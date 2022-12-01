package grep_test

import (
	"fmt"

	"github.com/rwxrob/grep"
)

func ExampleThis_Pretty() {
	results, err := grep.This(`some`, 20, `testdata/afile`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", results.Pretty())
	results.ShowFile = true
	fmt.Printf("%q\n", results.Pretty())
	// Output:
	// "Here is a file with \x1b[33msome\x1b[0m stuff in it\\n\\n\\nand b\nore stuff here with \x1b[33msome\x1b[0mthing else\\nhere.\n"
	// "\x1b[30mtestdata/afile: \x1b[0mHere is a file with \x1b[33msome\x1b[0m stuff in it\\n\\n\\nand b\n\x1b[30mtestdata/afile: \x1b[0more stuff here with \x1b[33msome\x1b[0mthing else\\nhere.\n"

}

func ExampleThis_String() {
	results, err := grep.This(`some`, 20, `testdata/afile`)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%q\n", results.String())
	results.ShowFile = true
	fmt.Printf("%q\n", results.String())

	// Output:
	// "Here is a file with some stuff in it\\n\\n\\nand b\nore stuff here with something else\\nhere.\n"
	// "Here is a file with some stuff in it\\n\\n\\nand b\nore stuff here with something else\\nhere.\n"

}

func ExampleThis_hit() {
	results, err := grep.This(`advent`, 90, `testdata/advent.md`)
	if err != nil {
		fmt.Println(err)
	}
	hit := results.Hits[0]
	fmt.Printf("%v\n", hit.Text[hit.Beg:hit.End])
	fmt.Printf("%v\n", hit.Text[hit.TextBeg:hit.TextEnd])
	fmt.Println(hit.Beg == hit.TextBeg)
	fmt.Println(hit.End == hit.TextEnd)

	// Output:
	// advent
	// advent
	// true
	// true

}
