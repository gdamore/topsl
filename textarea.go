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

import (
	"strings"
)

type TextArea struct {
	port     *ViewPort
	view     View
	content  Widget
	contentV *ViewPort
	cursor   bool
	cursorX  int
	cursorY  int
	style    Style
	lines    []string
}

func (ta *TextArea) Draw() {
	if ta.view == nil {
		return
	}
	ta.view.Clear(ta.style)
	if ta.cursor {
		ta.port.MakeVisible(ta.cursorX, ta.cursorY)
	}
	didcursor := false
	y := 0
	for _, line := range ta.lines {
		x := 0
		for _, l := range line {
			style := ta.style
			if ta.cursor && ta.cursorY == y && ta.cursorX == x {
				style = style.Reverse()
				didcursor = true
			}
			ta.port.SetCell(x, y, l, style)
			x++
		}
		y++
	}
	if ta.cursor && !didcursor {
		ta.port.SetCell(ta.cursorX, ta.cursorY, ' ', ta.style.Reverse())
	}
}

func (ta *TextArea) keyUp() {
	if !ta.cursor {
		ta.port.ScrollUp(1)
		return
	}
	if ta.cursorY > 0 {
		ta.cursorY--
	}
	ta.port.MakeVisible(ta.cursorX, ta.cursorY)
}

func (ta *TextArea) keyDown() {
	if !ta.cursor {
		ta.port.ScrollDown(1)
		return
	}
	_, end := ta.port.GetContentSize()
	if ta.cursorY < end-1 {
		ta.cursorY++
	}
	ta.port.MakeVisible(ta.cursorX, ta.cursorY)
}

func (ta *TextArea) keyLeft() {
	if !ta.cursor {
		ta.port.ScrollLeft(1)
		return
	}
	if ta.cursorX > 0 {
		ta.cursorX--
	}
}

func (ta *TextArea) keyRight() {
	if !ta.cursor {
		ta.port.ScrollRight(1)
		return
	}
	end, _ := ta.port.GetContentSize()
	if ta.cursorX < end {
		ta.cursorX++
	}
	ta.port.MakeVisible(ta.cursorX, ta.cursorY)
}

func (ta *TextArea) HandleEvent(e Event) bool {
	switch e := e.(type) {
	case *KeyEvent:
		switch e.Ch {
		case 0:
			switch e.Key {
			case KeyUp, KeyCtrlP:
				ta.keyUp()
				return true
			case KeyDown, KeyCtrlN:
				ta.keyDown()
				return true
			case KeyRight, KeyCtrlF:
				ta.keyRight()
				return true
			case KeyLeft, KeyCtrlB:
				ta.keyLeft()
				return true
			}
		case 'J', 'j':
			ta.keyDown()
			return true
		case 'K', 'k':
			ta.keyUp()
			return true
		}
	}
	return false
}

func (ta *TextArea) contentSize() (int, int) {
	width := 0
	height := 0
	if lines := ta.lines; lines != nil {
		height = len(lines)
		width = 0
		for _, l := range lines {
			if len(l) > width {
				width = len(l)
			}
		}
	}
	return width, height
}

func (ta *TextArea) SetLines(lines []string) {
	ta.lines = append([]string{}, lines...)
	if ta.port != nil {
		ta.port.SetContentSize(ta.contentSize())
		ta.port.ValidateView()
	}
}

func (ta *TextArea) SetContent(text string) {
	lines := strings.Split(strings.Trim(text, "\n"), "\n")
	ta.SetLines(lines)
}

func (ta *TextArea) SetView(view View) {
	ta.view = view
	width, height := view.Size()
	ta.port = NewViewPort(view, 0, 0, width, height)
	ta.port.SetContentSize(ta.contentSize())
	ta.Resize()
}

func (ta *TextArea) Resize() {
	// We might want to reflow text
	width, height := ta.view.Size()
	ta.port.Resize(0, 0, width, height)
	ta.port.ValidateView()
}

func (ta *TextArea) SetCursorEnabled(on bool) {
	if on && ta.port != nil {
		if !ta.cursor {
			ta.cursorX = 0
			ta.cursorY = 0
			ta.port.MakeVisible(ta.cursorX, ta.cursorY)
		}
	}
	ta.cursor = on
}

func NewTextArea() *TextArea {
	return &TextArea{
		style: StyleDefault,
	}
}
