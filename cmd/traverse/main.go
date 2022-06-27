package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/carlmjohnson/versioninfo"
	"github.com/joefitzgerald/traverse/command"
)

var cli struct {
	Find    command.Find `cmd:"" help:"Find a user"`
	Version VersionFlag  `name:"version" short:"v" help:"Print version information and quit"`
}

type VersionFlag string

func (v VersionFlag) Decode(ctx *kong.DecodeContext) error { return nil }
func (v VersionFlag) IsBool() bool                         { return true }
func (v VersionFlag) BeforeApply(app *kong.Kong, vars kong.Vars) error {
	version := vars["version"]
	if len(version) == 0 {
		version = "devel"
	}

	if len(vars["revision"]) > 0 {
		version += fmt.Sprintf(" → %s", vars["revision"])
	}

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
		kong.Vars{
			"version":  versioninfo.Version,
			"revision": versioninfo.Short(),
		})
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
