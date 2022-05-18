package commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type Spawn struct {
}

func (s Spawn) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)
	w := world.World{}
	w.AddEntity(p)

}
