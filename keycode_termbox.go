// +build freebsd linux netbsd openbsd darwin

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

type KeyCode termbox.Key

//
// We take a very simplistic approach here, and do not expose every possible
// key code.  Only those codes which are likely to be useful in real apps
// are listed.  Applications shouldn't rely on non-portable key codes.
//
const (
	KeyF1        KeyCode = KeyCode(termbox.KeyF1)
	KeyF2                = KeyCode(termbox.KeyF2)
	KeyF3                = KeyCode(termbox.KeyF3)
	KeyF4                = KeyCode(termbox.KeyF4)
	KeyF5                = KeyCode(termbox.KeyF5)
	KeyF6                = KeyCode(termbox.KeyF6)
	KeyF7                = KeyCode(termbox.KeyF7)
	KeyF8                = KeyCode(termbox.KeyF8)
	KeyF9                = KeyCode(termbox.KeyF9)
	KeyF10               = KeyCode(termbox.KeyF10)
	KeyF11               = KeyCode(termbox.KeyF11)
	KeyF12               = KeyCode(termbox.KeyF12)
	KeyInsert            = KeyCode(termbox.KeyInsert)
	KeyDelete            = KeyCode(termbox.KeyDelete)
	KeyHome              = KeyCode(termbox.KeyHome)
	KeyEnd               = KeyCode(termbox.KeyEnd)
	KeyUp                = KeyCode(termbox.KeyArrowUp)
	KeyDown              = KeyCode(termbox.KeyArrowDown)
	KeyLeft              = KeyCode(termbox.KeyArrowLeft)
	KeyRight             = KeyCode(termbox.KeyArrowRight)
	KeyPgDn              = KeyCode(termbox.KeyPgdn)
	KeyPgUp              = KeyCode(termbox.KeyPgup)
	KeyCtrlA             = KeyCode(termbox.KeyCtrlA)
	KeyCtrlB             = KeyCode(termbox.KeyCtrlB)
	KeyCtrlC             = KeyCode(termbox.KeyCtrlC)
	KeyCtrlD             = KeyCode(termbox.KeyCtrlD)
	KeyCtrlE             = KeyCode(termbox.KeyCtrlE)
	KeyCtrlF             = KeyCode(termbox.KeyCtrlF)
	KeyCtrlG             = KeyCode(termbox.KeyCtrlG)
	KeyCtrlH             = KeyCode(termbox.KeyCtrlH)
	KeyCtrlI             = KeyCode(termbox.KeyCtrlI)
	KeyCtrlJ             = KeyCode(termbox.KeyCtrlJ)
	KeyCtrlK             = KeyCode(termbox.KeyCtrlK)
	KeyCtrlL             = KeyCode(termbox.KeyCtrlL)
	KeyCtrlM             = KeyCode(termbox.KeyCtrlM)
	KeyCtrlN             = KeyCode(termbox.KeyCtrlN)
	KeyCtrlO             = KeyCode(termbox.KeyCtrlO)
	KeyCtrlP             = KeyCode(termbox.KeyCtrlP)
	KeyCtrlQ             = KeyCode(termbox.KeyCtrlQ)
	KeyCtrlR             = KeyCode(termbox.KeyCtrlR)
	KeyCtrlS             = KeyCode(termbox.KeyCtrlS)
	KeyCtrlT             = KeyCode(termbox.KeyCtrlT)
	KeyCtrlU             = KeyCode(termbox.KeyCtrlU)
	KeyCtrlV             = KeyCode(termbox.KeyCtrlV)
	KeyCtrlW             = KeyCode(termbox.KeyCtrlW)
	KeyCtrlX             = KeyCode(termbox.KeyCtrlX)
	KeyCtrlY             = KeyCode(termbox.KeyCtrlY)
	KeyCtrlZ             = KeyCode(termbox.KeyCtrlZ)
	KeyTab               = KeyCode(termbox.KeyTab)
	KeyBackspace         = KeyCode(termbox.KeyBackspace)
	KeyEnter             = KeyCode(termbox.KeyEnter)
	KeyEsc               = KeyCode(termbox.KeyEsc)
)
