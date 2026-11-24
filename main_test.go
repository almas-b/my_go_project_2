package main

import "testing"

func TestGenerateRandomElements(t *testing.T) {
	tests := []struct {
		name string
		size int
		want int
	}{
		{name: "zero size", size: 0, want: 0},
		{name: "negative size", size: -5, want: 0},
		{name: "one element", size: 1, want: 1},
		{name: "many elements", size: 1000, want: 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data := generateRandomElements(tt.size)

			if data == nil {
				t.Fatal("expected non-nil slice")
			}
			if len(data) != tt.want {
				t.Fatalf("expected length %d, got %d", tt.want, len(data))
			}
			for i, value := range data {
				if value <= 0 {
					t.Fatalf("element %d is not positive: %d", i, value)
				}
			}
		})
	}
}

func TestMaximum(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{42}, want: 42},
		{name: "max at start", data: []int{9, 1, 2, 3}, want: 9},
		{name: "max in middle", data: []int{1, 2, 9, 3}, want: 9},
		{name: "max at end", data: []int{1, 2, 3, 9}, want: 9},
		{name: "equal elements", data: []int{5, 5, 5}, want: 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.data); got != tt.want {
				t.Errorf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func TestMaxChunks(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{name: "nil slice", data: nil, want: 0},
		{name: "empty slice", data: []int{}, want: 0},
		{name: "one element", data: []int{7}, want: 7},
		{name: "less than chunks", data: []int{3, 8, 1}, want: 8},
		{name: "exactly chunks", data: []int{1, 2, 3, 4, 5, 6, 7, 8}, want: 8},
		{name: "max in remainder", data: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 100}, want: 100},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.data); got != tt.want {
				t.Errorf("expected %d, got %d", tt.want, got)
			}
		})
	}
}

func TestMaxChunksEqualsMaximum(t *testing.T) {
	data := generateRandomElements(100_003)

	if got, want := maxChunks(data), maximum(data); got != want {
		t.Errorf("expected %d, got %d", want, got)
	}
}
