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

type TextBar struct {
	view        View
	leftStyle   Style
	rightStyle  Style
	centerStyle Style
	padStyle    Style
	left        string
	right       string
	center      string
}

func (t *TextBar) Draw() {
	v := t.view
	if v == nil {
		return
	}
	width, _ := v.Size()
	x := 0
	y := 0
	if width == 0 {
		return
	}

	v.Clear(t.padStyle)

	// do left text
	for _, l := range t.left {
		v.SetCell(x, y, l, t.leftStyle)
		x++
	}
	// advance for center if there is space
	if start := (width - len(t.center)) / 2; start > x {
		x = start
	}

	// do center text
	for _, l := range t.center {
		v.SetCell(x, y, l, t.centerStyle)
		x++
	}

	// advance for right if there is space
	if start := width - len(t.right); start > x {
		x = start
	}

	// do right text
	for _, l := range t.right {
		v.SetCell(x, y, l, t.rightStyle)
		x++
	}
}

func (t *TextBar) SetView(view View) {
	t.view = view
}

func (t *TextBar) HandleEvent(Event) bool {
	return false
}

func (t *TextBar) SetCenter(s string, style Style) {
	t.center = s
	if style != StyleDefault {
		t.centerStyle = style
	}
}

func (t *TextBar) SetLeft(s string, style Style) {
	t.left = s
	if style != StyleDefault {
		t.leftStyle = style
	}
}

func (t *TextBar) SetRight(s string, style Style) {
	t.right = s
	if t.rightStyle != StyleDefault {
		t.rightStyle = style
	}
}

func (t *TextBar) SetStyle(style Style) {
	t.padStyle = style
}

func (t *TextBar) Resize() {
	// Nothing we can do.. move on.
}

func NewTextBar() *TextBar {
	t := &TextBar{
		left:        "",
		right:       "",
		center:      "",
		padStyle:    StyleTextBar,
		rightStyle:  StyleTextBar,
		centerStyle: StyleTextBar,
		leftStyle:   StyleTextBar,
	}
	return t
}
