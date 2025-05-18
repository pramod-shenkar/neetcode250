package neetcode250

import (
	"testing"

	require "github.com/stretchr/testify/assert"
)

func Test78(t *testing.T) {
	got := subsets([]int{1, 2, 3})
	want := [][]int{
		{}, {1}, {2}, {3},
		{1, 2}, {2, 3}, {1, 3}, {1, 2, 3},
	}

	require.Len(t, got, len(want))
	for _, subset := range want {
		require.Contains(t, got, subset)
	}
}

func Test1863(t *testing.T) {
	got := subsetXORSum([]int{2, 4, 4})
	require.Equal(t, 24, got)
}

func Test39(t *testing.T) {
	got := combinationSum([]int{2, 3, 6, 7}, 7)
	want := [][]int{{2, 2, 3}, {7}}
	require.Equal(t, want, got)
}

func Test40(t *testing.T) {
	got := combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8)
	want := [][]int{{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6}}
	require.Equal(t, want, got)
}

func Test77(t *testing.T) {
	got := combine(4, 2)
	want := [][]int{{1, 2}, {1, 3}, {1, 4}, {2, 3}, {2, 4}, {3, 4}}
	require.Equal(t, want, got)
}

func Test46(t *testing.T) {
	got := permute([]int{5, 4, 6, 2})
	want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
	require.Equal(t, want, got)
}

func Test90(t *testing.T) {
	got := subsetsWithDup([]int{1, 2, 2})
	want := [][]int{{}, {1}, {1, 2}, {1, 2, 2}, {2}, {2, 2}}
	require.Equal(t, want, got)
}

func Test47(t *testing.T) {
	got := permuteUnique([]int{1, 1, 2})
	want := [][]int{{}, {1, 1, 2}, {1, 2, 1}, {2, 1, 1}}
	require.Equal(t, want, got)
}

func Test79WordSearch(t *testing.T) {
	// got := exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")
	// want := true
	// require.Equal(t, want, got)

	// got := exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE")
	// want := true
	// require.Equal(t, want, got)

	var cases = []struct {
		input1 [][]byte
		input2 string
		want   bool
	}{
		{
			[][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			"ABAB",
			false,
		},
	}

	for _, tt := range cases {
		require.Equal(t, tt.want, exist(tt.input1, tt.input2))
	}
}

func Test125(t *testing.T) {
	var cases = []struct {
		input string
		want  bool
	}{
		{
			"aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga",
			true,
		},
	}

	for _, tt := range cases {
		require.Equal(t, tt.want, validPalindrome(tt.input))
	}
}

func Test15(t *testing.T) {
	// require.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, threeSum([]int{-1, 0, 1, 2, -1, -4}))
	require.Equal(t, [][]int{{0, 0, 0}}, threeSum([]int{0, 0, 0}))
}

func Test11(t *testing.T) {
	require.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func Test189(t *testing.T) {
	var input = []int{-1, -100, 3, 99}
	rotate(input, 2)
	require.Equal(t, []int{3, 99, -1, -100}, input)
}

func Test2461(t *testing.T) {
	require.Equal(t, 3, maximumSubarraySum([]int{1, 2, 2}, 2))
}

func Test121(t *testing.T) {
	require.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))
}

func Test3(t *testing.T) {
	require.Equal(t, 3, lengthOfLongestSubstring("abcabcbb"))
}
