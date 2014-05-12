package main

import (
	"io"
	"strconv"
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
	if p.s.Peek() == scanner.EOF {
		return nil
	}

	h := p.s.Scan()
	i := 0
	for {
		r := p.s.Scan()

		if r == '\n' {
			break
		}

		sign := 1
		if r == '-' {
			sign = -1
			r = p.s.Scan()
		}

		n, _ := strconv.Atoi(p.s.TokenText())
		p.scores[string([]rune{h, p.header[i]})] = n * sign
		i++
	}

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
		state = state(p)
	}

	return p.scores
}
