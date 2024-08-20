package gui

import (
	"github.com/common-nighthawk/go-figure"
	"github.com/danvergara/gocui"
)

type BannerWidget struct {
	name           string
	x0, y0, x1, y1 float32
	color          gocui.Attribute
	label          string
}

func NewBannerWidget(name string, x0, y0, x1, y1 float32, color gocui.Attribute, label string) *BannerWidget {
	return &BannerWidget{name: name, x0: x0, y0: y0, x1: x1, y1: y1, color: color, label: label}
}

func (bw *BannerWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	x0, y0, x1, y1 := handleWidgetSize(maxX, maxY, int(bw.x0*float32(maxX)), int(bw.y0*float32(maxY)), int(bw.x1*float32(maxX)), int(bw.y1*float32(maxY)))

	v, err := g.SetView(bw.name, x0, y0, x1, y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FrameColor = bw.color
		nameFigure := figure.NewColorFigure(bw.label, "", "blue", true)
		figure.Write(v, nameFigure)
	}
	return nil
}

type InputWidget struct {
	name           string
	x0, x1, y0, y1 float32
	label          string
	color          gocui.Attribute
}

func NewInputWidget(name string, x0, y0, x1, y1 float32, label string, color gocui.Attribute) *InputWidget {
	return &InputWidget{
		name:  name,
		x0:    x0,
		y0:    y0,
		x1:    x1,
		y1:    y1,
		label: label,
		color: color,
	}
}

func (iw *InputWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	x0, y0, x1, y1 := handleWidgetSize(maxX, maxY, int(iw.x0*float32(maxX)), int(iw.y0*float32(maxY)), int(iw.x1*float32(maxX)), int(iw.y1*float32(maxY)))

	v, err := g.SetView(iw.name, x0, y0, x1, y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FrameColor = iw.color
		v.Title = iw.label
		v.TitleColor = iw.color
		v.Editable = true
		v.Wrap = true
		v.Highlight = true

		if _, err := g.SetCurrentView(iw.name); err != nil {
			return err
		}
	}
	return nil
}

type OutputWidget struct {
	name           string
	x0, x1, y0, y1 float32
	label          string
	color          gocui.Attribute
}

func NewOutputWidget(name string, x0, y0, x1, y1 float32, label string, color gocui.Attribute) *OutputWidget {
	return &OutputWidget{
		name:  name,
		x0:    x0,
		y0:    y0,
		x1:    x1,
		y1:    y1,
		label: label,
		color: color,
	}
}

func (ow *OutputWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	x0, y0, x1, y1 := handleWidgetSize(maxX, maxY, int(ow.x0*float32(maxX)), int(ow.y0*float32(maxY)), int(ow.x1*float32(maxX)), int(ow.y1*float32(maxY)))

	v, err := g.SetView(ow.name, x0, y0, x1, y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FrameColor = ow.color
		v.Title = ow.label
		v.TitleColor = ow.color
	}
	return nil
}

type TableWidget struct {
	name           string
	x0, x1, y0, y1 float32
	label          string
	color          gocui.Attribute
	output         map[string]string
}

func NewTableWidget(name string, x0, y0, x1, y1 float32, label string, color gocui.Attribute, output map[string]string) *TableWidget {
	return &TableWidget{
		name:   name,
		x0:     x0,
		y0:     y0,
		x1:     x1,
		y1:     y1,
		label:  label,
		color:  color,
		output: output,
	}
}

func (tw *TableWidget) Layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()

	x0, y0, x1, y1 := handleWidgetSize(maxX, maxY, int(tw.x0*float32(maxX)), int(tw.y0*float32(maxY)), int(tw.x1*float32(maxX)), int(tw.y1*float32(maxY)))

	v, err := g.SetView(tw.name, x0, y0, x1, y1)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}

		v.FrameColor = tw.color
		v.Title = tw.label
		v.TitleColor = tw.color

		renderTable(v, tw.output)
	}
	return nil
}

func handleWidgetSize(maxX, maxY int, x0, y0, x1, y1 int) (int, int, int, int) {
	if x1 >= maxX {
		x1 = maxX - 1
	}

	if y1 >= maxY {
		y1 = maxY - 1
	}

	if x0 >= x1 {
		x1 = x0 + 1
	}

	if y0 >= y1 {
		y1 = y0 + 1
	}

	return x0, y0, x1, y1
}
