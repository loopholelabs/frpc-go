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
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

type svc struct {
	t *testing.T
}

func (s svc) Echo(ctx context.Context, request *Request) (*Response, error) {
	assert.Equal(s.t, "Hello World", request.Message)
	assert.Equal(s.t, RequestUNIVERSAL, request.Corpus)
	return &Response{Message: "Hello World", Test: &Data{
		Message: "Hello World",
		Checker: Potato,
	}}, nil
}

func (s svc) EchoStream(srv *EchoStreamServer) error {
	for {
		request, err := srv.Recv()
		if err == io.EOF {
			err = srv.CloseSend()
			assert.NoError(s.t, err)
			break
		}

		assert.NoError(s.t, err)

		assert.Equal(s.t, "Hello World", request.Message)
		assert.Equal(s.t, RequestUNIVERSAL, request.Corpus)

		response := &Response{Message: "Hello World", Test: &Data{
			Message: "Hello World",
			Checker: Potato,
		}}

		err = srv.Send(response)
		assert.NoError(s.t, err)
	}
	return nil
}

func (s svc) Testy(ctx context.Context, response *SearchResponse) (*StockPricesWrapper, error) {
	panic("not implemented")
}

func (s svc) Search(req *SearchResponse, srv *SearchServer) error {
	assert.Equal(s.t, 1, len(req.Results))
	for i := 0; i < 10; i++ {
		err := srv.Send(&Response{Message: "Hello World", Test: &Data{
			Message: "Hello World",
			Checker: Potato,
		}})
		assert.NoError(s.t, err)
	}
	return srv.CloseSend()
}

func (s svc) Upload(srv *UploadServer) error {
	println("upload")
	received := 0
	for {
		res, err := srv.Recv()
		if err == io.EOF {
			assert.Equal(s.t, 11, received)
			return srv.CloseAndSend(&Response{Message: "Hello World", Test: &Data{}})
		}
		received += 1
		assert.NoError(s.t, err)
		assert.Equal(s.t, "Hello World", res.Message)
	}
}
