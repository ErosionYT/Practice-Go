package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Feed struct {
}

func (f Feed) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)
	p.Saturate(20, 20)
	o.Print("You have fed yourself")
}
