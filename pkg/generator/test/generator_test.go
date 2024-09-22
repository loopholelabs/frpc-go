// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
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
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
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
