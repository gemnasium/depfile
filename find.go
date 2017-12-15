package depfile

import (
	"path/filepath"
	"regexp"
)

var regexps map[string]*regexp.Regexp

func Find(path string) *DepFile {
	fname := filepath.Base(path)
	for _, f := range DepFiles {
		for _, fn := range f.Filenames {
			if fname == fn {
				return &f
			}
		}
		for _, p := range f.Patterns {
			if r, ok := regexps[p]; ok && r.MatchString(fname) {
				return &f
			}
		}
	}
	return nil
}

func init() {
	regexps = make(map[string]*regexp.Regexp)
	for _, f := range DepFiles {
		for _, p := range f.Patterns {
			regexps[p] = regexp.MustCompile(p)
		}
	}
}
