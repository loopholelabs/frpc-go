// SPDX-License-Identifier: Apache-2.0

package test

import (
	"context"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type svc struct {
	t *testing.T
}

func (s svc) Echo(_ context.Context, request *Request) (*Response, error) {
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

func (s svc) Testy(_ context.Context, _ *SearchResponse) (*StockPricesWrapper, error) {
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
	received := 0
	for {
		res, err := srv.Recv()
		if err == io.EOF {
			assert.Equal(s.t, 11, received)
			return srv.CloseAndSend(&Response{Message: "Hello World", Test: &Data{}})
		}
		received++
		assert.NoError(s.t, err)
		assert.Equal(s.t, "Hello World", res.Message)
	}
}
