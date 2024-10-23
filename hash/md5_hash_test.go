/**
 * Package hash
 * @Author fuqiang.li <fuqiang.li@baishan.com>
 * @Date 2024/10/23
 */

package hash

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func md5Hash(ss ...string) string {
	var buf bytes.Buffer
	for _, s := range ss {
		buf.WriteString(s)
	}
	return fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
}

func md5Hash2(ss ...string) string {
	hash := md5.New()
	for _, s := range ss {
		hash.Write([]byte(s))
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func BenchmarkMd5Hash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5Hash("hello", "world")
	}
}

func BenchmarkMd5Hash2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		md5Hash2("hello", "world")
	}
}
