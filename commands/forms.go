package commands

import (
	"FirstGo/forms"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

type Forms struct {
}

func (f Forms) Run(src cmd.Source, o *cmd.Output) {
	p := src.(*player.Player)

	forms.SendForms(p)
}
