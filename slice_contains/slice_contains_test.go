package main

import (
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkSliceContains(b *testing.B) {
	b.Run("SliceContainsV1", func(b *testing.B) {
		b.ReportAllocs()
		
		b.StopTimer()
		s, target := getData()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			SliceContainsV1(s, target)
		}

	})

	b.Run("SliceContainsV0", func(b *testing.B) {
		b.ReportAllocs()

		b.StopTimer()
		s, target := getData()
		b.StartTimer()

		for i := 0; i < b.N; i++ {
			SliceContainsV0(s, target)
		}
	})
}

func TestSliceContains(t *testing.T) {
	t.Parallel()

	t.Run("latest match", func(t *testing.T) {
		t.Parallel()
		s, target := getData()
		require.Equal(t, target, s[len(s)-1])

		// always last match
		require.False(t, slices.Contains(s[:len(s)-1], target))

		simd := SliceContainsV0(s, target)
		simple := SliceContainsV1(s, target)

		require.True(t, simd)
		require.True(t, simple)
	})

	t.Run("middle match", func(t *testing.T) {
		t.Parallel()

		s, target := getData()
		s = s[:len(s)-1]
		s[len(s)/2] = target

		simd := SliceContainsV0(s, target)
		simple := SliceContainsV1(s, target)

		require.True(t, simd)
		require.True(t, simple)
	})

	t.Run("no match", func(t *testing.T) {
		t.Parallel()

		s, target := getData()
		s[len(s)-1]--

		simd := SliceContainsV0(s, target)
		simple := SliceContainsV1(s, target)

		require.False(t, simd)
		require.False(t, simple)
	})
}

// 16 alignment
func getData() ([]uint8, uint8) {
	s := make([]uint8, 1_000_000)

	for i := 0; i < len(s); i++ {
		s[i] = rand.N[uint8](5)
	}

	// always last match
	s[len(s)-1] = 10

	return s, 10
}
