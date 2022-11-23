package main

import "fmt"

func main() {
	a := []string{"s", "a", "a", "s", "d", "z", "a", "z", "v", "w", "w", "a", "a"}
	removeRepeatedStr(&a)
	fmt.Println(a)
}

func removeRepeatedStr(ptr *[]string) {
	s := *ptr
	for i := 0; i < len(s)-1; i++ {
		current := s[i]
		next := s[i+1]
		if current == next {
			if i == len(s)-2 {
				s = s[:i+1]
			} else {
				s = append(s[:i+1], s[i+2:]...)
			}
		}
	}
	*ptr = s
}
