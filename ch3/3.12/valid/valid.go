package valid

import "bytes"

func IsValid(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	var bs []byte
	for i := range s1 {
		bs = append(bs, s1[i])
	}

	for i := range s2 {
		index := bytes.IndexByte(bs, s2[i])
		if index == -1 {
			return false
		}
		bs = append(bs[:index], bs[index+1:]...)
	}

	return len(bs) == 0
}
