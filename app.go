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

var appWidget Widget
var appStopq chan struct{}

func SetApplication(app Widget) {
	appWidget = app
}

func AppInit() error {
	appStopq = make(chan struct{})
	return GetScreen().Init()
}

func AppFini() {
	if appStopq != nil {
		close(appStopq)
		appStopq = nil
	}
	GetScreen().Fini()
}

func AppRedraw() {
	if appWidget != nil {
		GetScreen().Sync()
	}
}

func AppDraw() {
	if appWidget != nil {
		go GetScreen().Interrupt()
	}
}

func AppLock() {
	GetScreen().Lock()
}

func AppUnlock() {
	GetScreen().Unlock()
}

func RunApplication() {

	scr := GetScreen()
	if appWidget == nil {
		return
	}
	stopq := appStopq

	scr.Clear(StyleDefault)
	appWidget.SetView(scr)

	for {
		if stopq != nil {
			select {
			case <-stopq:
				return
			default:
			}
		}

		appWidget.Draw()
		scr.Flush()

		ev := scr.PollEvent()
		switch ev.(type) {
		case *ResizeEvent:
			scr.Sync()
			appWidget.Resize()
		case *KeyEvent:
			appWidget.HandleEvent(ev)
		}
	}
}
