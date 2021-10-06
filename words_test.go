package profiling_demo

import (
	"os"
	"os/exec"
	"testing"

	"gotest.tools/v3/assert"
	"gotest.tools/v3/assert/cmp"
)

const testInput = "testdata/2701-0.txt"

func Benchmark_words(b *testing.B) {
	for i := 0; i < b.N; i++ {
		func() {
			f, err := os.Open(testInput)
			assert.Assert(b, err)
			b.Cleanup(func() {
				assert.Check(b, f.Close())
			})

			count, err := words(f)
			assert.Check(b, err)
			assert.Check(b, cmp.Equal(count, 181320))
		}()
	}
}

func Benchmark_native(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd := exec.Command("wc", "-w", testInput)
		out, err := cmd.Output()
		assert.Check(b, err)
		assert.Check(b, cmp.Contains(string(out), "215864"))
	}
}