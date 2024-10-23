# hash
Contains benchmarks for hash related cases.

Reference results:
```
goos: darwin
goarch: arm64
cpu: Apple M3 Pro
BenchmarkMd5Hash-12              4911375               225.1 ns/op           128 B/op          4 allocs/op
BenchmarkMd5Hash2-12             8835208               134.8 ns/op            80 B/op          3 allocs/op
PASS
ok      command-line-arguments  3.063s
```

## md5_hash_test
Comparing the performance of different ways to hash a string using the `md5` algorithm.