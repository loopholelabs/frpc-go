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

package utils

import (
	"google.golang.org/protobuf/reflect/protoreflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

// CamelCase returns the CamelCased name.
// If there is an interior underscore followed by a lower case letter,
// drop the underscore and convert the letter to upper case.
// There is a remote possibility of this rewrite causing a name collision,
// but it's so remote we're prepared to pretend it's nonexistent - since the
// C++ generator lowercases names, it's extremely unlikely to have two fields
// with different capitalizations.
// In short, _my_field_name_2 becomes XMyFieldName_2.
func CamelCase(s string) string {
	if s == "" {
		return ""
	}
	t := make([]byte, 0, 32)
	i := 0
	if s[0] == '_' { // Keep the initial _ if it exists
		i++
	}
	// Invariant: if the next letter is lower case, it must be converted
	// to upper case.
	// That is, we process a word at a time, where words are marked by _ or
	// upper case letter. Digits are treated as words.
	for ; i < len(s); i++ {
		c := s[i]
		if c == '_' && i+1 < len(s) && 'a' <= s[i+1] && s[i+1] <= 'z' {
			continue // Skip the underscore in s.
		}
		if c == '.' {
			continue
		}
		if '0' <= c && c <= '9' {
			t = append(t, c)
			continue
		}
		// Assume we have a letter now - if not, it's a bogus identifier.
		// The next word is a sequence of characters that must start upper case.
		if 'a' <= c && c <= 'z' {
			c ^= ' ' // Make it a capital letter.
		}
		t = append(t, c) // Guaranteed not lower case.
		// Accept lower case sequence that follows.
		for i+1 < len(s) && 'a' <= s[i+1] && s[i+1] <= 'z' {
			i++
			t = append(t, s[i])
		}
	}
	return string(t)
}

func CamelCaseFullName(name protoreflect.FullName) string {
	return CamelCase(string(name))
}

func CamelCaseName(name protoreflect.Name) string {
	return CamelCase(string(name))
}

func AppendString(inputs ...string) string {
	builder := new(strings.Builder)
	for _, s := range inputs {
		builder.WriteString(s)
	}

	return builder.String()
}

func FirstLowerCase(s string) string {
	if s == "" {
		return ""
	}
	r, n := utf8.DecodeRuneInString(s)
	return string(unicode.ToLower(r)) + s[n:]
}

func FirstLowerCaseName(name protoreflect.Name) string {
	return FirstLowerCase(string(name))
}

func MakeIterable(length int) []struct{} {
	return make([]struct{}, length)
}

func Counter(initial int) func() int {
	i := initial
	return func() int {
		i++
		return i
	}
}
