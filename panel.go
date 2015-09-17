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

type Panel struct {
	view     View
	title    Widget
	titleV   *ViewPort
	status   Widget
	statusV  *ViewPort
	bottom   Widget
	bottomV  *ViewPort
	content  Widget
	contentV *ViewPort
}

func (p *Panel) Draw() {
	if p.title != nil {
		p.title.Draw()
	}
	if p.status != nil {
		p.status.Draw()
	}
	if p.bottom != nil {
		p.bottom.Draw()
	}
	if p.content != nil {
		p.content.Draw()
	}
}

func (p *Panel) HandleEvent(e Event) bool {
	rv := false
	// Only the content pane will do anything with the event.
	if p.content != nil {
		rv = p.content.HandleEvent(e)
	}
	return rv
}

func (p *Panel) SetView(view View) {
	p.view = view
	if view == nil {
		return
	}
	w, h := view.Size()
	y := 0
	if p.title != nil {
		p.titleV = NewViewPort(view, 0, y, w, 1)
		p.title.SetView(p.titleV)
		y++
	}
	if p.status != nil {
		p.statusV = NewViewPort(view, 0, y, w, 1)
		p.status.SetView(p.statusV)
		y++
	}
	if p.bottom != nil {
		p.bottomV = NewViewPort(view, 0, h-1, w, 1)
		p.bottom.SetView(p.bottomV)
		h--
	}
	if p.content != nil {
		p.contentV = NewViewPort(view, 0, y, w, h-y)
		p.content.SetView(p.contentV)
	}
	p.Resize()
}

func (p *Panel) Resize() {
	if p.view == nil {
		return
	}
	w, h := p.view.Size()
	y := 0

	if p.title != nil {
		p.titleV.Resize(0, y, w, 1)
		y++
	}
	if p.status != nil {
		p.statusV.Resize(0, y, w, 1)
		y++
	}
	if p.bottom != nil {
		p.bottomV.Resize(0, h-1, w, 1)
		h--
	}
	if p.content != nil {
		p.contentV.Resize(0, y, w, h-y)
	}
	if p.title != nil {
		p.title.Resize()
	}
	if p.status != nil {
		p.status.Resize()
	}
	if p.bottom != nil {
		p.bottom.Resize()
	}
	if p.content != nil {
		p.content.Resize()
	}
}

func (p *Panel) SetTitle(w Widget) {
	p.titleV = NewViewPort(p.view, 0, 0, 0, 0)
	p.title = w
	w.SetView(p.titleV)
	p.Resize()
}

func (p *Panel) SetStatus(w Widget) {
	p.statusV = NewViewPort(p.view, 0, 0, 0, 0)
	p.status = w
	w.SetView(p.statusV)
	p.Resize()
}

func (p *Panel) SetBottom(w Widget) {
	p.bottomV = NewViewPort(p.view, 0, 0, 0, 0)
	p.bottom = w
	w.SetView(p.bottomV)
	p.Resize()
}

func (p *Panel) SetContent(w Widget) {
	p.contentV = NewViewPort(p.view, 0, 0, 0, 0)
	p.content = w
	w.SetView(p.contentV)
	p.Resize()
}

func NewPanel() *Panel {
	return &Panel{}
}
