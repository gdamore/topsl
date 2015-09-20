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

// Attribute represents a single color, or additional attributes such as
// bold, reverse, underline.  For the moment, we are not supporting the entire
// 256 color palette.  We should explore that later.
type Attribute uint16

// Style incorporates the various elements of a style, both foreground and
// background.  Logically the foreground occupies the upper 16 bits, and the
// background the lower 16, but applications must not depend on that.
type Style uint32

func (s Style) Fg() Attribute {
	return Attribute(uint32(s) >> 16)
}

func (s Style) Bg() Attribute {
	return Attribute(uint32(s) & 0xffff)
}

func (s Style) Normal() Style {
	fg := s.Fg()
	bg := s.Bg()
	fg &= ^(AttrReverse | AttrBold | AttrUnderline)
	bg &= ^(AttrReverse | AttrBold | AttrUnderline)
	return NewStyle(fg, bg)
}

func (s Style) Bold() Style {
	return NewStyle(s.Fg()|AttrBold, s.Bg())
}
func (s Style) Blink() Style {
	return NewStyle(s.Fg(), s.Bg()|AttrBold)
}
func (s Style) Underline() Style {
	return NewStyle(s.Fg()|AttrUnderline, s.Bg()|AttrUnderline)
}
func (s Style) Reverse() Style {
	return NewStyle(s.Fg()|AttrReverse, s.Bg()|AttrReverse)
}

func NewStyle(fg, bg Attribute) Style {
	return Style((uint32(fg) << 16) | uint32(bg))
}

// Some stock styles
var (
	StyleDefault     = NewStyle(ColorWhite, ColorBlack)
	StyleText        = NewStyle(ColorWhite, ColorBlack)
	StyleTextBar     = NewStyle(ColorBlack, ColorWhite)
	StyleTitle       = NewStyle(ColorBlack, ColorWhite)
	StyleGood        = NewStyle(ColorGreen, ColorBlack)
	StyleError       = NewStyle(ColorRed, ColorBlack)
	StyleWarn        = NewStyle(ColorYellow, ColorBlack)
	StyleStatus      = NewStyle(ColorBlack, ColorWhite)
	StyleStatusError = NewStyle(ColorWhite|AttrBold, ColorRed)
	StyleStatusGood  = NewStyle(ColorWhite|AttrBold, ColorGreen)
	StyleStatusWarn  = NewStyle(ColorWhite|AttrBold, ColorYellow)
	StyleStatusDim   = NewStyle(ColorWhite|AttrBold, ColorWhite)
	StyleKey         = NewStyle(ColorBlack, ColorWhite)
	StyleKeyHot      = NewStyle(ColorBlue|AttrBold, ColorWhite)
)
