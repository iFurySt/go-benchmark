# indexing
Contains benchmarks for indexing related cases.

Reference results:
```
goos: darwin
goarch: arm64
pkg: github.com/ifuryst/go-benchmark/indexing
cpu: Apple M3 Pro
BenchmarkIsValidSwitch-12               13969291                81.64 ns/op            0 B/op          0 allocs/op
BenchmarkIsValidMap-12                    531158              2259 ns/op               0 B/op          0 allocs/op
BenchmarkIsValidSwitchRandom-12         132003404                9.295 ns/op           0 B/op          0 allocs/op
BenchmarkIsValidMapRandom-12            26443974                44.69 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/ifuryst/go-benchmark/indexing        6.208s
```

# switch_vs_map_test
Contains benchmarks for comparing the performance between a `switch`-based implementation and a `map`-based implementation when indexing a value.
