// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/loopholelabs/polyglot/v2"
	"github.com/loopholelabs/testing/conn/pair"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRPC(t *testing.T) {
	cConn, sConn, err := pair.New()
	assert.NoError(t, err, "Client server pair creation failed")

	client, err := NewClient(nil, nil)
	assert.NoError(t, err, "Client creation failed")

	err = client.FromConn(cConn)
	assert.NoError(t, err, "Client connection assignment failed")

	serverMethods := new(svc)
	serverMethods.t = t
	server, err := NewServer(serverMethods, nil, nil)
	assert.NoError(t, err, "Server creation failed")

	go server.ServeConn(sConn)

	t.Run("Synchronous Request", func(t *testing.T) {
		testSynchronous(client, t)
	})

	t.Run("Bi-directional Stream", func(t *testing.T) {
		testBidirectional(client, t)
	})

	t.Run("Server Stream", func(t *testing.T) {
		testServerStreaming(client, t)
	})

	t.Run("Client Stream", func(t *testing.T) {
		testClientStreaming(client, t)
	})
}

func TestClientStreamIDs(t *testing.T) {
	ids := newClientStreamIDs()
	seen := make(map[uint16]struct{}, 1<<16)
	for range 1 << 16 {
		id, err := ids.acquire()
		require.NoError(t, err)
		_, exists := seen[id]
		require.False(t, exists)
		seen[id] = struct{}{}
	}

	_, err := ids.acquire()
	require.Error(t, err)

	ids.release(42)
	id, err := ids.acquire()
	require.NoError(t, err)
	require.Equal(t, uint16(42), id)
}

func testSynchronous(client *Client, t *testing.T) {
	ctx := context.Background()
	req := &Request{
		Message: "Hello World",
		Corpus:  RequestUNIVERSAL,
	}
	response, err := client.EchoService.Echo(ctx, req)
	assert.NoError(t, err, "Request error")
	assert.Equal(t, "Hello World", response.Message)
	assert.Equal(t, "Hello World", response.Test.Message)
	assert.Equal(t, Potato, response.Test.Checker)
}

func testBidirectional(client *Client, t *testing.T) {
	ctx := context.Background()
	req := &Request{
		Message: "Hello World",
		Corpus:  RequestUNIVERSAL,
	}

	stream, err := client.EchoService.EchoStream(ctx, req)
	assert.NoError(t, err)

	roundtrips := 0
	for {
		res, err := stream.Recv()
		roundtrips++
		assert.NoError(t, err, "Request error")
		assert.Equal(t, "Hello World", res.Message)
		assert.Equal(t, "Hello World", res.Test.Message)
		assert.Equal(t, Potato, res.Test.Checker)
		if roundtrips == 100 {
			err = stream.CloseSend()
			assert.NoError(t, err)
			_, err = stream.Recv()
			assert.ErrorIs(t, err, io.EOF)
			break
		}

		err = stream.Send(req)
		assert.NoError(t, err)
	}
	assert.Equal(t, 100, roundtrips)
}

func testServerStreaming(client *Client, t *testing.T) {
	ctx := context.Background()
	search := NewSearchResponse()
	search.Results = []*SearchResponseResult{
		{
			Url:      "https://google.com",
			Title:    "Google",
			Snippets: []string{},
		},
	}
	stream, err := client.EchoService.Search(ctx, search)
	assert.NoError(t, err)

	received := 0
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			assert.Equal(t, 10, received)
			break
		}
		received++
		assert.NoError(t, err)
		assert.Equal(t, "Hello World", res.Message)
		assert.Equal(t, "Hello World", res.Test.Message)
		assert.Equal(t, Potato, res.Test.Checker)
	}
}

func testClientStreaming(client *Client, t *testing.T) {
	ctx := context.Background()
	data := &Data{Message: "Hello World", Checker: Potato}
	stream, err := client.EchoService.Upload(ctx, data)
	assert.NoError(t, err)

	for i := 0; i < 10; i++ {
		err := stream.Send(data)
		assert.NoError(t, err)
	}
	res, err := stream.CloseAndRecv()
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", res.Message)
}

func TestRPCInvalidConnection(t *testing.T) {
	// Create non-Frisbee server the client can connect to but not exchange
	// messages, so the connection will be broken soon after connect.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "test")
	}))
	t.Cleanup(ts.Close)

	// Create a client and connect to test server.
	client, err := NewClient(nil, nil)
	require.NoError(t, err)

	err = client.Connect(ts.Listener.Addr().String())
	require.NoError(t, err)

	// Make RPC request with a 3s timeout.
	timeout := 3 * time.Second
	if os.Getenv("CI") != "" {
		timeout = 10 * time.Second
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	t.Cleanup(cancel)

	req := &Request{
		Message: "Hello World",
		Corpus:  RequestUNIVERSAL,
	}
	response, err := client.EchoService.Echo(ctx, req)

	// Verify request doesn't block forever.
	require.NoError(t, ctx.Err())

	// Verify request fails.
	require.Error(t, err)
	require.Nil(t, response)
}

func TestEncodeDecodePreservesNilFields(t *testing.T) {
	r := &Response{Message: "test", Test: nil}
	b := polyglot.NewBuffer()
	r.Encode(b)

	got := &Response{}
	err := got.Decode(b.Bytes())
	require.NoError(t, err)
	require.Equal(t, r, got)
}

func TestOneofRoundTrip(t *testing.T) {
	roundTrip := func(t *testing.T, want *OneofPayload) *OneofPayload {
		t.Helper()
		b := polyglot.NewBuffer()
		want.Encode(b)
		got := NewOneofPayload()
		require.NoError(t, got.Decode(b.Bytes()))
		require.Equal(t, want, got)
		return got
	}

	t.Run("not set", func(t *testing.T) {
		got := roundTrip(t, &OneofPayload{Ordinary: "ordinary"})
		require.Equal(t, OneofPayloadValueCaseNotSet, got.WhichValue())
		require.Equal(t, OneofPayloadSecondaryCaseNotSet, got.WhichSecondary())
	})

	t.Run("scalar default values remain selected", func(t *testing.T) {
		for name, set := range map[string]func(*OneofPayload){
			"string": func(x *OneofPayload) { x.SetText("") },
			"enum":   func(x *OneofPayload) { x.SetEnumValue(Potato) },
			"bool":   func(x *OneofPayload) { x.SetEnabled(false) },
			"int64":  func(x *OneofPayload) { x.SetNumber(0) },
		} {
			t.Run(name, func(t *testing.T) {
				want := NewOneofPayload()
				set(want)
				got := roundTrip(t, want)
				require.Equal(t, want.WhichValue(), got.WhichValue())
			})
		}
	})

	t.Run("bytes", func(t *testing.T) {
		want := NewOneofPayload()
		want.SetRaw([]byte("raw payload"))
		roundTrip(t, want)
	})

	t.Run("message", func(t *testing.T) {
		want := NewOneofPayload()
		want.SetPayload(&Data{Message: "nested", Checker: Monkey})
		roundTrip(t, want)

		want.SetPayload(nil)
		got := roundTrip(t, want)
		require.Equal(t, OneofPayloadValueCasePayload, got.WhichValue())
		require.Nil(t, got.GetPayload())
	})

	t.Run("multiple independent groups", func(t *testing.T) {
		want := NewOneofPayload()
		want.SetText("primary")
		want.SetCode(42)
		got := roundTrip(t, want)
		require.Equal(t, "primary", got.GetText())
		require.Equal(t, uint32(42), got.GetCode())
	})
}

func TestOneofSettersAndReuse(t *testing.T) {
	x := NewOneofPayload()
	x.SetText("old")
	x.SetRaw([]byte("new"))
	require.Equal(t, OneofPayloadValueCaseRaw, x.WhichValue())
	require.Empty(t, x.Text)
	require.Empty(t, x.GetText())
	require.Equal(t, []byte("new"), x.GetRaw())

	want := NewOneofPayload()
	want.SetNumber(123)
	b := polyglot.NewBuffer()
	want.Encode(b)
	require.NoError(t, x.Decode(b.Bytes()))
	require.Equal(t, OneofPayloadValueCaseNumber, x.WhichValue())
	require.Nil(t, x.Raw)
	require.Equal(t, int64(123), x.GetNumber())

	x.ClearValue()
	require.Equal(t, OneofPayloadValueCaseNotSet, x.WhichValue())
}

func TestOneofEncodesOnlyActiveMember(t *testing.T) {
	clean := NewOneofPayload()
	clean.SetText("active")
	cleanBuffer := polyglot.NewBuffer()
	clean.Encode(cleanBuffer)

	withInactiveData := NewOneofPayload()
	withInactiveData.SetText("active")
	withInactiveData.Raw = make([]byte, 1<<20)
	withInactiveData.Payload = &Data{Message: "inactive"}
	withInactiveData.Number = 42
	withInactiveDataBuffer := polyglot.NewBuffer()
	withInactiveData.Encode(withInactiveDataBuffer)

	require.Equal(t, cleanBuffer.Bytes(), withInactiveDataBuffer.Bytes())
}

func TestOneofMemberNamedError(t *testing.T) {
	want := NewErrorOneof()
	want.SetError(&Data{Message: "application error", Checker: Monkey})
	b := polyglot.NewBuffer()
	want.Encode(b)

	got := NewErrorOneof()
	require.NoError(t, got.Decode(b.Bytes()))
	require.Equal(t, ErrorOneofResultCaseError, got.WhichResult())
	require.Equal(t, want.GetError(), got.GetError())
}

func TestOneofRejectsUnknownCase(t *testing.T) {
	x := NewOneofPayload()
	x.SetText("value")
	b := polyglot.NewBuffer()
	x.Encode(b)
	encoded := append([]byte(nil), b.Bytes()...)

	d := polyglot.Decoder(encoded)
	require.False(t, d.Nil())
	_, err := d.Error()
	require.Error(t, err)
	_, err = d.Uint8()
	require.NoError(t, err)
	_, err = d.String()
	require.NoError(t, err)
	caseOffset := len(encoded) - len(*d)
	require.Less(t, caseOffset+1, len(encoded))
	encoded[caseOffset+1] = 99

	err = NewOneofPayload().Decode(encoded)
	require.ErrorIs(t, err, ErrUnknownOneofCase)
}

func BenchmarkOneofEncode(b *testing.B) {
	x := NewSampleMessage()
	x.SetName("a representative oneof payload")
	buf := polyglot.NewBuffer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		x.Encode(buf)
	}
}

func BenchmarkOneofDecode(b *testing.B) {
	x := NewSampleMessage()
	x.SetName("a representative oneof payload")
	buf := polyglot.NewBuffer()
	x.Encode(buf)
	encoded := buf.Bytes()
	got := NewSampleMessage()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := got.Decode(encoded); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkRegularFieldEncode(b *testing.B) {
	x := &RegularPayload{Text: "a representative oneof payload"}
	buf := polyglot.NewBuffer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		buf.Reset()
		x.Encode(buf)
	}
}

func BenchmarkRegularFieldDecode(b *testing.B) {
	x := &RegularPayload{Text: "a representative oneof payload"}
	buf := polyglot.NewBuffer()
	x.Encode(buf)
	encoded := buf.Bytes()
	got := NewRegularPayload()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := got.Decode(encoded); err != nil {
			b.Fatal(err)
		}
	}
}
