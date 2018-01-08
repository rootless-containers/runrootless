package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransformRunCreate(t *testing.T) {
	newBundle := "/some/bundle/runrootless"
	cases := []struct {
		osArgs   []string
		expected []string
	}{
		{
			osArgs:   []string{"runc", "--root", "/foo", "run", "-b", "/bar", "baz"},
			expected: []string{"--root", "/foo", "run", "baz", "--bundle", newBundle},
		},
		{
			osArgs:   []string{"runc", "--root", "/foo", "run", "-b=/bar", "baz"},
			expected: []string{"--root", "/foo", "run", "baz", "--bundle", newBundle},
		},
		{
			osArgs:   []string{"runc", "--root", "/foo", "run", "baz"},
			expected: []string{"--root", "/foo", "run", "baz", "--bundle", newBundle},
		},
		{
			osArgs:   []string{"runc", "--root", "/foo", "list"},
			expected: []string{"--root", "/foo", "list"},
		},
	}

	for _, c := range cases {
		actual := _transformRunCreate(c.osArgs, newBundle)
		require.Equal(t, c.expected, actual)
	}
}
