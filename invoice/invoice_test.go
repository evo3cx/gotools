package invoice

import "testing"

func TestGenerate(t *testing.T) {
	Generate()
}

func BenchmarkGenrate(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Generate()
		if err != nil {
			b.Fatal(err)
		}
	}
}
