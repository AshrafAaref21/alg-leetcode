package p0003_container

import "testing"

func TestContainer(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "example",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "two bars",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "empty",
			height: []int{},
			want:   0,
		},
		{
			name:   "single bar",
			height: []int{1},
			want:   0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Container(tc.height)
			if got != tc.want {
				t.Fatalf("Container(%v) = %d, want %d", tc.height, got, tc.want)
			}
		})
	}
}
