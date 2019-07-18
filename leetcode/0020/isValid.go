package leetcode

var data = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		_, ok := data[c]
		if ok == true {
			var top rune
			if len(stack) == 0 {
				top = '#'
			} else {
				top = stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
			}
			if top != data[c] {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}
	return len(stack) == 0
}
