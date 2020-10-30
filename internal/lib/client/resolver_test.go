package client_test

import (
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/invit/ec2meta/internal/lib/client"
	"github.com/stretchr/testify/assert"
)

type Tree map[string]Tree

func TreeResolver(tree Tree, parts []string, index int) ([]string, bool) {
	out := []string{}

	if len(parts) == index+1 {
		if parts[index] == "" {
			// root node
			for k, v := range tree {
				if v != nil {
					k = k + "/"
				}

				out = append(out, k)
			}
		} else {
			// any other leaf
			if l, ok := tree[parts[index]]; ok {
				if l == nil {
					out = append(out, "")
				} else {
					for k, v := range tree[parts[index]] {
						if v != nil {
							k = k + "/"
						}

						out = append(out, k)
					}
				}

			} else {
				return nil, false
			}
		}

		for i, _ := range out {
			post := ""
			if strings.HasSuffix(out[i], "/") {
				post = "/"
			}

			out[i] = filepath.Join(append(parts[:index+1], out[i])...) + post
		}

		return out, true
	}

	return TreeResolver(tree[parts[index]], parts, index+1)
}

var tree = Tree{
	"a1": Tree{
		"1": Tree{
			"b1": Tree{
				"10": nil,
				"20": nil,
				"30": nil,
			},
			"b2": Tree{
				"10": nil,
				"20": nil,
				"30": nil,
			},
		},
	},
	"a2": Tree{
		"1": Tree{
			"b1": Tree{
				"10": nil,
				"20": nil,
				"30": nil,
			},
			"b2": Tree{
				"10": nil,
				"20": nil,
				"30": nil,
			},
		},
	},
}

func TestPathResolver(t *testing.T) {
	testMap := []struct {
		Path     string
		Expected []string
	}{
		{
			Path:     "/",
			Expected: []string{"a1/", "a2/"},
		},
		{
			Path:     "/xx",
			Expected: []string{},
		},
		{
			Path:     "/a1",
			Expected: []string{"a1/1/"},
		},
		{
			Path:     "/a1/1",
			Expected: []string{"a1/1/b1/", "a1/1/b2/"},
		},
		{
			Path:     "a1/1/b1/10",
			Expected: []string{"a1/1/b1/10"},
		},
		{
			Path:     "/a1/[*]",
			Expected: []string{"a1/1/b1/", "a1/1/b2/"},
		},
		{
			Path:     "/a1/[*]/b1",
			Expected: []string{"a1/1/b1/10", "a1/1/b1/20", "a1/1/b1/30"},
		},
		{
			Path:     "/a1/[*]/b1/10",
			Expected: []string{"a1/1/b1/10"},
		},
		{
			Path:     "/a1/[0]/b1/10",
			Expected: []string{"a1/1/b1/10"},
		},
		{
			Path: "/[*]/[*]/[*]/10",
			Expected: []string{
				"a1/1/b1/10",
				"a1/1/b2/10",
				"a2/1/b1/10",
				"a2/1/b2/10",
			},
		},
		{
			Path:     "/a1/1/[*]/10",
			Expected: []string{"a1/1/b1/10", "a1/1/b2/10"},
		},
		{
			Path:     "/a1/1/[*]/[0]",
			Expected: []string{"a1/1/b1/10", "a1/1/b2/10"},
		},
		{
			Path:     "/a1/[*]/xx",
			Expected: []string{},
		},
	}

	resolver := func(p string) ([]string, error) {
		path := strings.Split(strings.Trim(p, "/"), "/")
		r, _ := TreeResolver(tree, path, 0)
		sort.Strings(r)
		return r, nil
	}

	for _, m := range testMap {
		p, err := client.PathResolver(m.Path, resolver)
		assert.NoError(t, err)
		assert.ElementsMatch(t, m.Expected, p)
	}
}
