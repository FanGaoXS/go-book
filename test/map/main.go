package main

func main() {

}

func equal(m1, m2 map[string]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k1, v1 := range m1 {
		if v2, ok := m2[k1]; !ok || v2 != v1 {
			return false
		}
	}

	return true
}
