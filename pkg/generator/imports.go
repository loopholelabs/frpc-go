// SPDX-License-Identifier: Apache-2.0

package generator

var (
	requiredImports = []string{
		"errors",
		"net",
		"sync",
		"context",
		"github.com/loopholelabs/polyglot/v2",
	}

	serviceImports = []string{
		"crypto/tls",
		"github.com/loopholelabs/frisbee-go",
		"github.com/loopholelabs/frisbee-go/pkg/packet",
		"github.com/loopholelabs/logging/types",
	}

	streamMethodImports = []string{
		"sync/atomic",
		"io",
	}
)
