package main

import (
	"fmt"
	"io"
	"text/scanner"
)

type parser struct {
	scores map[string]int
	header []rune
	s      scanner.Scanner
}

type stateFn func(*parser) stateFn

func parseComment(p *parser) stateFn {
	if p.s.Peek() != '#' {
		return parseHeader
	}

	for p.s.Scan() != '\n' {
		// Consume both \r\n - why is this required?
		// p.s.Scan()
	}

	return parseComment
}

func parseHeader(p *parser) stateFn {
	for p.s.Peek() != '\n' {
		p.header = append(p.header, p.s.Scan())
	}

	p.s.Scan()

	return parseValues
}

func parseValues(p *parser) stateFn {
	if p.s.Peek() == '\n' {
		fmt.Println("the end")
		return nil
	}

	for p.s.Scan() != '\n' {
	}

	fmt.Println("new values")

	return parseValues
}

func parseScores(r io.Reader) map[string]int {
	p := &parser{
		scores: make(map[string]int),
		header: make([]rune, 0),
	}
	p.s.Init(r)
	p.s.Mode = scanner.ScanInts
	p.s.Whitespace = 1<<'\t' | 1<<'\r' | 1<<' '

	for state := parseComment; state != nil; {
		fmt.Printf("%g\n", state)
		state = state(p)
	}

	return p.scores
}
