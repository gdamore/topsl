// +build darwin freebsd linux netbsd openbsd windows

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
	"github.com/gdamore/tcell/termbox"
)

// Style incorporates the various elements of a style, both foreground and
// background.  Logically the foreground occupies the upper 16 bits, and the
// background the lower 16, but applications must not depend on that.

const (
	ColorDefault  Attribute = Attribute(termbox.ColorDefault)
	ColorBlack              = Attribute(termbox.ColorBlack)
	ColorWhite              = Attribute(termbox.ColorWhite)
	ColorRed                = Attribute(termbox.ColorRed)
	ColorCyan               = Attribute(termbox.ColorCyan)
	ColorGreen              = Attribute(termbox.ColorGreen)
	ColorBlue               = Attribute(termbox.ColorBlue)
	ColorYellow             = Attribute(termbox.ColorYellow)
	ColorMagenta            = Attribute(termbox.ColorMagenta)
	AttrBold                = Attribute(termbox.AttrBold)
	AttrReverse             = Attribute(termbox.AttrReverse)
	AttrUnderline           = Attribute(termbox.AttrUnderline)
)

func (s Style) attrs() (termbox.Attribute, termbox.Attribute) {
	fg := termbox.Attribute(uint32(s) >> 16)
	bg := termbox.Attribute(uint32(s) & 0xffff)
	return fg, bg
}
