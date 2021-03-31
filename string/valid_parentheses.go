package string


// leetcode 20
// https://leetcode.com/problems/valid-parentheses/

//First approach using stack without mapping
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, r := range s {
		switch r {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack = append(stack, r)
		case ')':
			if len(stack) > 0 && stack[len(stack) - 1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack) - 1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack) - 1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

//clean code using mapping brackets before string validation
func isValid2(s string) bool {
	brackets := map[rune]rune{'(':')', '[':']', '{':'}'}
	if len(s) == 0 {
		return true
	}
	stack := make([]rune, 0)
	for _, r := range s {
		switch r {
		case '(':
			fallthrough
		case '[':
			fallthrough
		case '{':
			stack = append(stack, r)
		case ')':
			fallthrough
		case ']':
			fallthrough
		case '}':
			if len(stack) > 0 && brackets[stack[len(stack) - 1]] == r {
				stack = stack[:len(stack) - 1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
