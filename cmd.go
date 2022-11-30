package grep

import (
	"log"

	Z "github.com/rwxrob/bonzai/z"
	_ "github.com/rwxrob/conf"
	"github.com/rwxrob/help"
	"github.com/rwxrob/term"
	_ "github.com/rwxrob/vars"
)

func init() {
	Z.Vars.SoftInit()
	Z.Conf.SoftInit()
}

var Cmd = &Z.Cmd{
	Name:     `grep`,
	Aliases:  []string{`bongrep`},
	Summary:  `async grep with regex`,
	Usage:    `(help|PATTERN)`,
	UseVars:  true,
	UseConf:  true,
	MinArgs:  1,
	Commands: []*Z.Cmd{help.Cmd},

	Description: `
				The {{aka}} command is what you *actually* want. It transforms the
				command line and a collection of flat text files into the closest
				thing to a database that you can get.

				{{aka}} is a version of grep that can be composed into other
				composite stateful command branches with the benefits of being
				asynchronous so that every file searched gets its own goroutine.

				Plus {{aka}} provides the ability to set stateful filters for all
				searches that layer on top of one another. These predefined search
				filters can be saved and named and mapped based on the current
				working directory.

				The number of goroutines allocated to any search
				can also be allocated allowing you to turn up or down the amount of
				memory consumed for a given grep.

				Assume Current Directory

				Unlike traditional grep this command assumes you want to grep
				the current directory unless more than one argument is passed.

			`,

	Call: func(x *Z.Cmd, args ...string) error {
		pattern := args[0]
		targets := []string{}
		if len(args) > 1 {
			targets = args[1:]
		}
		// TODO get pad from vars
		pad := 10
		results, err := This(pattern, pad, targets...)
		if err != nil {
			return err
		}
		log.Print(results)

		show, err := x.Get(`show-files`)
		// TODO add text/template, with default template
		if err != nil {
			return err
		}
		log.Print(show)

		if term.IsInteractive() {
			//results.PrintPretty(tmpl)
			// TODO
			return nil
		}
		// TODO if terminal is interactive return colored results based on
		// variables and configuration
		// TODO if terminal is not interactive assume it is cat or vim and
		// provide a GitHub compatible markdown link to the line in the file with no color or other escaping.
		return nil
	},
}
