package comma

import (
	"bytes"
	"strconv"
)

// Comma inserts commas in a non-negative decimal integer string.
func Comma(s string) string {
	if _, err := strconv.ParseInt(s, 10, 64); err != nil {
		return s
	}
	n := len(s)
	if n < 3 {
		return s
	}

	var bs []byte
	count := 0
	for i := n - 1; i >= 0; i-- {
		if count != 0 && count%3 == 0 {
			// insert ',' to head
			bs = append([]byte{','}, bs...)
		}
		b := s[i]
		// insert s[i] to head
		bs = append([]byte{b}, bs...)
		count++
	}

	buffer := bytes.Buffer{}
	for _, b := range bs {
		buffer.WriteByte(b)
	}
	return buffer.String()
}
