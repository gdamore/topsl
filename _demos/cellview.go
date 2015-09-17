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
	x    int
	y    int
	endx int
	endy int
	hide bool
	enab bool
	loc  string
}

func (m *model) GetBounds() (int, int) {
	return m.endx, m.endy
}

func (m *model) MoveCursor(offx, offy int) {
	m.x += offx
	m.y += offy
	m.limitCursor()
}

func (m *model) limitCursor() {
	if m.x < 0 {
		m.x = 0
	}
	if m.x > m.endx-1 {
		m.x = m.endx - 1
	}
	if m.y < 0 {
		m.y = 0
	}
	if m.y > m.endy-1 {
		m.y = m.endy - 1
	}
	m.loc = fmt.Sprintf("Cursor is %d,%d", m.x, m.y)
}

func (m *model) GetCursor() (int, int, bool, bool) {
	return m.x, m.y, m.enab, !m.hide
}

func (m *model) SetCursor(x int, y int) {
	m.x = x
	m.y = y

	m.limitCursor()
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
	panel  *topsl.Panel
	view   topsl.View
	main   *topsl.CellView
	keybar *topsl.KeyBar
	status *topsl.StatusBar
	model  *model
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
				a.model.hide = false
				a.updateKeys()
				return true
			case 'H', 'h':
				a.model.hide = true
				a.updateKeys()
				return true
			case 'E', 'e':
				a.model.enab = true
				a.updateKeys()
				return true
			case 'D', 'd':
				a.model.enab = false
				a.updateKeys()
				return true
			}
		}
	}
	return false
}

func (a *MyApp) Draw() {
	a.status.SetStatus(a.model.loc)
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
	m := a.model
	w := []string{"[Q] Quit"}
	if !m.enab {
		w = append(w, "[E] Enable cursor")
	} else {
		w = append(w, "[D] Disable cursor")
		if !m.hide {
			w = append(w, "[H] Hide cursor")
		} else {
			w = append(w, "[S] Show cursor")
		}
	}
	a.keybar.SetKeys(w)
}

func main() {

	app := &MyApp{}

	app.model = &model{endx: 60, endy: 15}

	topsl.AppInit()

	title := topsl.NewTitleBar()
	title.SetCenter("CellView Test")
	title.SetRight("Example v1.0")

	app.keybar = topsl.NewKeyBar()
	app.keybar.SetKeys([]string{"[Q] Quit"})

	app.status = topsl.NewStatusBar()
	app.status.SetStatus("My status is here.")

	app.main = topsl.NewCellView()
	app.main.SetModel(app.model)

	app.panel = topsl.NewPanel()
	app.panel.SetBottom(app.keybar)
	app.panel.SetTitle(title)
	app.panel.SetContent(app.main)
	app.panel.SetStatus(app.status)

	app.updateKeys()
	topsl.SetApplication(app)
	topsl.RunApplication()
}
