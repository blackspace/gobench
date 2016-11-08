package gobench

import (
	"testing"
	"sync"
)

func BenchmarkCreateGoroutineWait(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg:=sync.WaitGroup{}
		wg.Add(10000)
		for n:=0;n<10000;n++ {
			go func() {
				wg.Wait()
			}()
		}
		wg.Add(-10000)
	}


}



func BenchmarkCreateGoroutineWaitEnd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wg_all_halt_point :=sync.WaitGroup{}
		wg_all_halt_point.Add(10000)

		wg_all_end :=sync.WaitGroup{}
		wg_all_end.Add(10000)
		for n:=0;n<10000;n++ {
			go func() {
				wg_all_halt_point.Wait()
				wg_all_end.Done()
			}()
		}

		wg_all_halt_point.Add(-10000)
		wg_all_end.Wait()
	}
}

