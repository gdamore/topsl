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
}

// A CellView is a flexible view of a CellModel, offering both cursor
// management and a panning.
type CellView struct {
	port       *ViewPort
	view       View
	content    Widget
	contentV   *ViewPort
	hideCursor bool
	hasCursor  bool
	cursorX    int
	cursorY    int
	style      Style
	lines      []string
	model      CellModel
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
	if a.hasCursor {
		port.MakeVisible(a.cursorX, a.cursorY)
	}

	ex, ey := model.GetBounds()
	for y := 0; y < ey; y++ {
		for x := 0; x < ex; x++ {
			ch, style := model.GetCell(x, y)
			if ch == 0 {
				ch = ' '
				style = a.style
			}

			if a.hasCursor && x == a.cursorX && y == a.cursorY &&
				!a.hideCursor {
				style = style.Reverse()
			}
			port.SetCell(x, y, ch, style)
		}
	}
}

func (a *CellView) keyUp() {
	if !a.hasCursor {
		a.port.ScrollUp(1)
		return
	}
	if a.cursorY > 0 {
		a.cursorY--
	}
	a.model.SetCursor(a.cursorX, a.cursorY)
	a.port.MakeVisible(a.cursorX, a.cursorY)
}

func (a *CellView) keyDown() {
	if !a.hasCursor {
		a.port.ScrollDown(1)
		return
	}
	_, end := a.model.GetBounds()
	if a.cursorY < end-1 {
		a.cursorY++
	}
	a.model.SetCursor(a.cursorX, a.cursorY)
	a.port.MakeVisible(a.cursorX, a.cursorY)
}

func (a *CellView) keyLeft() {
	if !a.hasCursor {
		a.port.ScrollLeft(1)
		return
	}
	if a.cursorX > 0 {
		a.cursorX--
	}
	a.model.SetCursor(a.cursorX, a.cursorY)
	a.port.MakeVisible(a.cursorX, a.cursorY)
}

func (a *CellView) keyRight() {
	if !a.hasCursor {
		a.port.ScrollRight(1)
		return
	}
	end, _ := a.model.GetBounds()
	if a.cursorX < end-1 {
		a.cursorX++
	}
	a.model.SetCursor(a.cursorX, a.cursorY)
	a.port.MakeVisible(a.cursorX, a.cursorY)
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
		case 'J', 'j':
			a.keyDown()
			return true
		case 'K', 'k':
			a.keyUp()
			return true
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
}

func (a *CellView) HideCursor(on bool) {
	a.hideCursor = on
}

func (a *CellView) EnableCursor(on bool) {
	if on {
		if !a.hasCursor {
			a.cursorX = 0
			a.cursorY = 0
			a.port.MakeVisible(0, 0)
		}
	}
	a.hasCursor = on
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
