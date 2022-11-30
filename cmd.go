package grep

import (
	"fmt"
	"strconv"

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
	Name:      `grep`,
	Aliases:   []string{`bongrep`},
	Copyright: `Copyright 2021 Robert S Muhlestein`,
	Version:   `v0.1.1`,
	License:   `Apache-2.0`,
	Summary:   `async grep with regex`,
	Usage:     `(help|PATTERN)`,
	UseVars:   true,
	UseConf:   true,
	Commands:  []*Z.Cmd{help.Cmd},

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
		var padding = 20
		if len(args) == 0 {
			return help.Cmd.Call(x, args...)
		}
		// FIXME Get is hanging infinitely
		pad, err := x.Get(`padding`)
		if err != nil {
			return err
		}
		if pad != "" {
			padding, err = strconv.Atoi(pad)
			if err != nil {
				return err
			}
		}

		/*
			show, err := x.Get(`show-files`)
			if err != nil {
				return err
			}
		*/
		if len(args) == 1 {
			args = append(args, ".")
		}
		results, err := This(args[0], padding, args[1:]...)
		if err != nil {
			return err
		}
		results.ShowFile = true
		if term.IsInteractive() {
			fmt.Print(results.Pretty())
			return nil
		}
		fmt.Print(results)
		return nil
	},
}
