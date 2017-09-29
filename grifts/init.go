package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/spankie/buffalo_auth/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
