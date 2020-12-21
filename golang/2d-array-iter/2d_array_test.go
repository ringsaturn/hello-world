package arrayiter

import (
	"testing"
)

func benchmark2DArrayIter(b *testing.B) {
	sample := make([][]int, 1000)
	for _i := 0; _i < b.N; _i++ {
		func() int {
			s := 0
			for i := range sample {
				x := sample[i]
				for j := range x {
					s += x[j]
				}
			}
			return s
		}()
	}
}
