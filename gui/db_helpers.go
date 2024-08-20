package gui

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/danvergara/gocui"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func (gui *Gui) executeQuery() func(g *gocui.Gui, v *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		v.Rewind()

		ov, err := gui.g.View("result")
		if err != nil {
			return err
		}
		lv, err := gui.g.View("log")
		if err != nil {
			return err
		}

		ov.Clear()

		cmd := v.Buffer()

		switch {
		case strings.Contains(strings.ToLower(cmd), "get"):
			args := strings.Split(cmd, " ")
			if len(args) == 2 {
				val, err := gui.c.Get(context.Background(), strings.TrimSpace(args[1]))
				if err != nil {
					renderLog(lv, err)
				}
				if val == "key not found" {
					renderLog(lv, errors.New(val))
				} else {
					gui.render(v, "result", map[string]string{args[1]: val})
				}
			}
		case strings.Contains(strings.ToLower(cmd), "set"):
			args := strings.Split(cmd, " ")
			if len(args) == 3 {
				err := gui.c.Set(context.Background(), args[1], args[2])
				if err != nil {
					renderLog(lv, err)
				}
				renderLog(lv, "Key set")
				renderAllVals(gui, ov, lv)
			}
		case strings.Contains(strings.ToLower(cmd), "del"):
			args := strings.Split(cmd, " ")
			if len(args) == 2 {
				val, err := gui.c.Get(context.Background(), strings.TrimSpace(args[1]))
				if err != nil {
					renderLog(lv, err)
				}
				if val == "key not found" {
					renderLog(lv, errors.New(val))
				} else {
					renderLog(lv, "key deleted")
					renderAllVals(gui, ov, lv)
				}
			}
		case strings.Contains(strings.ToLower(cmd), "upd"):
			args := strings.Split(cmd, " ")
			if len(args) == 3 {
				err := gui.c.Set(context.Background(), args[1], args[2])
				if err != nil {
					renderLog(lv, err)
				}
				renderLog(lv, "Key updated")
				renderAllVals(gui, ov, lv)
			}
		}

		return nil
	}

}

func (gui *Gui) render(
	v *gocui.View,
	viewName string,
	resultSet map[string]string,
) error {
	v.Rewind()

	ov, err := gui.g.View(viewName)
	if err != nil {
		return err
	}

	// Cleans the view.
	ov.Clear()

	if err := ov.SetCursor(0, 0); err != nil {
		return err
	}

	renderTable(ov, resultSet)

	return nil
}

func renderTable(v *gocui.View, result map[string]string) {
	table := tablewriter.NewWriter(v)
	table.SetBorders(tablewriter.Border{Left: true, Top: false, Right: true, Bottom: true})
	table.SetHeader([]string{"key", "value"})
	// Add Bulk Data.
	var resultSet [][]string
	for k, v := range result {
		resultSet = append(resultSet, []string{k, v})
	}
	table.AppendBulk(resultSet)
	table.Render()
}

func renderLog(v *gocui.View, log interface{}) {
	l := fmt.Sprintf("%+v", time.Now().Format(time.RFC3339))
	switch log.(type) {
	case error:
		red := color.New(color.FgRed)
		boldRed := red.Add(color.Bold)
		fmt.Fprintf(v, "%s ", l)
		boldRed.Fprintf(v, "%s\n", log)
	case string:
		fmt.Fprintf(v, "%s %s\n", l, log)

	}
}

func renderAllVals(gui *Gui, ov *gocui.View, lv *gocui.View) {
	result, err := gui.c.GetAll(context.Background())
	if err != nil {
		renderLog(lv, err)
	} else {
		renderTable(ov, result)
	}
}
