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

package main

import (
	"fmt"
	"os"

	"github.com/gdamore/topsl"
)

type MyApp struct {
	panel topsl.Widget
	view  topsl.View
}

func (a *MyApp) HandleEvent(ev topsl.Event) bool {
	if a.panel != nil {
		if a.panel.HandleEvent(ev) {
			return true
		}

		switch ev := ev.(type) {
		case *topsl.KeyEvent:
			switch ev.Ch {
			case 0:
			case 'Q', 'q':
				topsl.AppFini()
				return true
			}
		}
	}
	return false
}

func (a *MyApp) Draw() {
	if a.panel != nil {
		a.panel.Draw()
	}
}

func (a *MyApp) Resize() {
	if a.panel != nil {
		a.panel.Resize()
	}
}

func (a *MyApp) SetView(view topsl.View) {
	if a.panel != nil {
		a.panel.SetView(view)
	}
	a.view = view
}

func main() {

	app := &MyApp{}

	if e := topsl.AppInit(); e != nil {
		fmt.Fprintln(os.Stderr, e.Error())
		os.Exit(1)
	}

	title := topsl.NewTitleBar()
	title.SetCenter("TextArea Test")
	title.SetRight("Example v1.0")

	keyb := topsl.NewKeyBar()
	keyb.SetKeys([]string{"[Q] Quit"})

	panel := topsl.NewPanel()

	panel.SetBottom(keyb)
	panel.SetTitle(title)

	content := topsl.NewTextArea()
	content.SetContent("This is a test\nAnd another line\nAnd more and more\n" +
		"A very very very long line.  The quick brown fox jumps over the " +
		"lazy dog.  The mouse ran up the clock.  Blah blah blah.")
	content.EnableCursor(true)

	panel.SetContent(content)

	status := topsl.NewStatusBar()
	panel.SetStatus(status)
	status.SetStatus("Number 5 is Alive!")

	app.panel = panel
	topsl.SetApplication(app)
	topsl.RunApplication()
}
