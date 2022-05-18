package forms

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
)

type Forms struct {
	rules, help form.Button
}

func (r Forms) Submit(submitter form.Submitter, pressed form.Button) {
	p := submitter.(*player.Player)
	switch pressed.Text {
	case r.rules.Text:
		p.Message("Rules")
		break
	case r.help.Text:
		p.Message("help")
		break
	}
}

func SendForms(p *player.Player) {
	f := form.NewMenu(form.MenuSubmittable(Forms{
		rules: form.Button{
			Text: "Rules!",
		},
		help: form.Button{
			Text: "Help!",
		},
	}), "Forms")
	p.SendForm(f)
}
