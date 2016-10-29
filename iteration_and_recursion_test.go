package gobench

import (
	"testing"
)

func Benchmark10000Iteration(b *testing.B) {
	for i:=0;i<b.N;i++ {
		for x:=0;x<10000;x++ {
		}
	}
}

var c10000=0
func R10000() int {
	if c10000==10000-1 {
		return c10000
	}
	c10000++
	return R10000()

}

func Benchmark10000Recursion(b *testing.B) {
	for i:=0;i<b.N;i++ {
		c10000=0
		v:=R10000()
		_=v
	}

}

func Benchmark100Iteration(b *testing.B) {
	for i:=0;i<b.N;i++ {
		for x:=0;x<100;x++ {
		}
	}
}

var c100=0
func R100() int {
	if c100==100-1 {
		return c100
	}
	c100++
	return R100()

}

func Benchmark100Recursion(b *testing.B) {
	for i:=0;i<b.N;i++ {
		c100=0
		v:=R100()
		_=v
	}

}

func Benchmark10Iteration(b *testing.B) {
	for i:=0;i<b.N;i++ {
		for x:=0;x<10;x++ {
		}
	}
}

var c10=0
func R10() int {
	if c10==10-1 {
		return c10
	}
	c10++
	return R10()

}

func Benchmark10Recursion(b *testing.B) {
	for i:=0;i<b.N;i++ {
		c10=0
		v:=R10()
		_=v
	}

}

func Benchmark1Iteration(b *testing.B) {
	for i:=0;i<b.N;i++ {
		for x:=0;x<1;x++ {
		}
	}
}



var c1=0
func R1() int {
	if c1==0 {
		return c1
	}
	c1++
	return R1()

}

func Benchmark1Recursion(b *testing.B) {
	for i:=0;i<b.N;i++ {
		c1=0
		v:=R10()
		_=v
	}

}


