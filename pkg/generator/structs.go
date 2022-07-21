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
	"errors"
	"fmt"
	"github.com/loopholelabs/frisbee-go/protoc-gen-frpc/internal/utils"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	errUnknownKind        = errors.New("unknown or unsupported protoreflect.Kind")
	errUnknownCardinality = errors.New("unknown or unsupported protoreflect.Cardinality")
)

var (
	typeLUT = map[protoreflect.Kind]string{
		protoreflect.BoolKind:     "bool",
		protoreflect.Int32Kind:    "int32",
		protoreflect.Sint32Kind:   "int32",
		protoreflect.Uint32Kind:   "uint32",
		protoreflect.Int64Kind:    "int64",
		protoreflect.Sint64Kind:   "int64",
		protoreflect.Uint64Kind:   "uint64",
		protoreflect.Sfixed32Kind: "int32",
		protoreflect.Sfixed64Kind: "int64",
		protoreflect.Fixed32Kind:  "uint32",
		protoreflect.Fixed64Kind:  "uint64",
		protoreflect.FloatKind:    "float32",
		protoreflect.DoubleKind:   "float64",
		protoreflect.StringKind:   "string",
		protoreflect.BytesKind:    "[]byte",
	}

	encodeLUT = map[protoreflect.Kind]string{
		protoreflect.BoolKind:     ".Bool",
		protoreflect.Int32Kind:    ".Int32",
		protoreflect.Sint32Kind:   ".Int32",
		protoreflect.Uint32Kind:   ".Uint32",
		protoreflect.Int64Kind:    ".Int64",
		protoreflect.Sint64Kind:   ".Int64",
		protoreflect.Uint64Kind:   ".Uint64",
		protoreflect.Sfixed32Kind: ".Int32",
		protoreflect.Sfixed64Kind: ".Int64",
		protoreflect.Fixed32Kind:  ".Uint32",
		protoreflect.Fixed64Kind:  ".Uint64",
		protoreflect.StringKind:   ".String",
		protoreflect.FloatKind:    ".Float32",
		protoreflect.DoubleKind:   ".Float64",
		protoreflect.BytesKind:    ".Bytes",
		protoreflect.EnumKind:     ".Uint32",
	}

	decodeLUT = map[protoreflect.Kind]string{
		protoreflect.BoolKind:     ".Bool",
		protoreflect.Int32Kind:    ".Int32",
		protoreflect.Sint32Kind:   ".Int32",
		protoreflect.Uint32Kind:   ".Uint32",
		protoreflect.Int64Kind:    ".Int64",
		protoreflect.Sint64Kind:   ".Int64",
		protoreflect.Uint64Kind:   ".Uint64",
		protoreflect.Sfixed32Kind: ".Int32",
		protoreflect.Sfixed64Kind: ".Int64",
		protoreflect.Fixed32Kind:  ".Uint32",
		protoreflect.Fixed64Kind:  ".Uint64",
		protoreflect.StringKind:   ".String",
		protoreflect.FloatKind:    ".Float32",
		protoreflect.DoubleKind:   ".Float64",
		protoreflect.BytesKind:    ".Bytes",
		protoreflect.EnumKind:     ".Uint32",
	}

	kindLUT = map[protoreflect.Kind]string{
		protoreflect.BoolKind:     "packet.BoolKind",
		protoreflect.Int32Kind:    "packet.Int32Kind",
		protoreflect.Sint32Kind:   "packet.Int32Kind",
		protoreflect.Uint32Kind:   "packet.Uint32Kind",
		protoreflect.Int64Kind:    "packet.Int64Kind",
		protoreflect.Sint64Kind:   "packet.Int64Kind",
		protoreflect.Uint64Kind:   "packet.Uint64Kind",
		protoreflect.Sfixed32Kind: "packet.Int32Kind",
		protoreflect.Sfixed64Kind: "packet.Int64Kind",
		protoreflect.Fixed32Kind:  "packet.Uint32Kind",
		protoreflect.Fixed64Kind:  "packet.Uint64Kind",
		protoreflect.StringKind:   "packet.StringKind",
		protoreflect.FloatKind:    "packet.Float32Kind",
		protoreflect.DoubleKind:   "packet.Float64Kind",
		protoreflect.BytesKind:    "packet.BytesKind",
		protoreflect.EnumKind:     "packet.Uint32Kind",
	}
)

func findValue(field protoreflect.FieldDescriptor) string {
	if kind, ok := typeLUT[field.Kind()]; !ok {
		switch field.Kind() {
		case protoreflect.EnumKind:
			switch field.Cardinality() {
			case protoreflect.Optional, protoreflect.Required:
				return utils.CamelCase(string(field.Enum().FullName()))
			case protoreflect.Repeated:
				return utils.CamelCase(utils.AppendString(slice, string(field.Enum().FullName())))
			default:
				panic(errUnknownCardinality)
			}
		case protoreflect.MessageKind:
			if field.IsMap() {
				return utils.CamelCase(utils.AppendString(string(field.FullName()), mapSuffix))
			} else {
				switch field.Cardinality() {
				case protoreflect.Optional, protoreflect.Required:
					return utils.AppendString(pointer, utils.CamelCase(string(field.Message().FullName())))
				case protoreflect.Repeated:
					return utils.AppendString(slice, pointer, utils.CamelCase(string(field.Message().FullName())))
				default:
					panic(errUnknownCardinality)
				}
			}
		default:
			panic(errUnknownKind)
		}
	} else {
		if field.Cardinality() == protoreflect.Repeated {
			kind = slice + kind
		}
		return kind
	}
}

type encodingFields struct {
	MessageFields []protoreflect.FieldDescriptor
	SliceFields   []protoreflect.FieldDescriptor
	Values        []string
}

func getEncodingFields(fields protoreflect.FieldDescriptors) encodingFields {
	var messageFields []protoreflect.FieldDescriptor
	var sliceFields []protoreflect.FieldDescriptor
	var values []string

	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if field.Cardinality() == protoreflect.Repeated && !field.IsMap() {
			sliceFields = append(sliceFields, field)
		} else {
			if encoder, ok := encodeLUT[field.Kind()]; !ok {
				switch field.Kind() {
				case protoreflect.MessageKind:
					messageFields = append(messageFields, field)
				default:
					panic(errUnknownKind)
				}
			} else {
				if field.Kind() == protoreflect.EnumKind {
					values = append(values, fmt.Sprintf("%s(uint32(x.%s))", encoder, utils.CamelCase(string(field.Name()))))
				} else {
					values = append(values, fmt.Sprintf("%s(x.%s)", encoder, utils.CamelCase(string(field.Name()))))
				}
			}
		}
	}
	return encodingFields{
		MessageFields: messageFields,
		SliceFields:   sliceFields,
		Values:        values,
	}
}

type decodingFields struct {
	MessageFields []protoreflect.FieldDescriptor
	SliceFields   []protoreflect.FieldDescriptor
	Other         []protoreflect.FieldDescriptor
}

func getDecodingFields(fields protoreflect.FieldDescriptors) decodingFields {
	var messageFields []protoreflect.FieldDescriptor
	var sliceFields []protoreflect.FieldDescriptor
	var other []protoreflect.FieldDescriptor

	for i := 0; i < fields.Len(); i++ {
		field := fields.Get(i)
		if field.Cardinality() == protoreflect.Repeated && !field.IsMap() {
			sliceFields = append(sliceFields, field)
		} else {
			if _, ok := decodeLUT[field.Kind()]; !ok {
				switch field.Kind() {
				case protoreflect.MessageKind:
					messageFields = append(messageFields, field)
				default:
					panic(errUnknownKind)
				}
			} else {
				other = append(other, field)
			}
		}
	}

	return decodingFields{
		MessageFields: messageFields,
		SliceFields:   sliceFields,
		Other:         other,
	}
}

func getKind(kind protoreflect.Kind) string {
	var outKind string
	var ok bool
	if outKind, ok = kindLUT[kind]; !ok {
		switch kind {
		case protoreflect.MessageKind:
			outKind = packetAnyKind
		default:
			panic(errUnknownKind)
		}
	}
	return outKind
}

func getLUTEncoder(kind protoreflect.Kind) string {
	return encodeLUT[kind]
}

func getLUTDecoder(kind protoreflect.Kind) string {
	return decodeLUT[kind]
}

func getKindLUT(kind protoreflect.Kind) string {
	return kindLUT[kind]
}
