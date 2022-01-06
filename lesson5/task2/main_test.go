package main_test

import (
	"fmt"
	rset "task2/rset"
	set "task2/set"
	"testing"
)

var m = map[string]struct {
	add int
	has int
}{
	"10% запись, 90% чтение": {100, 900},
	"50% запись, 50% чтение": {500, 500},
	"90% запись, 10% чтение": {900, 100},
}

var parallelism = 1000

func Benchmark(b *testing.B) {
	var s = set.NewSet()
	var rs = rset.NewRSet()

	for k, v := range m {

		b.Run(fmt.Sprintf("%s ADD (add=%v has=%v)", k, v.add, v.has), func(b *testing.B) {
			b.SetParallelism(parallelism)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					for i := 0; i < v.add; i++ {
						s.Add(i)
					}
				}
			})
		})

		b.Run(fmt.Sprintf("%s HAS (add=%v has=%v)", k, v.add, v.has), func(b *testing.B) {
			b.SetParallelism(parallelism)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					for i := 0; i < v.has; i++ {
						s.Has(i)
					}
				}
			})
		})

		b.Run(fmt.Sprintf("%s RW ADD (add=%v has=%v)", k, v.add, v.has), func(b *testing.B) {
			b.SetParallelism(parallelism)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					for i := 0; i < v.add; i++ {
						rs.Add(i)
					}
				}
			})
		})

		b.Run(fmt.Sprintf("%s RW HAS (add=%v has=%v)", k, v.add, v.has), func(b *testing.B) {
			b.SetParallelism(parallelism)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					for i := 0; i < v.has; i++ {
						rs.Has(i)
					}
				}
			})
		})
	}
}

//Benchmark/10%_запись,_90%_чтение_ADD_(add=100_has=900)-8                           83512             14007 ns/op              11 B/op          0 allocs/op
//Benchmark/10%_запись,_90%_чтение_HAS_(add=100_has=900)-8                            8326            132393 ns/op             119 B/op          1 allocs/op
//Benchmark/10%_запись,_90%_чтение_RW_ADD_(add=100_has=900)-8                        76944             15565 ns/op              12 B/op          0 allocs/op
//Benchmark/10%_запись,_90%_чтение_RW_HAS_(add=100_has=900)-8                        14649             81122 ns/op              17 B/op          0 allocs/op

//Если кол-во операций чтения преобладает, то выгодно использовать вариант с sync.RWMutex

//Benchmark/50%_запись,_50%_чтение_ADD_(add=500_has=500)-8                           15415             76907 ns/op              64 B/op          1 allocs/op
//Benchmark/50%_запись,_50%_чтение_HAS_(add=500_has=500)-8                           16723             71015 ns/op              58 B/op          0 allocs/op
//Benchmark/50%_запись,_50%_чтение_RW_ADD_(add=500_has=500)-8                        13341             87314 ns/op              73 B/op          1 allocs/op
//Benchmark/50%_запись,_50%_чтение_RW_HAS_(add=500_has=500)-8                        26589             45703 ns/op               9 B/op          0 allocs/op

//Если кол-во операций чтения равно кол-ву записей, то также выгодно использовать вариант с sync.RWMutex

//Benchmark/90%_запись,_10%_чтение_ADD_(add=900_has=100)-8                            8624            140267 ns/op             112 B/op          1 allocs/op
//Benchmark/90%_запись,_10%_чтение_HAS_(add=900_has=100)-8                           96453             12423 ns/op              10 B/op          0 allocs/op
//Benchmark/90%_запись,_10%_чтение_RW_ADD_(add=900_has=100)-8                         7892            154669 ns/op             121 B/op          1 allocs/op
//Benchmark/90%_запись,_10%_чтение_RW_HAS_(add=900_has=100)-8                       129376              9059 ns/op               1 B/op          0 allocs/op

//Если кол-во операций записи преобладает, то выгодно использовать обычный sync.Mutex
//PASS
//ok
