// Copyright (c) 2019 The Tor Project, inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptext

import (
	"bytes"
	"log"
	"testing"

	pt "git.torproject.org/pluggable-transports/goptlib.git"
)

func TestLogSeverity(t *testing.T) {
	tests := [...]struct {
		value    LogSeverity
		expected string
	}{
		// The good cases.
		{Debug, "debug"},
		{Info, "info"},
		{Notice, "notice"},
		{Warning, "warning"},
		{Error, "error"},

		// The bad cases.
		{99, "error"},
		{-99, "debug"},
	}

	for _, input := range tests {
		value := input.value.String()

		if input.expected != value {
			t.Errorf("%v.String() → \"%v\" (expected \"%v\")",
				input.value, value, input.expected)
		}
	}
}

func TestLogger(t *testing.T) {
	var buffer bytes.Buffer
	logger := log.New(logger{writer: &buffer, severity: Debug}, "", 0)

	defer func() {
		logger.SetOutput(pt.Stdout)
	}()

	logger.Print("Foo")
	logger.Print("Foo bar baz")

	if buffer.String() != "LOG MESSAGE=Foo SEVERITY=debug\nLOG MESSAGE=\"Foo bar baz\" SEVERITY=debug\n" {
		t.Errorf("Incorrect log output")
	}
}
