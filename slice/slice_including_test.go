/**
 * Package slice
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2024/10/14
 */

package slice

import (
	"github.com/ifuryst/lol"
	"sort"
	"testing"
)

func Includes(strArray []string, target string) bool {
	sort.Strings(strArray)
	index := sort.SearchStrings(strArray, target)
	if index < len(strArray) && strArray[index] == target {
		return true
	}
	return false
}

func BenchmarkWithin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Includes([]string{"a", "b", "c"}, "a")
	}
}

func BenchmarkLol(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lol.Include([]string{"a", "b", "c"}, "a")
	}
}
