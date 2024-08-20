package gui

import (
	"context"

	"github.com/danvergara/gocui"
)

func (gui *Gui) setLayout() error {
	bannerWidget := NewBannerWidget(
		"banner",
		0,
		0,
		0.36,
		0.14,
		gocui.ColorRed,
		"MEMCLIENT",
	)

	query := NewInputWidget(
		"query",
		0.37,
		0,
		1,
		0.25,
		"Query",
		gocui.ColorBlue,
	)

	log := NewOutputWidget(
		"log",
		0,
		0.16,
		0.36,
		1,
		"Log",
		gocui.ColorGreen,
	)

	data, err := gui.c.GetAll(context.Background())
	if err != nil {
		return err
	}

	result := NewTableWidget(
		"result",
		0.37,
		0.27,
		1,
		1,
		"Result",
		gocui.ColorYellow,
		data,
	)

	gui.g.SetManager(bannerWidget, query, log, result)

	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}
