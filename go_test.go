package gobench

import (
	"testing"
)

func BenchmarkIteration(b *testing.B) {
	for i:=0;i<b.N;i++ {
		for x:=0;x<10000;x++ {
			_=x+x
		}
	}
}

var c=0

func R() {
	if c==10000-1 {
		return
	}
	c++
	_=c+c
	R()

}

func BenchmarkRecursion(b *testing.B) {
	for i:=0;i<b.N;i++ {
		R()
	}

}


