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
	Version:   `v0.2.3`,
	License:   `Apache-2.0`,
	Summary:   `async grep with regex`,
	Usage:     `(help|PATTERN [TARGET ...])`,
	UseVars:   true,
	UseConf:   true,
	Commands:  []*Z.Cmd{help.Cmd},

	Description: `
				The {{aka}} command is a simple utility similar to grep that can
				be composed into bonzai stateful command trees. It uses Go
				regular expressions to return all occurrences within a given
				file. Each TARGET is either a file or directory which will be
				recursively searched. If no TARGET is provided assumes the
				current directory.

			`,

	Call: func(x *Z.Cmd, args ...string) error {

		var padding = 20
		if len(args) == 0 {
			return help.Cmd.Call(x, args...)
		}

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
