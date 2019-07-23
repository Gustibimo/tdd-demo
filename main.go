package main

import (
	"fmt"
	"github.com/Gustibimo/tdd-stack/set"
)

func main() {
	s := set.New()

	s.Add("Peter")
	s.Add("David")
	s.Add("Davidsss")

	fmt.Println(s.Has("Peter"))  // True
	fmt.Println(s.Has("George")) // False

	fmt.Println(s)
	s.Remove("David")
	fmt.Println(s.Has("David")) // False
}
