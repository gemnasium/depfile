package depfile

import (
	"path"
	"testing"
)

func TestFind(t *testing.T) {

	var check = func(query string, want string, got *DepFile) {
		if got == nil {
			t.Errorf("Expected %s to resolve to %s but got nil", query, want)
		} else if got.Name != want {
			t.Errorf("Expected %s to resolve to %s but got %v", query, want, got.Name)
		}
	}

	var fnames = []string{
		"requirements.txt",
		"package.json",
		"package-lock.json",
		"npm-shrinkwrap.json",
		"composer.json",
		"composer.lock",
		"Gemfile",
		"Gemfile.lock",
		"bower.json",
		"yarn.lock",
		"pom.xml",
	}

	t.Run("Canonical", func(t *testing.T) {
		for _, fname := range fnames {
			got := Find(fname)
			check(fname, fname, got)
		}
	})

	t.Run("Subdir", func(t *testing.T) {
		for _, fname := range fnames {
			q := path.Join("sub/dir", fname)
			got := Find(q)
			check(q, fname, got)
		}
	})

	t.Run("Aliases", func(t *testing.T) {
		var tcs = []struct {
			Queries []string
			Result  string
		}{
			{
				Queries: []string{
					"requirements.pip",
					"requires.txt",
					"requirements-dev.txt",
				},
				Result: "requirements.txt",
			},
			{
				Queries: []string{
					"xyz.gemspec",
					"activerecord.gemspec",
				},
				Result: "gemspec",
			},
			{
				Queries: []string{
					"gems.rb",
				},
				Result: "Gemfile",
			},
			{
				Queries: []string{
					"gems.locked",
				},
				Result: "Gemfile.lock",
			},
		}

		for _, tc := range tcs {
			for _, q := range tc.Queries {
				got := Find(q)
				check(q, tc.Result, got)
			}
		}
	})
}
