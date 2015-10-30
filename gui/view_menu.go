package gui

import (
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/twstrike/coyim/config"
)

type viewMenu struct {
	merge   *gtk.CheckMenuItem
	offline *gtk.CheckMenuItem
}

func (v *viewMenu) setFromConfig(c *config.Accounts) {
	glib.IdleAdd(func() bool {
		v.merge.SetActive(c.MergeAccounts)
		v.offline.SetActive(!c.ShowOnlyOnline)
		return false
	})
}

func (u *gtkUI) toggleMergeAccounts() {
	if u.config != nil {
		u.config.MergeAccounts = u.viewMenu.merge.GetActive()
		u.saveConfigOnly()
	}

	u.roster.redrawIfRosterVisible()
}

func (u *gtkUI) toggleShowOffline() {
	if u.config != nil {
		u.config.ShowOnlyOnline = !u.viewMenu.offline.GetActive()
		u.saveConfigOnly()
	}

	u.roster.redrawIfRosterVisible()
}