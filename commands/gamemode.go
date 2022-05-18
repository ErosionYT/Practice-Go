package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Gamemode struct {
	Number int
}

func (c Gamemode) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)
	u := "/gamemode <0, 1, 2>"
	switch c.Number {
	case 0:
		p.SetGameMode(world.GameModeSurvival)
		o.Print("You have set your gamemode to Survival")
	case 1:
		p.SetGameMode(world.GameModeCreative)
		o.Print("You have set your gamemode to Creative")
	case 2:
		p.SetGameMode(world.GameModeAdventure)
		o.Print("You have set your gamemode to Adventure")
	default:
		o.Print(u)
	}
}
