// SPDX-License-Identifier: Apache-2.0

package main

import (
	"io"
	"os"

	"github.com/loopholelabs/frpc-go/pkg/generator"
)

func main() {
	gen := generator.New()

	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	req, err := gen.UnmarshalRequest(data)
	if err != nil {
		panic(err)
	}

	res, err := gen.Generate(req)
	if err != nil {
		panic(err)
	}

	data, err = gen.MarshalResponse(res)
	if err != nil {
		panic(err)
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		panic(err)
	}
}
