### kmp
Go version of Knuth–Morris–Pratt algorithm for custom user types, where user type must be
a slice implementing the following interface:
```go
type interfaceKMP interface {
	At(i int) interface{}
	Len() int
	EqualTo(i int, to interface{}) bool
}
```

### Example
```go
package main

import (
	"fmt"
	"strings"

	"github.com/vchezganov/kmp"
)

type StringList []string

func (l StringList) At(i int) interface{} {
	return l[i]
}

func (l StringList) Len() int {
	return len(l)
}

func (l StringList) EqualTo(i int, to interface{}) bool {
	return strings.EqualFold(l[i], to.(string))
}

func main() {
	pattern := StringList{
		"hello",
		"WORLD",
	}

	kmpSearch, err := kmp.New(pattern)
	if err != nil {
		panic(err)
	}

	list := StringList{
		"abc",
		"World",
		"hello",
		"hELLo",
		"world",
		"xyz",
	}

	firstIndex := kmpSearch.FindPatternIndex(list)
	fmt.Printf("Index: %d\n", firstIndex)
	fmt.Printf("List: %v\n", list[firstIndex:firstIndex+len(pattern)])
}
```

```
Index: 3
List: [hELLo world]
```