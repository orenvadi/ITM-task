package main

func testEq(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// MaskLinks masks links in the input string by replacing them with asterisks.
func MaskLinks(input string) string {
	data := []rune(input)
	linkPrefix := []rune("http://")

	isLink := false
	c := 6
	for i := 0; i < len(data)-6; i++ {

		if testEq(data[i:i+7], linkPrefix) {
			isLink = true
			continue
		}

		if isLink && c > 0 {
			c--
			continue
		}

		if data[i] == rune(' ') {
			isLink = false
		}

		if isLink {
			data[i] = '*'
		}

	}

	return string(data)
}
