package externaldata

import (
	"encoding/json"
	"fmt"
	"github.com/xujiamingtxwd123/douyin/open/context"
	"github.com/xujiamingtxwd123/douyin/util"
)

const (
	//获取视频基础数据
	videoBaseURL string = "https://open.douyin.com/data/external/item/base/?open_id=%s&item_id=%s"
	//获取视频点赞数据
	videoLikeURL string = "https://open.douyin.com/data/external/item/like/?open_id=%s&item_id=%s&date_type=%d"
	//获取视频评论数据
	videoCommentURL string = "https://open.douyin.com/data/external/item/comment/?open_id=%s&item_id=%s"
	//获取视频播放数据
	videoPlayURL string = "https://open.douyin.com/data/external/item/play/?open_id=%s&item_id=%s"
	//获取视频分享数据
	videoShareURL string = "https://open.douyin.com/data/external/item/share/?open_id=%s&item_id=%s"
)

// Video 视频
type ExternalVideo struct {
	*context.Context
}

// NewVideo .
func NewExternalVideo(context *context.Context) *ExternalVideo {
	externalVideo := new(ExternalVideo)
	externalVideo.Context = context
	return externalVideo
}

type BaseInfo struct {
	util.CommonError
	Result struct {
		TotalLike           int64   `json:"total_like"`
		TotalComment        int64   `json:"total_comment"`
		TotalShare          int64   `json:"total_share"`
		AveragePlayDuration float32 `json:"avg_play_duration"`
		TotalPlay           int64   `json:"total_play"`
	} `json:"result"`
}

type BaseInfoRes struct {
	util.CommonErrorExtra
	Data BaseInfo `json:"data"`
}

func (externalVideo *ExternalVideo) Base(accessToken, openid, itemId string) (info *BaseInfo, err error) {

	uri := fmt.Sprintf(videoBaseURL, openid, itemId)
	var response []byte
	response, err = util.HTTPGet2(uri, accessToken)
	if err != nil {
		return
	}
	var result BaseInfoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Base error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}

	fmt.Printf("Base %+v err:%+v\n", result, err)
	info = &result.Data
	return
}

type LikeInfo struct {
	util.CommonError
	ResultList []struct {
		Date string `json:"date"`
		Like int64  `json:"like"`
	} `json:"result_list"`
}

type LikeInfoRes struct {
	util.CommonErrorExtra
	Data LikeInfo `json:"data"`
}

func (externalVideo *ExternalVideo) Like(accessToken, openid, itemId string, dateType int) (info *LikeInfo, err error) {

	uri := fmt.Sprintf(videoLikeURL, openid, itemId, dateType)
	var response []byte
	response, err = util.HTTPGet2(uri, accessToken)
	if err != nil {
		return
	}
	var result LikeInfoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Base error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	info = &result.Data
	return
}
