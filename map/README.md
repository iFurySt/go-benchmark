# map
Compare 1.24 and previous versions of go in terms of map performance. Because the map implementation has been optimized in go 1.24, the performance of the map has been improved. The performance of the map has been improved by 1.24 which implements map based on Swiss Tables.

Reference results:

## Benchmark
### 1.24
```
goos: linux
goarch: arm64
BenchmarkLookup-12              51288534                22.83 ns/op            0 B/op          0 allocs/op
BenchmarkInsertion-12           1000000000               0.02921 ns/op         0 B/op          0 allocs/op
BenchmarkDeletion-12            1000000000               1.116 ns/op           0 B/op          0 allocs/op
PASS
ok      command-line-arguments  3.493s
```

### 1.23
```
goos: linux
goarch: arm64
BenchmarkLookup-12              35707200                32.74 ns/op            0 B/op          0 allocs/op
BenchmarkInsertion-12           1000000000               0.03368 ns/op         0 B/op          0 allocs/op
BenchmarkDeletion-12            822345778                1.432 ns/op           0 B/op          0 allocs/op
PASS
ok      command-line-arguments  3.903s
```

## Direct Run
### 1.24
```
Lookup time: 257.288417ms
Insertion time: 67.836209ms
Deletion time: 63.128875ms
```

### 1.23
```
Lookup time: 350.119917ms
Insertion time: 172.39625ms
Deletion time: 45.574167ms
```