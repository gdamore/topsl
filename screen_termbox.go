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
	"sync"
)

type screen struct {
	sync.Mutex
}

type Screen interface {
	PollEvent() Event
	Init() error
	Fini()
	Sync()
	Flush()
	Interrupt()
	View
	sync.Locker
}

// Screen represents the default screen.
var DefaultScreen = &screen{}

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

func (*screen) Interrupt() {
	termbox.Interrupt()
}

func (*screen) Sync() {
	termbox.Sync()
}

func (*screen) Fini() {
	termbox.Close()
}

func (*screen) Init() error {
	return termbox.Init()
}

func (*screen) Flush() {
	termbox.Flush()
}

func (*screen) PollEvent() Event {
	ev := termbox.PollEvent()
	switch ev.Type {
	case termbox.EventResize:
		return &ResizeEvent{}
	case termbox.EventKey:
		return &KeyEvent{Ch: ev.Ch, Key: KeyCode(ev.Key)}
	}
	return nil
}

func GetScreen() Screen {
	return DefaultScreen
}
