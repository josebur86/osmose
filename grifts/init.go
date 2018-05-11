package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/josebur86/osmose/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
