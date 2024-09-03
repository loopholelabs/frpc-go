// SPDX-License-Identifier: Apache-2.0

package generator

import "github.com/loopholelabs/polyglot/v2/utils"

const extension = ".frpc.go"

func FileName(name string) string {
	return utils.AppendString(name, extension)
}
