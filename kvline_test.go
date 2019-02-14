// Copyright (c) 2019 The Tor Project, inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptext

import (
	"testing"
)

func TestKVLineEncoder(t *testing.T) {
	tests := [...]struct {
		value    map[string]string
		expected string
	}{
		{map[string]string{"A": "B", "CD": "EF"}, "A=B CD=EF"},
		{map[string]string{"AB": "C", "CDE": "F G"}, "AB=C CDE=\"F G\""},
		{map[string]string{"A": "Foo Bar Baz\r\t\n\"'"}, "A=\"Foo Bar Baz\\r\\t\\n\\\"\\'\""},
	}

	for _, input := range tests {
		encoded := kvlineEncode(input.value)
		if input.expected != encoded {
			t.Errorf("kvline_encode(%v) → \"%v\" (expected \"%v\")",
				input.value, encoded, input.expected)
		}
	}
}
