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

type Widget interface {
	// Draw is called to inform the widget to draw itself.
	Draw()

	// Resize is called in response to a window resize.  Unlike with
	// other events, Resize performed by parents first, and they must
	// then call their children.  This is because the children need to
	// see the updated sizes from the parents before they are called.
	Resize()

	// HandleEvent is called to ask the widget to handle any events.
	// If the widget has consumed the event, it should return true.
	// Generally, events are handled by the lower layers first, that
	// is for example, a button may have a chance to handle an event
	// before the enclosing window or panel.
	HandleEvent(ev Event) bool

	// SetView is used by callers to set the visual parent of the
	// Widget.  The Widget should use the View as a context for
	// drawing.
	SetView(view View)
}
