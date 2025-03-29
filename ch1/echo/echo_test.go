// Exercise 1.3

/*
   Note: the println() functions are commented out because I want to benchmark
   the speed of each function, not the speed of writes to stdout. I might be
   better off using some king of nop in Go (not sure what the keyword is yet
   though).

   Results when run on my home machine:

   $ go test -bench=.
   goos: linux
   goarch: amd64
   pkg: github.com/pbrowne011/tgpl/ch1/echo
   cpu: Intel(R) Core(TM) i7-8565U CPU @ 1.80GHz
   BenchmarkEcho1-8           10000            217177 ns/op
   BenchmarkEcho2-8           10000            200898 ns/op
   BenchmarkEcho3-8        11907384                94.27 ns/op
   PASS
   ok      github.com/pbrowne011/tgpl/ch1/echo     5.424s

   Clearly the third is optimized. The other two execute in quadratic time and
   involve memory reassignment, which I think invokes the garbage collector,
   slowing things down even further.
*/

package echo

import (
	//	"fmt"
	"os"
	"strings"
	"testing"
)

func BenchmarkEcho1(b *testing.B) {
	var s, sep string
	for i := 0; i < b.N; i++ {
		for j := 1; j < len(os.Args); j++ {
			s += sep + os.Args[j]
			sep = " "
		}
		// fmt.Println(s)
	}
}

func BenchmarkEcho2(b *testing.B) {
	s, sep := "", ""
	for i := 0; i < b.N; i++ {
		for _, arg := range os.Args[1:] {
			s += sep + arg
			sep = " "
		}
		// fmt.Println(s)
	}
}

// Includes program name. Reassigns to string to prevent the compiler from
// optimizing out the code, with check at the end.
func BenchmarkEcho3(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		// fmt.Println()
		s = strings.Join(os.Args[0:], " ")
	}

	if len(s) < 1 {
		b.Fatalf("impossible")
	}
}
