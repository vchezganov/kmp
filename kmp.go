package kmp

import (
	"errors"
)

// interfaceKMP is a required interface to be implemented for Knuth–Morris–Pratt algorithm
type interfaceKMP interface {
	At(i int) interface{}
	Len() int
	EqualTo(i int, to interface{}) bool
}

// kmp is a structure for Knuth–Morris–Pratt algorithm
type kmp struct {
	pattern interfaceKMP
	prefix  []int
	size    int
}

// FindPatternIndex returns index of first occurrences of pattern in argument and -1 otherwise
func (kmp *kmp) FindPatternIndex(s interfaceKMP) int {
	length := s.Len()

	if length < kmp.size {
		return -1
	}

	m := 0
	i := 0

	for m+i < length {
		if kmp.pattern.EqualTo(i, s.At(m+i)) {
			if i == kmp.size-1 {
				return m
			}

			i++
		} else {
			m = m + i - kmp.prefix[i]

			if kmp.prefix[i] > -1 {
				i = kmp.prefix[i]
			} else {
				i = 0
			}
		}
	}

	return -1
}

// ContainedIn returns true if pattern matched at least once in argument
func (kmp *kmp) ContainedIn(s interfaceKMP) bool {
	return kmp.FindPatternIndex(s) >= 0
}

// New creates instance for provided pattern where pattern has to be a slice
func New(pattern interfaceKMP) (*kmp, error) {
	prefix, err := computePrefix(pattern)

	if err != nil {
		return nil, err
	}

	return &kmp{
		pattern: pattern,
		prefix:  prefix,
		size:    pattern.Len(),
	}, nil
}

// computePrefix calculates array containing indexes of matches and returns error if pattern is less than one char
func computePrefix(pattern interfaceKMP) ([]int, error) {
	length := pattern.Len()

	if length < 2 {
		if length == 0 {
			return nil, errors.New("pattern must contain at least one element")
		}

		return []int{-1}, nil
	}

	shifts := make([]int, length)
	shifts[0] = -1
	shifts[1] = 0

	for pos, count := 2, 0; pos < length; {
		if pattern.EqualTo(pos-1, pattern.At(count)) {
			count++
			shifts[pos] = count
			pos++
		} else {
			if count > 0 {
				count = shifts[count]
			} else {
				shifts[pos] = 0
				pos++
			}
		}
	}

	return shifts, nil
}
