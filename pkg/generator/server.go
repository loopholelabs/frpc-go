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

import (
	"github.com/loopholelabs/frisbee-go/protoc-gen-frpc/internal/utils"
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
)

func getServerFields(services protoreflect.ServiceDescriptors) string {
	builder := new(strings.Builder)
	for i := 0; i < services.Len(); i++ {
		service := services.Get(i)
		serviceName := utils.CamelCase(string(service.Name()))
		builder.WriteString(utils.FirstLowerCase(serviceName))
		builder.WriteString(space)
		builder.WriteString(serviceName)
		builder.WriteString(comma)
		builder.WriteString(space)
	}
	serverFields := builder.String()
	serverFields = serverFields[:len(serverFields)-2]
	return serverFields
}
