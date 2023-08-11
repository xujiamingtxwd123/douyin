package groupon

import "github.com/xujiamingtxwd123/douyin/open/context"

// Groupon .
type Groupon struct {
	*context.Context
}

// NewGroupon .
func NewGroupon(context *context.Context) *Groupon {
	groupon := new(Groupon)
	groupon.Context = context
	return groupon
}
