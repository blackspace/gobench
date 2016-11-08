package gobench

import (
	"testing"
	"log"
)

func BenchmarkPrintln(b *testing.B) {
	for i:=0;i<b.N;i++ {
		log.Println("Hello World")
	}
}
