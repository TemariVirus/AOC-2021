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

func (s Set[T]) len() int {
	return len(s.data)
}

func (s Set[T]) toArray() []T {
	keys := make([]T, 0, len(s.data))
	for k := range s.data {
		keys = append(keys, k)
	}
	return keys
}

func (s Set[T]) add(item T) {
	s.data[item] = true
}

func (s Set[T]) remove(item T) {
	delete(s.data, item)
}

func (s Set[T]) contains(item T) bool {
	_, ok := s.data[item]
	return ok
}

func (s Set[T]) containsSet(set Set[T]) bool {
	if s.len() < set.len() {
		return false
	}

	for key := range set.data {
		if !s.contains(key) {
			return false
		}
	}

	return true
}

func (s Set[T]) intersect(other Set[T]) Set[T] {
	result := make(map[T]bool)
	smaller := s.data
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

func (s Set[T]) union(other Set[T]) Set[T] {
	result := make(map[T]bool)
	maps.Copy(result, s.data)
	maps.Copy(result, other.data)

	return Set[T]{result}
}

func (s Set[T]) exclue(other Set[T]) Set[T] {
	result := make(map[T]bool)
	maps.Copy(result, s.data)
	for key := range other.data {
		delete(result, key)
	}

	return Set[T]{result}
}

func (s Set[T]) copy() Set[T] {
	clone := make(map[T]bool, s.len())
	maps.Copy(clone, s.data)
	return Set[T]{clone}
}

func (s Set[T]) equals(other Set[T]) bool {
	if s.len() != other.len() {
		return false
	}

	for key := range s.data {
		if !other.contains(key) {
			return false
		}
	}

	return true
}

func (s Set[T]) eq(other Equatable) bool {
	if set, ok := other.(Set[T]); ok {
		return s.equals(set)
	}
	return false
}

func (s Set[T]) ne(other Equatable) bool {
	if set, ok := other.(Set[T]); ok {
		return !s.equals(set)
	}
	return false
}
