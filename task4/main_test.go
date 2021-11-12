package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"testing"
)

const textToProcess = "random, random, something random"

var errRead = errors.New("read err")
var errWrite = errors.New("write err")

type errReader struct {
	bytes.Buffer
}

func (*errReader) Read([]byte) (int, error) {
	return 0, errRead
}

func TestCountString(t *testing.T) {
	tests := []struct {
		name   string
		reader io.Reader
		s      string
		count  int
		err    error
	}{
		{name: "countFromString", reader: strings.NewReader(textToProcess), s: "random", count: 3, err: nil},
		{name: "countFromStringErr", reader: &errReader{*bytes.NewBufferString(textToProcess)}, s: "random", count: 0, err: errRead},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			count, err := CountString(tt.reader, tt.s)
			if count != tt.count || !errors.Is(err, tt.err) {
				t.Errorf("%s: countString(reader, %s) = %d, %v want %d, %v", tt.name, tt.s, count, err, tt.count, tt.err)
			}
		})
	}
}

type eerrWriter struct {
	bytes.Buffer
}

func (*eerrWriter) Write([]byte) (int, error) {
	return 0, errWrite
}

// reading result to buffer and initializing new buf for
// every test case with anon func
func TestReplaceString(t *testing.T) {
	// buffer to write result to
	var buf *bytes.Buffer

	tests := []struct {
		name              string
		rw                func() io.ReadWriter
		buf               *bytes.Buffer
		old, want, s1, s2 string
		err               error
	}{
		{name: "ok", rw: func() io.ReadWriter {
			buf = bytes.NewBufferString(textToProcess)
			return buf

		}, s1: "random", s2: "concrete", want: "concrete, concrete, something concrete", err: nil},

		{name: "Errreader", rw: func() io.ReadWriter {
			buf = bytes.NewBufferString(textToProcess)
			return &errReader{*buf}
		}, s1: "random", s2: "concrete", want: "random, random, something random", err: errRead},

		{name: "ErrWriter", rw: func() io.ReadWriter {
			buf = bytes.NewBufferString(textToProcess)
			return &eerrWriter{*buf}
		}, s1: "random", s2: "concrete", want: "random, random, something random", err: errWrite},
	}
	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			err := ReplaceString(tt.rw(), tt.s1, tt.s2)

			if !errors.Is(err, tt.err) || buf.String() != tt.want {
				t.Errorf("%s ReplaceString(wr, %s, %s) = %s, %v, want %s, %v", tt.name, tt.s1, tt.s2, buf.String(), err, tt.want, tt.err)
			}
		})
	}
}
