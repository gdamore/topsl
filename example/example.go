package main

import (
	"github.com/gdamore/topsl"
	//	"github.com/nsf/termbox-go"
	"log"
	"os"
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

	if out, e := os.Create("/tmp/debuglog"); e == nil {
		log.SetOutput(out)
	}

	topsl.AppInit()
	defer topsl.AppFini()
	//termbox.Init()
	//defer termbox.Close()
	//s := topsl.Screen

	//	_, height := s.Size()
	//s.Clear(topsl.StyleDefault)

	title := topsl.NewTitleBar("My Test")
	title.SetRight("Example v1.0", topsl.StyleTextBar)
	keyb := topsl.NewKeyBar([]string{"_Quit", "_About", "e_Xit"})

	panel := topsl.NewPanel()

	panel.SetBottom(keyb)

	panel.SetTitle(title)

	content := topsl.NewTextArea()
	content.SetContent("This is a test\nAnd another line\nAnd more and more\n" +
		"A very very very long line.  The quick brown fox jumps over the " +
		"lazy dog.  The mouse ran up the clock.  Blah blah blah.")
	content.SetCursorEnabled(true)

	panel.SetContent(content)

	status := topsl.NewStatusBar("Some thing happened")
	panel.SetStatus(status)
	status.SetStatus("Heck YEAH!! (KookAid Man)")

	//	panel.Draw()
	//	termbox.Flush()

	//	app.SetView(topsl.Screen)
	log.Printf("Here goes nuttin")
	app.panel = panel
	topsl.SetApplication(app)
	topsl.AppInit()
	topsl.RunApplication()
}
