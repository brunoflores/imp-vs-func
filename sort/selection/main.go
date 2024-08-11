package main

import (
	"strings"
)

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i := range ts {
		us[i] = f(ts[i])
	}
	return us
}

type Comp int

const (
	Eq Comp = iota
	Less
	More
)

type Comparable[T any] interface {
	CompareTo(T) Comp
	Show() string
}

type MyString struct {
	c string
}

func NewMyString(str string) *MyString {
	return &MyString{c: str}
}

func (s *MyString) Show() string {
	return s.c
}

func (s *MyString) CompareTo(other *MyString) Comp {
	if s.c == other.c {
		return Eq
	} else if s.c < other.c {
		return Less
	} else {
		return More
	}
}

type Sort[T Comparable[T]] interface {
	Sort([]T)
	Show([]T) string
	exch(int, int)
}

type SelectionSort[T Comparable[T]] struct {
	Items []T
}

func (s *SelectionSort[T]) exch(i, j int) {
	t := s.Items[i]
	s.Items[i] = s.Items[j]
	s.Items[j] = t
}

func (s *SelectionSort[T]) Sort() {
	var n int = len(s.Items)
	for i := 0; i < n; i++ {
		var min int = i
		for j := i + 1; j < n; j++ {
			if s.Items[j].CompareTo(s.Items[min]) == Less {
				min = j
			}
		}
		s.exch(i, min)
	}
}

func (s *SelectionSort[T]) Show() string {
	return strings.Join(Map(s.Items, func(i T) string { return i.Show() }), " ")
}
