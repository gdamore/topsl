// Copyright 2015 The Tops'l Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use file except in compliance with the License.
// You may obtain a copy of the license at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"

	"github.com/gdamore/topsl"
)

type model struct {
	x      int
	y      int
	loc    string
	status *topsl.StatusBar
}

func (*model) GetBounds() (int, int) {
	return 60, 15
}

func (m *model) SetCursor(x int, y int) {
	m.x = x
	m.y = y
	m.loc = fmt.Sprintf("Cursor is %d,%d", m.x, m.y)
	m.status.SetStatus(m.loc)
}

func (m *model) GetCell(x, y int) (rune, topsl.Style) {
	dig := []rune{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
	var ch rune
	var style topsl.Style
	if x >= 60 || y >= 15 {
		return ch, style
	}
	colors := []topsl.Attribute{
		topsl.ColorWhite,
		topsl.ColorGreen,
		topsl.ColorRed,
		topsl.ColorBlue,
		topsl.ColorYellow}
	if y == 0 && x < len(m.loc) {
		style = topsl.StyleStatusGood
		ch = rune(m.loc[x])
	} else {
		ch = dig[(x)%len(dig)]
		style = topsl.NewStyle(colors[(y)%len(colors)],
			topsl.ColorBlack)
	}
	return ch, style
}

type MyApp struct {
	panel         topsl.Widget
	view          topsl.View
	main          *topsl.CellView
	keybar        *topsl.KeyBar
	cursorHide    bool
	cursorDisable bool
}

func (a *MyApp) HandleEvent(ev topsl.Event) bool {
	if a.panel != nil {
		if a.panel.HandleEvent(ev) {
			return true
		}

		switch ev := ev.(type) {
		case *topsl.KeyEvent:
			switch ev.Ch {
			case 0:
			case 'Q', 'q':
				topsl.AppFini()
				return true
			case 'S', 's':
				a.cursorHide = false
				a.main.HideCursor(a.cursorHide)
				a.updateKeys()
			case 'H', 'h':
				a.cursorHide = true
				a.main.HideCursor(a.cursorHide)
				a.updateKeys()
			case 'E', 'e':
				a.cursorDisable = false
				a.main.EnableCursor(!a.cursorDisable)
				a.updateKeys()
			case 'D', 'd':
				a.cursorDisable = true
				a.main.EnableCursor(!a.cursorDisable)
				a.updateKeys()
			}
		}
	}
	return false
}

func (a *MyApp) Draw() {
	if a.panel != nil {
		a.panel.Draw()
	}
}

func (a *MyApp) Resize() {
	if a.panel != nil {
		a.panel.Resize()
	}
}

func (a *MyApp) SetView(view topsl.View) {
	if a.panel != nil {
		a.panel.SetView(view)
	}
	a.view = view
}

func (a *MyApp) updateKeys() {
	w := []string{"_Quit"}
	if a.cursorDisable {
		w = append(w, "_Enable cursor")
	} else {
		w = append(w, "_Disable cursor")
		if !a.cursorHide {
			w = append(w, "_Hide cursor")
		} else {
			w = append(w, "_Show cursor")
		}
	}
	a.keybar.SetKeys(w)
}

func main() {

	app := &MyApp{}

	topsl.AppInit()
	defer topsl.AppFini()

	title := topsl.NewTitleBar("CellView Test")
	title.SetRight("Example v1.0", topsl.StyleTextBar)
	keyb := topsl.NewKeyBar([]string{"_Quit"})

	panel := topsl.NewPanel()

	panel.SetBottom(keyb)

	panel.SetTitle(title)

	status := topsl.NewStatusBar("Some thing happened")
	status.SetStatus("My status is here.")

	content := topsl.NewCellView()
	content.SetModel(&model{status: status})
	content.EnableCursor(true)

	panel.SetContent(content)
	panel.SetStatus(status)

	app.panel = panel
	app.main = content
	app.keybar = keyb
	app.updateKeys()
	topsl.SetApplication(app)
	topsl.RunApplication()
}
