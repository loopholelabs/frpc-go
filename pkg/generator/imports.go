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

package generator

var (
	requiredImports = []string{
		"errors",
		"github.com/loopholelabs/polyglot-go",
		"net",
	}

	serviceImports = []string{
		"context",
		"crypto/tls",
		"github.com/loopholelabs/frisbee-go",
		"github.com/loopholelabs/frisbee-go/pkg/packet",
		"github.com/rs/zerolog",
	}

	streamMethodImports = []string{
		"go.uber.org/atomic",
		"io",
	}

	methodImports = []string{
		"sync",
	}
)
