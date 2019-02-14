// Copyright (c) 2019 The Tor Project, inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptext

import (
	"bytes"
	"fmt"
	"sort"
	"unicode"
)

func kvlineValueNeedsEscape(input string) bool {
	for _, c := range input {
		if c == '\'' || c == '"' || !unicode.IsPrint(c) || unicode.IsSpace(c) {
			return true
		}
	}

	return false
}

func kvlineEscapeValue(input string) string {
	if !kvlineValueNeedsEscape(input) {
		return input
	}

	// This code could benefit from the strings.Builder, but we cannot use that
	// because we want to work on Debian releases that have older Go versions.
	var result bytes.Buffer

	for _, c := range input {
		switch c {
		case '\'':
			result.WriteRune('\\')
			result.WriteRune('\'')
		case '"':
			result.WriteRune('\\')
			result.WriteRune('"')
		case '\n':
			result.WriteRune('\\')
			result.WriteRune('n')
		case '\t':
			result.WriteRune('\\')
			result.WriteRune('t')
		case '\r':
			result.WriteRune('\\')
			result.WriteRune('r')
		default:
			result.WriteRune(c)
		}
	}

	return fmt.Sprintf("\"%s\"", result.String())
}

func kvlineEncode(input map[string]string) string {
	// We need to make sure the order of keys are deterministic when we
	// serialize the K/V data structure into a string.
	keys := make([]string, 0, len(input))

	for key := range input {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	// This code could be refactored into using Go 1.10+ `strings.Builder`
	// type, but for now we can perfectly fine with the additional allocated
	// done by doing `+=` on strings :-)
	var result string

	for index := range keys {
		key := keys[index]
		value := input[key]

		// Are we not the first iteration we need to add a space.
		if index != 0 {
			result += " "
		}

		result += fmt.Sprintf("%s=%s",
			key, kvlineEscapeValue(value))
	}

	return result
}
