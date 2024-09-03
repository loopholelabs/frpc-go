// SPDX-License-Identifier: Apache-2.0

package version

import (
	_ "embed"
)

//go:embed current_version
var Version string
