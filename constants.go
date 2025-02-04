package main

const pi = 3.141
const (
	hello = "g"
	e     = 2
)

const (
	zero = iota
	_
	three
)

const (
	_         = iota
	KB uint64 = 1 << (10 * iota)
	MB
)

const (
	year          = 2017
	yearTyped int = 2017
)
