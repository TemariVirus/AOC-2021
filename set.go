package main

import (
	"maps"
)

type Set[T comparable] struct {
	data map[T]bool
}

func makeSet[T comparable](capacity int) Set[T] {
	return Set[T]{make(map[T]bool, capacity)}
}

func makeSetFrom[T comparable](items []T) Set[T] {
	set := Set[T]{make(map[T]bool, len(items))}
	for _, item := range items {
		set.add(item)
	}
	return set
}

func (this Set[T]) len() int {
	return len(this.data)
}

func (this Set[T]) toArray() []T {
	keys := make([]T, 0, len(this.data))
	for k := range this.data {
		keys = append(keys, k)
	}
	return keys
}

func (this Set[T]) add(item T) {
	this.data[item] = true
}

func (this Set[T]) remove(item T) {
	delete(this.data, item)
}

func (this Set[T]) contains(item T) bool {
	_, ok := this.data[item]
	return ok
}

func (this Set[T]) containsSet(set Set[T]) bool {
	if this.len() < set.len() {
		return false
	}

	for key := range set.data {
		if !this.contains(key) {
			return false
		}
	}

	return true
}

func (this Set[T]) intersection(other Set[T]) Set[T] {
	result := make(map[T]bool)
	smaller := this.data
	bigger := other.data
	if len(bigger) < len(smaller) {
		smaller, bigger = bigger, smaller
	}
	for key := range smaller {
		if _, ok := bigger[key]; ok {
			result[key] = true
		}
	}

	return Set[T]{result}
}

func (this Set[T]) union(other Set[T]) Set[T] {
	result := make(map[T]bool)
	maps.Copy(result, this.data)
	maps.Copy(result, other.data)

	return Set[T]{result}
}

func (this Set[T]) exclue(other Set[T]) Set[T] {
	result := make(map[T]bool)
	maps.Copy(result, this.data)
	for key := range other.data {
		delete(result, key)
	}

	return Set[T]{result}
}

func (this Set[T]) equals(other Set[T]) bool {
	if this.len() != other.len() {
		return false
	}

	for key := range this.data {
		if !other.contains(key) {
			return false
		}
	}

	return true
}

func (this Set[T]) eq(other Equatable) bool {
	if set, ok := other.(Set[T]); ok {
		return this.equals(set)
	}
	return false
}

func (this Set[T]) ne(other Equatable) bool {
	if set, ok := other.(Set[T]); ok {
		return !this.equals(set)
	}
	return false
}
