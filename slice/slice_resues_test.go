/**
 * Package slice
 * @Author iFurySt <ifuryst@gmail.com>
 * @Date 2023/12/12
 */

package slice

import (
	"github.com/google/uuid"
	"github.com/ifuryst/lol"
	"testing"
)

func BenchmarkSliceNotReuse(b *testing.B) {
	testMap := make(map[int]string, 20000)
	for i := 20000; i > 0; i-- {
		testMap[i] = uuid.New().String()
	}

	for i := 0; i < b.N; i++ {
		keys := lol.Keys(testMap)
		_ = keys
	}
}

func BenchmarkSliceReuse(b *testing.B) {
	testMap := make(map[int]string, 20000)
	for i := 20000; i > 0; i-- {
		testMap[i] = uuid.New().String()
	}

	keys := make([]int, 0, 20000)

	for i := 0; i < b.N; i++ {
		keys = keys[:0]
		for k := range testMap {
			keys = append(keys, k)
		}
		_ = keys
	}
}
