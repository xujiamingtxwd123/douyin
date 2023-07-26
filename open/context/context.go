package context

import (
	"github.com/gzw13999/douyin/open/config"
	"github.com/gzw13999/douyin/open/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
