package search

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFindUrl(t *testing.T) {
	cases := []struct {
		name    string
		logUrls []string
		target  string
		want    string
	}{
		{
			name: "empty",
		},
		{
			name:    "one found",
			logUrls: []string{"one"},
			target:  "one",
			want:    "one",
		},
		{
			name:    "one not found",
			logUrls: []string{"one"},
			target:  "two",
			want:    "",
		},
		{
			name:    "two",
			logUrls: []string{"a1", "b2"},
			target:  "a1",
			want:    "a1",
		},
		{
			name:    "two",
			logUrls: []string{"a1", "b2"},
			target:  "b2",
			want:    "b2",
		},
		{
			name:    "two a0",
			logUrls: []string{"a1", "b2"},
			target:  "a0",
			want:    "",
		},
		{
			name:    "two a2",
			logUrls: []string{"a1", "b2"},
			target:  "a2",
			want:    "",
		},
		{
			name:    "two b3",
			logUrls: []string{"a1", "b2"},
			target:  "b3",
			want:    "",
		},
		{
			name:    "three 1",
			logUrls: []string{"a1", "b2", "c3"},
			target:  "a1",
			want:    "a1",
		},
		{
			name:    "three 2",
			logUrls: []string{"a1", "b2", "c3"},
			target:  "b2",
			want:    "b2",
		},
		{
			name:    "three 3",
			logUrls: []string{"a1", "b2", "c3"},
			target:  "c3",
			want:    "c3",
		},
		{
			name:    "three not found",
			logUrls: []string{"a1", "b2", "c3"},
			target:  "nope",
			want:    "",
		},
		{
			name:    "three not found a0",
			logUrls: []string{"a1", "b2", "c3"},
			target:  "a0",
			want:    "",
		},
		{
			name:    "three not found a2",
			logUrls: []string{"a1", "b2", "c3"},
			target:  "a2",
			want:    "",
		},
		{
			name:    "three not found c0",
			logUrls: []string{"a1", "b2", "c3"},
			target:  "c0",
			want:    "",
		},
		{
			name:    "three not found c4",
			logUrls: []string{"a1", "b2", "c3"},
			target:  "c4",
			want:    "",
		},

		{
			name:    "four 1",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "a1",
			want:    "a1",
		},
		{
			name:    "four 2",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "b2",
			want:    "b2",
		},
		{
			name:    "four 3",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "c3",
			want:    "c3",
		},
		{
			name:    "four 4",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "d4",
			want:    "d4",
		},
		{
			name:    "four not found",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "nope",
			want:    "",
		},
		{
			name:    "four not a0",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "a0",
			want:    "",
		},
		{
			name:    "four not a2",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "a2",
			want:    "",
		},
		{
			name:    "four not d0",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "a0",
			want:    "",
		},
		{
			name:    "four not d5",
			logUrls: []string{"a1", "b2", "c3", "d4"},
			target:  "a2",
			want:    "",
		},
	}
	require := require.New(t)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			require.Equal(tc.want, findUrlSet(tc.logUrls, tc.target))
			require.Equal(tc.want, findUrlBinary(tc.logUrls, tc.target))
		})
	}
}

func TestFindUrls(t *testing.T) {
	fiveUrls := []string{"a1", "b2", "c3", "d4", "e5"}
	cases := []struct {
		name    string
		logUrls []string
		targets []string
		want    []string
	}{
		{
			name: "empty",
			want: []string{},
		},
		{
			name:    "one",
			logUrls: fiveUrls[:1],
			targets: []string{"a1"},
			want:    []string{"a1"},
		},
		{
			name:    "one",
			logUrls: fiveUrls[:1],
			targets: []string{"a1"},
			want:    []string{"a1"},
		},
		{
			name:    "one empty",
			logUrls: fiveUrls[:1],
			targets: []string{"x"},
			want:    []string{},
		},

		{
			name:    "five 1",
			logUrls: fiveUrls,
			targets: []string{"a1"},
			want:    []string{"a1"},
		},
		{
			name:    "five 3",
			logUrls: fiveUrls,
			targets: []string{"c3"},
			want:    []string{"c3"},
		},
		{
			name:    "five 5",
			logUrls: fiveUrls,
			targets: []string{"e5"},
			want:    []string{"e5"},
		},
		{
			name:    "five 5 2a",
			logUrls: fiveUrls,
			targets: []string{"a1", "e5"},
			want:    []string{"a1", "e5"},
		},
		{
			name:    "five 5 xx 2a",
			logUrls: fiveUrls,
			targets: []string{"a1", "xx", "e5"},
			want:    []string{"a1", "e5"},
		},
	}
	require := require.New(t)
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			require.Equal(tc.want, findUrlsSort(tc.logUrls, clone(tc.targets)))
			require.Equal(tc.want, findUrlsMap(tc.logUrls, tc.targets))
		})
	}
}

func clone[T any](s []T) []T {
	tmp := make([]T, len(s))
	copy(tmp, s)
	return tmp
}
