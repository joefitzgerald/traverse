package main

import (
	"github.com/alecthomas/kong"
	"github.com/joefitzgerald/traverse/command"
)

var cli struct {
	Find command.Find `cmd:"" help:"Find a user"`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
