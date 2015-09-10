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
	"github.com/nsf/termbox-go"
)

type View interface {
	// SetCell writes a single Unicode rune to a given x/y coordinate,
	// using the specified Style.
	SetCell(x int, y int, ch rune, style Style)

	// Clear clears the visible view, by filling it with spaces using
	// the specified Style.
	Clear(style Style)

	// Size represents the visible size.  The actual content may be
	// larger or smaller.
	Size() (int, int)
}

type screen struct{}

// Screen represents the default screen.
var Screen = &screen{}

func (*screen) SetCell(x, y int, ch rune, s Style) {
	fg, bg := s.attrs()
	termbox.SetCell(x, y, ch, fg, bg)
}
func (*screen) Clear(s Style) {
	fg, bg := s.attrs()
	termbox.Clear(fg, bg)
}
func (*screen) Size() (int, int) {
	return termbox.Size()
}

type ViewPort struct {
	physx  int // Anchor to the real world, usually 0
	physy  int // Again, anchor to the real world, usually 3
	viewx  int // Logical offset of the view
	viewy  int // Logical offset of the view
	limx   int // Content limits -- can't right scroll past this
	limy   int // Content limits -- can't down scroll past this
	width  int // View width
	height int // View height
	v      View
}

func (v *ViewPort) Clear(s Style) {
	for y := 0; y < v.height; y++ {
		for x := 0; x < v.width; x++ {
			v.v.SetCell(x+v.physx, y+v.physy, ' ', s)
		}
	}
}

func (v *ViewPort) Size() (int, int) {
	return v.width, v.height
}

// Reset resets the record of content, and also resets the offset back
// to the origin.  It doesn't alter the dimensions of the view port, nor
// the physical location relative to its parent.
func (v *ViewPort) Reset() {
	v.limx = 0
	v.limy = 0
	v.viewx = 0
	v.viewy = 0
}
func (v *ViewPort) SetCell(x, y int, ch rune, s Style) {
	if x > v.limx {
		v.limx = x
	}
	if y > v.limy {
		v.limy = y
	}
	if x < v.viewx || y < v.viewy {
		return
	}
	if x >= (v.viewx + v.width) {
		return
	}
	if y >= (v.viewy + v.height) {
		return
	}
	v.v.SetCell(x-v.viewx+v.physx, y-v.viewy+v.physy, ch, s)
}

// This moves the ViewPort the minimum necessary to make the given
// point visible.
func (v *ViewPort) MakeVisible(x, y int) {
	if x < 0 || y < 0 {
		return
	}
	if x >= v.limx || y > v.limy {
		return
	}
	if x >= v.viewx+v.width {
		v.viewx = x - (v.width - 1)
	}
	if x < v.viewx {
		v.viewx = x
	}
	if y >= v.viewy+v.height {
		v.viewy = y - (v.height - 1)
	}
	if y < v.viewy {
		v.viewy = y
	}
	v.ValidateView()
}

// ValidateViewY ensures that the Y offset of the view port is limited so that
// it cannot scroll away from the content.
func (v *ViewPort) ValidateViewY() {
	if v.viewy >= v.limy-v.height {
		v.viewy = (v.limy - v.height)
	}
	if v.viewy < 0 {
		v.viewy = 0
	}
}

// ValidateViewX ensures that the X offset of the view port is limited so that
// it cannot scroll away from the content.
func (v *ViewPort) ValidateViewX() {
	if v.viewx >= v.limx-v.width {
		v.viewx = (v.limx - v.width)
	}
	if v.viewx < 0 {
		v.viewx = 0
	}
}

// ValidateView does both ValidateViewX and ValidateViewY, ensuring both
// offsets are valid.
func (v *ViewPort) ValidateView() {
	v.ValidateViewX()
	v.ValidateViewY()
}

// This centers the point, if possible, in the view.
func (v *ViewPort) Center(x, y int) {
	if x < 0 || y < 0 || x > v.limx || y > v.limy {
		return
	}
	v.viewx = x - (v.width / 2)
	v.viewy = y - (v.height / 2)
	v.ValidateView()
}

func (v *ViewPort) ScrollUp(rows int) {
	v.viewy -= rows
	v.ValidateViewY()
}

func (v *ViewPort) ScrollDown(rows int) {
	v.viewy += rows
	v.ValidateViewY()
}

func (v *ViewPort) ScrollLeft(cols int) {
	v.viewx -= cols
	v.ValidateViewX()
}

func (v *ViewPort) ScrollRight(cols int) {
	v.viewx += cols
	v.ValidateViewX()
}

func (v *ViewPort) SetSize(width, height int) {
	v.height = height
	v.width = width
	v.ValidateView()
}

func (v *ViewPort) SetContentSize(width, height int) {
	v.limx = width
	v.limy = height
	v.ValidateView()
}

func (v *ViewPort) GetContentSize() (int, int) {
	return v.limx, v.limy
}

// Resize is called with the new dimensions, and also the new location in the
// the parent.
func (v *ViewPort) Resize(x, y, width, height int) {
	if v.v == nil {
		return
	}
	px, py := v.v.Size()
	if x >= 0 && x < px {
		v.physx = x
	}
	if y >= 0 && y < py {
		v.physy = y
	}
	if width < 0 {
		width = px - x
	}
	if height < 0 {
		height = py - y
	}
	if width <= x+px {
		v.width = width
	}
	if height <= y+py {
		v.height = height
	}
}

func (v *ViewPort) SetView(view View) {
	v.v = view
}

// Return a new view.  The x and y coordinates are an offset
// relative to the parent.  0,0 represents the upper left.
// The width and height indicate a width and height. If the value
// -1 is supplied, then the dimension is calculated from the parent.
func NewViewPort(view View, x, y, width, height int) *ViewPort {
	v := &ViewPort{v: view}
	// initial (and possibly poor) assumptions -- all visible
	// cells are addressible, but none beyond that.
	v.limx = width
	v.limy = height
	v.Resize(x, y, width, height)
	return v
}
