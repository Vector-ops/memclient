package gui

import "github.com/danvergara/gocui"

type keyBinding struct {
	view    string
	key     interface{}
	mod     gocui.Modifier
	handler func(*gocui.Gui, *gocui.View) error
}

func initialKeyBindings() []keyBinding {
	bindings := []keyBinding{
		{
			view:    "query",
			key:     gocui.KeyCtrlH,
			mod:     gocui.ModNone,
			handler: nextView("query", "result"),
		},
		{
			view:    "result",
			key:     gocui.KeyCtrlQ,
			mod:     gocui.ModNone,
			handler: nextView("result", "query"),
		},
	}

	for _, viewName := range []string{"query", "result"} {
		bindings = append(bindings, []keyBinding{
			{
				view:    viewName,
				key:     gocui.KeyArrowUp,
				mod:     gocui.ModNone,
				handler: moveCursorVertically("up"),
			},
			{
				view:    viewName,
				key:     gocui.KeyArrowDown,
				mod:     gocui.ModNone,
				handler: moveCursorVertically("down"),
			},
			{
				view:    viewName,
				key:     gocui.KeyArrowRight,
				mod:     gocui.ModNone,
				handler: moveCursorHorizontally("right"),
			},
			{
				view:    viewName,
				key:     gocui.KeyArrowLeft,
				mod:     gocui.ModNone,
				handler: moveCursorHorizontally("left"),
			},
		}...)
	}

	return bindings
}

func (gui *Gui) keybindings() error {
	for _, k := range initialKeyBindings() {
		if err := gui.g.SetKeybinding(k.view, k.key, k.mod, k.handler); err != nil {
			return err
		}
	}

	// SQL helpers
	if err := gui.g.SetKeybinding("query", gocui.KeyCtrlE, gocui.ModNone, gui.executeQuery()); err != nil {
		return err
	}

	return nil
}
