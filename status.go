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

type StatusBar struct {
	bar       *TextBar
	normStyle Style
	goodStyle Style
	failStyle Style
	warnStyle Style
	text      string
}

func (s *StatusBar) Draw() {
	s.bar.Draw()
}

func (s *StatusBar) HandleEvent(Event) bool {
	return false
}

func (s *StatusBar) SetView(view View) {
	s.bar.SetView(view)
}

func (s *StatusBar) SetNormal() {
	s.bar.SetStyle(s.normStyle)
	s.bar.SetLeft(s.text, s.normStyle)
}

func (s *StatusBar) SetFail() {
	s.bar.SetStyle(s.failStyle)
	s.bar.SetLeft(s.text, s.failStyle)
}

func (s *StatusBar) SetGood() {
	s.bar.SetStyle(s.goodStyle)
	s.bar.SetLeft(s.text, s.goodStyle)
}

func (s *StatusBar) SetWarn() {
	s.bar.SetStyle(s.warnStyle)
	s.bar.SetLeft(s.text, s.warnStyle)
}

func (s *StatusBar) SetStatus(text string) {
	s.text = text
	s.bar.SetLeft(text, StyleDefault)
}

func (s *StatusBar) Resize() {
}

func NewStatusBar() *StatusBar {
	s := &StatusBar{}
	s.normStyle = StyleStatus
	s.goodStyle = StyleStatusGood
	s.failStyle = StyleStatusError
	s.warnStyle = StyleStatusWarn

	s.bar = NewTextBar()
	s.bar.SetStyle(s.normStyle)
	return s
}
