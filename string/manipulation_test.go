/**
 * Package string
 * @Author fuqiang.li <fuqiang.li@baishan.com>
 * @Date 2024/10/17
 */

package string

import (
	"strings"
	"testing"
)

func sanitize1(s string) string {
	s = strings.ToLower(s)
	if strings.Contains(s, "-") {
		return strings.ReplaceAll(s, "-", ":")
	}
	return s
}

func sanitize2(s string) string {
	return strings.ReplaceAll(strings.ToLower(s), "-", ":")
}

func BenchmarkFindBeforeReplace_Exist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize1("D4-57-63-EB-0F-01")
	}
}

func BenchmarkReplace_Exist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize2("D4-57-63-EB-0F-01")
	}
}

func BenchmarkFindBeforeReplace_NotExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize1("D45763EB0F01")
	}
}

func BenchmarkReplace_NotExist(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = sanitize2("D45763EB0F01")
	}
}
