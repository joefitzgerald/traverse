package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/joefitzgerald/traverse/command"
)

var version = "devel"

var cli struct {
	Find    command.Find `cmd:"" help:"Find a user"`
	Version VersionFlag  `name:"version" short:"v" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	fmt.Println(version)
	app.Exit(0)
	return nil
}

func main() {
	ctx := kong.Parse(&cli,
		kong.Name("traverse"),
		kong.Description("Build and maintain groups on the basis of reporting hierarchy"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
		}),
	)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
