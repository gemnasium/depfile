package depfile

import "testing"

func TestFind(t *testing.T) {

	t.Run("Canonical", func(t *testing.T) {
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

		for _, fname := range fnames {
			got := Find(fname)
			if got == nil {
				t.Errorf("Expected %s to resolve to %s but got nil", fname, fname)
			} else if got.Name != fname {
				t.Errorf("Expected %s to resolve to %s but got %v", fname, fname, got.Name)
			}
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
				if got == nil {
					t.Errorf("Expected %s to resolve to %s but got nil", q, tc.Result)
				} else if got.Name != tc.Result {
					t.Errorf("Expected %s to resolve to %s but got %v", q, tc.Result, got.Name)
				}
			}
		}
	})
}
