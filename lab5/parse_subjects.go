package main

import (
	"io"
	"strings"
	"text/scanner"
)

func parseSubjects(r io.Reader) []*subject {
	subjects := make([]*subject, 0)
	var s scanner.Scanner
	s.Init(r)
	s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '

	for {
		if s.Peek() == scanner.EOF {
			break
		}
		s.Scan()

		name := make([]string, 0)
		for s.Scan() != '\n' {
			name = append(name, s.TokenText())
		}

		s.Scan()
		subjects = append(subjects, &subject{name: strings.Join(name, ""), value: s.TokenText()})
		s.Scan()
	}

	return subjects
}
