package gwh

import "strings"

// 解析路由
func parsePattern(pattern string) []string {
	portion := strings.Split(pattern, "/")
	parts := make([]string, 0)
	for _, part := range portion {
		if part != "" {
			parts = append(parts, part)
			// 只允许一个 * 通配符
			if part[0] == '*' {
				break
			}
		}
	}
	return parts
}
