package main

type Equatable interface {
	eq(Equatable) bool
	ne(Equatable) bool
}
