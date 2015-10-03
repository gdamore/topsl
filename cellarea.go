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

package topsl

// CellModel models the content of a CellView.  The dimensions used within
// a CellModel are always logical, and have origin 0, 0.
type CellModel interface {
	GetCell(x, y int) (rune, Style)
	GetBounds() (int, int)
	SetCursor(int, int)
	GetCursor() (int, int, bool, bool)
	MoveCursor(offx, offy int)
}

// A CellView is a flexible view of a CellModel, offering both cursor
// management and a panning.
type CellView struct {
	port     *ViewPort
	view     View
	content  Widget
	contentV *ViewPort
	cursorX  int
	cursorY  int
	style    Style
	lines    []string
	model    CellModel
}

func (a *CellView) Draw() {

	port := a.port
	model := a.model
	port.Clear(a.style)
	if a.view == nil {
		return
	}
	if model == nil {
		return
	}
	ex, ey := model.GetBounds()
	vx, vy := port.Size()
	if ex < vx {
		ex = vx
	}
	if ey < vy {
		ey = vy
	}
	for y := 0; y < ey; y++ {
		for x := 0; x < ex; x++ {
			ch, style := model.GetCell(x, y)
			if ch == 0 {
				ch = ' '
				style = a.style
			}
			cx, cy, en, sh := a.model.GetCursor()
			if en && x == cx && y == cy && sh {
				style = style.Reverse()
			}
			port.SetCell(x, y, ch, style)
		}
	}
}

func (a *CellView) keyUp() {
	if _, _, en, _ := a.model.GetCursor(); !en {
		a.port.ScrollUp(1)
		return
	}
	a.model.MoveCursor(0, -1)
	a.MakeCursorVisible()
}

func (a *CellView) keyDown() {
	if _, _, en, _ := a.model.GetCursor(); !en {
		a.port.ScrollDown(1)
		return
	}
	a.model.MoveCursor(0, 1)
	a.MakeCursorVisible()
}

func (a *CellView) keyLeft() {
	if _, _, en, _ := a.model.GetCursor(); !en {
		a.port.ScrollLeft(1)
		return
	}
	a.model.MoveCursor(-1, 0)
	a.MakeCursorVisible()
}

func (a *CellView) keyRight() {
	if _, _, en, _ := a.model.GetCursor(); !en {
		a.port.ScrollRight(1)
		return
	}
	a.model.MoveCursor(+1, 0)
	a.MakeCursorVisible()
}

func (a *CellView) MakeCursorVisible() {
	if a.model == nil {
		return
	}
	x, y, enabled, _ := a.model.GetCursor()
	if enabled {
		a.MakeVisible(x, y)
	}
}

func (a *CellView) HandleEvent(e Event) bool {
	if a.model == nil {
		return false
	}
	switch e := e.(type) {
	case *KeyEvent:
		switch e.Ch {
		case 0:
			switch e.Key {
			case KeyUp, KeyCtrlP:
				a.keyUp()
				return true
			case KeyDown, KeyCtrlN:
				a.keyDown()
				return true
			case KeyRight, KeyCtrlF:
				a.keyRight()
				return true
			case KeyLeft, KeyCtrlB:
				a.keyLeft()
				return true
			}
		}
	}
	return false
}

func (a *CellView) SetModel(model CellModel) {
	w, h := model.GetBounds()
	model.SetCursor(0, 0)
	a.model = model
	a.port.SetContentSize(w, h)
	a.port.ValidateView()
}

func (a *CellView) SetView(view View) {
	port := a.port
	port.SetView(view)
	a.view = view
	if view == nil {
		return
	}
	width, height := view.Size()
	a.port.Resize(0, 0, width, height)
	if a.model != nil {
		w, h := a.model.GetBounds()
		a.port.SetContentSize(w, h)
	}
	a.Resize()
}

func (a *CellView) Resize() {
	// We might want to reflow text
	width, height := a.view.Size()
	a.port.Resize(0, 0, width, height)
	a.port.ValidateView()
	a.MakeCursorVisible()
}

func (a *CellView) SetCursor(x, y int) {
	a.cursorX = x
	a.cursorY = y
	a.model.SetCursor(x, y)
}

func (a *CellView) SetCursorX(x int) {
	a.SetCursor(x, a.cursorY)
}

func (a *CellView) SetCursorY(y int) {
	a.SetCursor(a.cursorX, y)
}

func (a *CellView) MakeVisible(x, y int) {
	a.port.MakeVisible(x, y)
}

func NewCellView() *CellView {
	return &CellView{
		port:  NewViewPort(nil, 0, 0, 0, 0),
		style: StyleDefault,
	}
}
