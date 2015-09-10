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

type KeyBar struct {
	view     View
	style    Style
	hotstyle Style
	words    []string
}

func (k *KeyBar) Draw() {
	v := k.view
	if v == nil {
		return
	}
	v.Clear(k.style)
	x := 0
	y := 0
	for _, w := range k.words {
		hot := false
		for _, l := range w {
			if l == '_' {
				hot = true
				continue
			}
			if hot {
				v.SetCell(x, y, '[', k.style)
				x++
				v.SetCell(x, y, l, k.hotstyle)
				x++
				v.SetCell(x, y, ']', k.style)
				x++
			} else {
				v.SetCell(x, y, l, k.style)
				x++
			}
			hot = false
		}
		x += 2
	}
}

func (k *KeyBar) SetView(view View) {
	k.view = view
}

func (k *KeyBar) HandleEvent(Event) bool {
	return false
}

func (k *KeyBar) Resize() {
}

func (k *KeyBar) SetKeys(words []string) {
	k.words = words
}

func NewKeyBar(words []string) *KeyBar {
	k := &KeyBar{
		style:    StyleKey,
		hotstyle: StyleKeyHot,
		words:    words,
	}
	return k
}
