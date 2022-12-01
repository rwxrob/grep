package grep

import (
	"fmt"

	"github.com/rwxrob/term"
)

type Result struct {
	File       string // file or path to the file (rel, abs, whatever)
	Text       string // sample of text containing the hit
	Beg        int    // byte of beginning within overall file
	End        int    // byte of ending within overall file
	TextBeg    int    // beginning within Text
	TextEnd    int    // ending within Text
	FileColor  string // color to use for Pretty File
	MatchColor string // color to use for Pretty Match
	ShowFile   bool   // display file in Pretty or not
}

// Pretty returns a colorized pretty version of the Result with padding
// around the match. Note that no other characters (including line
// returns and invisible characters) are escaped. That is left to the
// caller.
func (r Result) Pretty() string {
	var str string
	if r.ShowFile {
		if r.FileColor == "" {
			r.FileColor = DefFileColor
		}
		str = r.FileColor + r.File + `: ` + term.X
	}
	if len(r.MatchColor) > 0 {
		str += r.Text[:r.TextBeg] + r.MatchColor
		str += r.Text[r.TextBeg:r.TextEnd] + term.X
		str += r.Text[r.TextEnd:]
	} else {
		str += r.Text
	}
	return str
}

func (r Result) Plain() string {
	var str string
	if r.ShowFile {
		str += r.File + ": "
	}
	str += r.Text
	return str
}

func (r Result) String() string {
	if term.IsInteractive() {
		return r.Pretty()
	}
	return r.Plain()
}

var DefFileColor = term.Black
var DefMatchColor = term.Yellow

type Results struct {
	ShowFile   bool     // include file in marshaled output
	FileColor  string   // color for file name (see DefFileColor)
	MatchColor string   // color for the match in text (see DefMatchColor
	Hits       []Result // actual results
}

// String fulfills fmt.Stringer interface in different ways depending on
// if the term.IsInteractive of not. Use the specific MD and Pretty
// methods instead if needed.
func (r Results) String() string {
	var str string
	for _, rr := range r.Hits {
		str += rr.String()
	}
	return str
}

func (r Results) Pretty() string {
	var buf string
	if r.FileColor == "" {
		r.FileColor = DefFileColor
	}
	if r.MatchColor == "" {
		r.MatchColor = DefMatchColor
	}
	for _, i := range r.Hits {
		i.FileColor = r.FileColor
		i.MatchColor = r.MatchColor
		i.ShowFile = r.ShowFile
		buf += fmt.Sprintf(i.Pretty())
	}
	return buf
}
