package gui

import (
	"errors"
	"log"

	"github.com/danvergara/gocui"
	"github.com/vector-ops/memclient/client"
)

type Gui struct {
	g *gocui.Gui
	c *client.Client
}

func New(g *gocui.Gui, c *client.Client) *Gui {
	return &Gui{
		g: g,
		c: c,
	}
}

func (gui *Gui) prepare() error {
	gui.g.Highlight = true
	gui.g.Cursor = true
	gui.g.Mouse = true

	gui.setLayout()

	if err := gui.keybindings(); err != nil {
		return err
	}
	if err := gui.g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}

	return nil
}

func (gui *Gui) Run() error {
	if err := gui.prepare(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		return err
	}
	if err := gui.g.MainLoop(); err != nil && !errors.Is(err, gocui.ErrQuit) {
		return err
	}

	return nil
}

// Gui returns a pointer of a gocui.Gui instance.
func (gui *Gui) Gui() *gocui.Gui {
	return gui.g
}
