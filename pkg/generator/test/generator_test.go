/*
	Copyright 2022 Loophole Labs

	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

		   http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/

package test

import (
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/loopholelabs/testing/conn/pair"
	"github.com/stretchr/testify/assert"
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
        fmt.Println("Sending data")
		err := stream.Send(data)
		assert.NoError(t, err)
	}
	res, err := stream.CloseAndRecv()
	assert.NoError(t, err)
	assert.Equal(t, "Hello World", res.Message)
}
