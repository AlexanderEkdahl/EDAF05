package main

import (
	"io"
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
		s.Scan()
		name := s.TokenText()
		for s.Scan() != '\n' {
		}

		s.Scan()
		subjects = append(subjects, &subject{name: name, value: s.TokenText()})
		s.Scan()
	}

	return subjects
}
