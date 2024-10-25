```
❯ go test *.go -bench=. -benchmem -cpu=1,2,4,8
goos: darwin
goarch: arm64
cpu: Apple M3 Max
BenchmarkSHA256                 22642249                52.49 ns/op           32 B/op          1 allocs/op
BenchmarkSHA256-2               23770260                50.35 ns/op           32 B/op          1 allocs/op
BenchmarkSHA256-4               23524568                49.83 ns/op           32 B/op          1 allocs/op
BenchmarkSHA256-8               24034749                50.08 ns/op           32 B/op          1 allocs/op
BenchmarkSHA512                 10898224               109.4 ns/op            64 B/op          1 allocs/op
BenchmarkSHA512-2               11039143               108.1 ns/op            64 B/op          1 allocs/op
BenchmarkSHA512-4               11005012               110.2 ns/op            64 B/op          1 allocs/op
BenchmarkSHA512-8               10984533               113.8 ns/op            64 B/op          1 allocs/op
BenchmarkArgon2                       92          34532030 ns/op        67116456 B/op         48 allocs/op
BenchmarkArgon2-2                     63          17768788 ns/op        67116713 B/op         49 allocs/op
BenchmarkArgon2-4                    127           9463131 ns/op        67116504 B/op         48 allocs/op
BenchmarkArgon2-8                    126           9860136 ns/op        67116666 B/op         48 allocs/op
BenchmarkBcrypt                       21          51237710 ns/op            5152 B/op         10 allocs/op
BenchmarkBcrypt-2                     22          51062104 ns/op            5168 B/op         10 allocs/op
BenchmarkBcrypt-4                     22          50939195 ns/op            5180 B/op         10 allocs/op
BenchmarkBcrypt-8                     22          51056252 ns/op            5198 B/op         10 allocs/op
BenchmarkScrypt                       56          19291911 ns/op        16781808 B/op         22 allocs/op
BenchmarkScrypt-2                     60          19518917 ns/op        16781809 B/op         22 allocs/op
BenchmarkScrypt-4                     60          19518851 ns/op        16781841 B/op         22 allocs/op
BenchmarkScrypt-8                     54          19063583 ns/op        16781837 B/op         22 allocs/op
BenchmarkPBKDF2_SHA256                56          21384875 ns/op             772 B/op         10 allocs/op
BenchmarkPBKDF2_SHA256-2              55          21266664 ns/op             772 B/op         10 allocs/op
BenchmarkPBKDF2_SHA256-4              56          21178376 ns/op             772 B/op         10 allocs/op
BenchmarkPBKDF2_SHA256-8              56          21391156 ns/op             772 B/op         10 allocs/op
BenchmarkPBKDF2_SHA512                25          48327123 ns/op            1348 B/op         10 allocs/op
BenchmarkPBKDF2_SHA512-2              25          48202653 ns/op            1348 B/op         10 allocs/op
BenchmarkPBKDF2_SHA512-4              24          52405974 ns/op            1348 B/op         10 allocs/op
BenchmarkPBKDF2_SHA512-8              22          54515831 ns/op            1349 B/op         10 allocs/op
```
