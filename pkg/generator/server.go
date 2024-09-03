// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"strings"

	"google.golang.org/protobuf/reflect/protoreflect"

	"github.com/loopholelabs/polyglot/v2/utils"
)

func getServerFields(services protoreflect.ServiceDescriptors) string {
	builder := new(strings.Builder)
	for i := 0; i < services.Len(); i++ {
		service := services.Get(i)
		serviceName := utils.CamelCase(string(service.Name()))
		builder.WriteString(utils.FirstLowerCase(serviceName))
		builder.WriteString(" ")
		builder.WriteString(serviceName)
		builder.WriteString(",")
		builder.WriteString(" ")
	}
	serverFields := builder.String()
	serverFields = serverFields[:len(serverFields)-2]
	return serverFields
}
