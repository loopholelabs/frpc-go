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
	"bytes"
	"github.com/loopholelabs/frpc-go/internal/version"
	"github.com/loopholelabs/frpc-go/templates"
	"github.com/loopholelabs/polyglot-go/pkg/generator"
	"github.com/loopholelabs/polyglot-go/pkg/utils"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/pluginpb"
	"text/template"
)

type Generator struct {
	options *protogen.Options
}

var templ *template.Template
var streamMethods = make(map[protoreflect.FullName]bool)

func init() {
	templ = template.Must(template.New("main").Funcs(template.FuncMap{
		"CamelCase":          utils.CamelCaseFullName,
		"CamelCaseName":      utils.CamelCaseName,
		"MakeIterable":       utils.MakeIterable,
		"Counter":            utils.Counter,
		"FirstLowerCase":     utils.FirstLowerCase,
		"FirstLowerCaseName": utils.FirstLowerCaseName,
		"GetServerFields":    getServerFields,
		"UsedForStreaming": func(typeName protoreflect.FullName) bool {
			return streamMethods[typeName]
		},
	}).ParseFS(templates.FS, "*"))
}

func New() *Generator {
	return &Generator{
		options: &protogen.Options{
			ParamFunc:         func(name string, value string) error { return nil },
			ImportRewriteFunc: func(path protogen.GoImportPath) protogen.GoImportPath { return path },
		},
	}
}

func (*Generator) UnmarshalRequest(buf []byte) (*pluginpb.CodeGeneratorRequest, error) {
	req := new(pluginpb.CodeGeneratorRequest)
	return req, proto.Unmarshal(buf, req)
}

func (*Generator) MarshalResponse(res *pluginpb.CodeGeneratorResponse) ([]byte, error) {
	return proto.Marshal(res)
}

func (g *Generator) Generate(req *pluginpb.CodeGeneratorRequest) (res *pluginpb.CodeGeneratorResponse, err error) {
	plugin, err := g.options.New(req)
	if err != nil {
		return nil, err
	}

	var tplBuffer bytes.Buffer
	if err := templ.ExecuteTemplate(&tplBuffer, "customEncode.templ", nil); err != nil {
		return nil, err
	}
	customEncode := tplBuffer.String()
	tplBuffer.Reset()

	if err := templ.ExecuteTemplate(&tplBuffer, "customDecode.templ", nil); err != nil {
		return nil, err
	}
	customDecode := tplBuffer.String()
	tplBuffer.Reset()

	if err := templ.ExecuteTemplate(&tplBuffer, "customFields.templ", nil); err != nil {
		return nil, err
	}
	customFields := tplBuffer.String()

	gen := generator.New()
	gen.CustomEncode = func() string {
		return customEncode
	}
	gen.CustomDecode = func() string {
		return customDecode
	}
	gen.CustomFields = func() string {
		return customFields
	}

	for _, f := range plugin.Files {
		if !f.Generate {
			continue
		}
		genFile := plugin.NewGeneratedFile(FileName(f.GeneratedFilenamePrefix), f.GoImportPath)

		packageName := string(f.Desc.Package().Name())
		if packageName == "" {
			packageName = string(f.GoPackageName)
		}

		numServices := f.Desc.Services().Len()

		numMethods := 0
		streamMethods = make(map[protoreflect.FullName]bool)
		for i := 0; i < numServices; i++ {
			nM := f.Desc.Services().Get(i).Methods().Len()
			numMethods += nM
			for m := 0; m < nM; m++ {
				method := f.Desc.Services().Get(i).Methods().Get(m)
				if method.IsStreamingClient() {
					streamMethods[method.Input().FullName()] = true
				}
				if method.IsStreamingServer() {
					streamMethods[method.Output().FullName()] = true
				}
			}
		}

		err = templ.ExecuteTemplate(genFile, "prebase.templ", map[string]interface{}{
			"pluginVersion":       version.Version,
			"sourcePath":          f.Desc.Path(),
			"package":             packageName,
			"requiredImports":     requiredImports,
			"serviceImports":      serviceImports,
			"methodImports":       methodImports,
			"streamMethodImports": streamMethodImports,
			"numServices":         numServices,
			"numMethods":          numMethods,
			"numStreamMethods":    len(streamMethods),
		})
		if err != nil {
			return nil, err
		}

		err = gen.ExecuteTemplate(genFile, f, packageName, false)
		if err != nil {
			return nil, err
		}

		err = templ.ExecuteTemplate(genFile, "base.templ", map[string]interface{}{
			"enums":            f.Desc.Enums(),
			"messages":         f.Desc.Messages(),
			"services":         f.Desc.Services(),
			"numServices":      numServices,
			"numMethods":       numMethods,
			"numStreamMethods": len(streamMethods),
		})
		if err != nil {
			return nil, err
		}
	}

	return plugin.Response(), nil
}
