package externaldata

import "github.com/xujiamingtxwd123/douyin/open/context"

// Externaldata 数据开放服务.
type Externaldata struct {
	*context.Context
}

// NewExternaldata .
func NewExternaldata(context *context.Context) *Externaldata {
	externaldata := new(Externaldata)
	externaldata.Context = context
	return externaldata
}
