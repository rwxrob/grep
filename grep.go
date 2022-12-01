package grep

import (
	"os"
	"regexp"

	"github.com/rwxrob/fs"
	"github.com/rwxrob/fs/dir"
)

var cache map[string]*regexp.Regexp

func cached(pattern string) (*regexp.Regexp, error) {
	var err error
	v, have := cache[pattern]
	if !have {
		v, err = regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
	}
	return v, nil
}

// This (as in grep.This) searches the targets for all instances of the
// regular expression pattern which is cached. Subsequent searches for
// the same patter will use the cached compiled regular expression for
// that pattern. Does not return an error if no results are found, only
// if something related to reading the targets involved.
func This(pattern string, pad int, targets ...string) (*Results, error) {
	re, err := cached(pattern)

	var results = Results{Hits: []Result{}}
	if err != nil {
		return nil, err
	}
	for _, target := range targets {

		// just a file
		if fs.NotExists(target) {
			continue
		}

		// recurse if target is a directory
		if fs.IsDir(target) {
			for _, entry := range dir.Entries(target) {
				res, err := This(pattern, pad, entry)
				if err != nil {
					return nil, err
				}
				results.Hits = append(results.Hits, res.Hits...)
			}
			return &results, nil
		}

		buf, err := os.ReadFile(target)
		if err != nil {
			return nil, err
		}

		for _, match := range re.FindAllIndex(buf, -1) {
			res := Result{
				Beg:  match[0],
				End:  match[1],
				File: target,
			}
			b := res.Beg - pad
			e := res.End + pad
			if b < 0 {
				b = 0
			}
			if e > len(buf)-1 {
				e = len(buf) - 1
			}
			res.TextBeg = pad
			res.TextEnd = res.TextBeg + (res.End - res.Beg)
			res.Text = string(buf[b:e])
			results.Hits = append(results.Hits, res)
		}

	}
	return &results, nil
}
