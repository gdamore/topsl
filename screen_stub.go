// +build dragonfly nacl plan9 solaris

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
	"errors"
	"sync"
	"time"
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
}

func (*screen) Clear(s Style) {
}

func (*screen) Size() (int, int) {
	return 0, 0
}

func (*screen) Interrupt() {
}

func (*screen) Sync() {
}

func (*screen) Fini() {
}

func (*screen) Init() error {
	return errors.New("Curses not available")
}

func (*screen) Flush() {
}

func (*screen) PollEvent() Event {
	time.Sleep(1)
	return nil
}

func GetScreen() Screen {
	return DefaultScreen
}
