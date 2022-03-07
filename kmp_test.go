package kmp

import "testing"

type KMPInt []int

func (o KMPInt) At(i int) interface{} {
	return o[i]
}

func (o KMPInt) Len() int {
	return len(o)
}

func (o KMPInt) EqualTo(i int, to interface{}) bool {
	return o[i] == to.(int)
}

func TestKMPFindIntEmpty(t *testing.T) {
	var pattern = KMPInt{70}
	var text1 = KMPInt{70}
	var text2 = KMPInt{80}

	kmp, err := New(pattern)

	if err != nil {
		t.Error("No errors expected got", err)
	}

	index := kmp.FindPatternIndex(text1)
	result := 0

	if index != result {
		t.Error("Expected", result, "got", index)
	}

	index = kmp.FindPatternIndex(text2)
	result = -1

	if index != result {
		t.Error("Expected", result, "got", index)
	}
}

func TestKMPFindIntSmall(t *testing.T) {
	var pattern = KMPInt{70}
	var text1 = KMPInt{100, 90, 80, 70, 60, 5, 10, 15, 20, 10, 5}
	var text2 = KMPInt{100, 90, 80, 70, 60, 5, 11, 15, 20, 70, 5, 10}

	kmp, err := New(pattern)

	if err != nil {
		t.Error("No errors expected got", err)
	}

	index := kmp.FindPatternIndex(text1)
	result := 3

	if index != result {
		t.Error("Expected", result, "got", index)
	}

	index = kmp.FindPatternIndex(text2)
	result = 3

	if index != result {
		t.Error("Expected", result, "got", index)
	}
}

func TestKMPFindInt(t *testing.T) {
	var pattern = KMPInt{5, 10, 15}
	var text1 = KMPInt{100, 90, 80, 70, 60, 5, 10, 15, 20, 10, 5}
	var text2 = KMPInt{100, 90, 80, 70, 60, 5, 11, 15, 20, 10, 5, 10}

	kmp, err := New(pattern)

	if err != nil {
		t.Error("No errors expected got", err)
	}

	index := kmp.FindPatternIndex(text1)
	result := 5

	if index != result {
		t.Error("Expected", result, "got", index)
	}

	index = kmp.FindPatternIndex(text2)
	result = -1

	if index != result {
		t.Error("Expected", result, "got", index)
	}
}

func TestKMPContainInt(t *testing.T) {
	var pattern = KMPInt{20, 10, 5}
	var text1 = KMPInt{100, 90, 80, 70, 60, 5, 10, 15, 20, 10, 5}
	var text2 = KMPInt{100, 90, 80, 70, 60, 5, 11, 15, 20, 10, 55, 10}

	kmp, err := New(pattern)

	if err != nil {
		t.Error("No errors expected got", err)
	}

	contained := kmp.ContainedIn(text1)
	result := true

	if contained != result {
		t.Error("Expected", result, "got", contained)
	}

	contained = kmp.ContainedIn(text2)
	result = false

	if contained != result {
		t.Error("Expected", result, "got", contained)
	}
}
