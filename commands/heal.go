package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity/healing"
	"github.com/df-mc/dragonfly/server/player"
)

type Heal struct {
}

func (h Heal) Run(src cmd.Source, o *cmd.Output) {
	isPlayer := src.(*player.Player)
	isPlayer.Heal(isPlayer.MaxHealth(), healing.SourceCustom{})
	o.Print("You have healed yourself")
}
