package context

import (
	"github.com/xujiamingtxwd123/douyin/open/config"
	"github.com/xujiamingtxwd123/douyin/open/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
