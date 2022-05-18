package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Feed struct {
}

func (f Feed) Run(src cmd.Source, o *cmd.Output) {
	isPlayer := src.(*player.Player)
	isPlayer.SetFood(isPlayer.Food())
}
