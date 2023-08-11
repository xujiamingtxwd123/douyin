package externaldata

import (
	"github.com/xujiamingtxwd123/douyin/open"
	"github.com/xujiamingtxwd123/douyin/open/config"
	"github.com/xujiamingtxwd123/douyin/util"
	"testing"
)

var cfg = &config.Config{
	ClientKey:    "",
	ClientSecret: "",
	RedirectURL:  "",
	Scopes:       "",
	Cache:        util.NewMemCache(),
}

// errcode=2100004 , errmsg=系统繁忙，此时请开发者稍候再试
func TestBase(t *testing.T) {
	api := open.NewOpenAPI(cfg)

	externalVideo := NewExternalVideo(api.GetContext())

	videoInfo, err := externalVideo.Base("accessToken",
		"openid",
		"item_id")
	t.Logf("%+v err:%+v\n", videoInfo, err)
}

func TestLike(t *testing.T) {
	api := open.NewOpenAPI(cfg)

	externalVideo := NewExternalVideo(api.GetContext())

	videoInfo, err := externalVideo.Like("accessToken",
		"openid",
		"item_id", 30)
	t.Logf("%+v err:%+v\n", videoInfo, err)
}
