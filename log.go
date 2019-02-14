// Copyright (c) 2019 The Tor Project, inc. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptext

import (
	"fmt"
	"io"
	"log"
	"strings"

	pt "git.torproject.org/pluggable-transports/goptlib.git"
)

// LogSeverity level. This mapping is defined in Tor's pt-spec.txt.
type LogSeverity int

const (
	// Debug log level.
	Debug LogSeverity = 0

	// Info log level.
	Info LogSeverity = 1

	// Notice log level.
	Notice LogSeverity = 2

	// Warning log level.
	Warning LogSeverity = 3

	// Error log level.
	Error LogSeverity = 4
)

func (log_severity LogSeverity) String() string {
	labels := [...]string{
		"debug",
		"info",
		"notice",
		"warning",
		"error",
	}

	value := log_severity

	// Smaller values than Debug? Reset to Debug.
	if value < Debug {
		value = Debug
	}

	// Larger values than Error? Reset to Error.
	if value > Error {
		value = Error
	}

	return labels[value]
}

type logger struct {
	writer   io.Writer
	severity LogSeverity
}

func (l logger) Write(p []byte) (n int, err error) {
	data := map[string]string{
		"SEVERITY": l.severity.String(),
		// Remove the trailing new line that the `log` package appends to all strings.
		"MESSAGE": strings.TrimRight(string(p), "\n"),
	}
	log := fmt.Sprintf("LOG %s\n", kvlineEncode(data))

	return io.WriteString(l.writer, log)
}

// NewPTLogger creates a new Logger instance that is writes its events to
// Standard Out which the Tor process will handle and log accordingly to its
// event log.
func NewPTLogger(severity LogSeverity, prefix string, flags int) *log.Logger {
	result := log.New(logger{writer: pt.Stdout, severity: severity}, prefix, flags)
	return result
}
