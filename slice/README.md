# slice
Contains benchmarks for slice related cases.

Reference results:
```
goos: darwin
goarch: arm64
pkg: github.com/ifuryst/go-benchmark/slice
cpu: Apple M3 Pro
BenchmarkWithin-12              51610590                21.66 ns/op            0 B/op          0 allocs/op
BenchmarkLol-12                 601433917                2.009 ns/op           0 B/op          0 allocs/op
BenchmarkSliceNotReuse-12           6655            181671 ns/op          164169 B/op          7 allocs/op
BenchmarkSliceReuse-12              8010            148972 ns/op             293 B/op          4 allocs/op
PASS
ok      github.com/ifuryst/go-benchmark/slice   6.977s
```

## slice_including_test
Comparing the performance of finding an element in a slice using `for` loop and `sort.Search*` function.

## slice_reuses_test
Comparing the performance of reusing a slice and creating a new slice every time.
