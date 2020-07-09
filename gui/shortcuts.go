package gui

import (
	"github.com/coyim/gotk3adapter/gtki"
	"github.com/digitalautonomy/wahay/tor"
)

func (u *gtkUI) connectShortcut(accel string, w gtki.Window, action func(gtki.Window)) {
	gr, _ := u.g.gtk.AccelGroupNew()
	key, mod := u.g.gtk.AcceleratorParse(accel)

	gr.Connect2(key, mod, gtki.ACCEL_VISIBLE, func() {
		action(w)
	})

	w.AddAccelGroup(gr)
}

func (u *gtkUI) connectShortcutsMainWindow(w gtki.Window) {
	// <Primary> maps to Command and OS X, but Control on other platforms
	u.connectShortcut("<Primary>q", w, u.closeApplicationWindow)
	u.connectShortcut("<Alt>F4", w, u.closeApplicationWindow)
}

func (u *gtkUI) connectShortcutsHostingMeetingConfigurationWindow(w gtki.Window) {
	// <Primary> maps to Command and OS X, but Control on other platforms
	u.connectShortcut("<Primary>q", w, u.closeApplicationWindow)
	u.connectShortcut("<Alt>F4", w, u.closeApplicationWindow)
}

func (u *gtkUI) connectShortcutCurrentHostMeetingWindow(w gtki.Window, h *hostData) {
	// <Primary> maps to Command and OS X, but Control on other platforms
	u.connectShortcut("<Primary>q", w, func(w gtki.Window) {
		h.leaveHostMeeting()
	})
	u.connectShortcut("<Primary>w", w, func(w gtki.Window) {
		h.finishMeeting()
	})
}

func (u *gtkUI) connectShortcutCurrentMeetingWindow(w gtki.Window, m tor.Service) {
	// <Primary> maps to Command and OS X, but Control on other platforms
	u.connectShortcut("<Primary>q", w, func(w gtki.Window) {
		u.leaveMeeting(m)
	})
}

func (u *gtkUI) closeApplicationWindow(w gtki.Window) {
	u.quit()
}