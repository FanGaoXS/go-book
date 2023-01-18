package word

import (
	"fmt"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type Case struct {
		input    string
		expected bool
	}
	cases := []*Case{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, c := range cases {
		if actual := IsPalindrome(c.input); actual != c.expected {
			t.Errorf("IsPalindrome(%q) expected %v, actual %v", c.input, c.expected, actual)
		}
	}
}

func BenchmarkIsPalindrome(t *testing.B) {
	type Case struct {
		input    string
		expected bool
	}
	cases := []*Case{
		{"", true},
		{"a", true},
		{"aa", true},
		{"ab", false},
		{"kayak", true},
		{"detartrated", true},
		{"A man, a plan, a canal: Panama", true},
		{"Evil I did dwell; lewd did I live.", true},
		{"Able was I ere I saw Elba", true},
		{"été", true},
		{"Et se resservir, ivresse reste.", true},
		{"palindrome", false}, // non-palindrome
		{"desserts", false},   // semi-palindrome
	}

	for _, c := range cases {
		if actual := IsPalindrome(c.input); actual != c.expected {
			t.Errorf("IsPalindrome(%q) expected %v, actual %v", c.input, c.expected, actual)
		}
	}
}

func ExampleIsPalindrome() {
	fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
	fmt.Println(IsPalindrome("palindrome"))
	// Output:
	// true
	// false
}
