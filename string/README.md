# string
Contains benchmarks for string related cases.

Reference results:
```
goos: darwin
goarch: arm64
pkg: github.com/ifuryst/go-benchmark/string
cpu: Apple M3 Pro
BenchmarkFindBeforeReplace_Exist-12              8852304               129.0 ns/op            48 B/op          2 allocs/op
BenchmarkReplace_Exist-12                        9867271               118.9 ns/op            48 B/op          2 allocs/op
BenchmarkFindBeforeReplace_NotExist-12          29248496                39.89 ns/op           16 B/op          1 allocs/op
BenchmarkReplace_NotExist-12                    26392646                45.11 ns/op           16 B/op          1 allocs/op
PASS
ok      github.com/ifuryst/go-benchmark/string  5.448s
```

## manipulation_test
Comparing the performance of finding a substring in a string before replacing it and directly replacing it.
