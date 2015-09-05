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

type TitleBar struct {
	b *TextBar
}

func (t *TitleBar) Draw() {
	t.b.Draw()
}

func (t *TitleBar) SetCenter(s string, style Style) {
	t.b.SetCenter(s, style)
}

func (t *TitleBar) SetLeft(s string, style Style) {
	t.b.SetLeft(s, style)
}

func (t *TitleBar) SetRight(s string, style Style) {
	t.b.SetRight(s, style)
}

func (t *TitleBar) SetView(view View) {
	t.b.SetView(view)
}

func (t *TitleBar) Resize() {
}

func (t *TitleBar) HandleEvent(Event) bool {
	return false
}

func NewTitleBar(center string) *TitleBar {
	t := &TitleBar{}
	t.b = NewTextBar()
	t.b.SetCenter(center, StyleTitle)
	return t
}
