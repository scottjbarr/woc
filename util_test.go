package woc

import (
	"fmt"
	"testing"
)

func TestChunkSlice(t *testing.T) {
	tests := []struct {
		name     string
		max      int
		elements int
		want     int
	}{
		{
			name:     "zero elements",
			max:      20,
			elements: 0,
			want:     0,
		},
		{
			name:     "one chunk",
			max:      5,
			elements: 5,
			want:     1,
		},
		{
			name:     "two chunks",
			max:      5,
			elements: 6,
			want:     2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ary := make([]string, tt.elements, tt.elements)

			fmt.Printf("ary = %+#v\n", ary)

			got := ChunkSlice(ary, tt.max)

			fmt.Printf("got = %+#v\n", got)

			gotLen := len(got)

			if gotLen != tt.want {
				t.Fatalf("got %v want %v", gotLen, tt.want)
			}
		})
	}
}
