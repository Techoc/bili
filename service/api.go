package service

import (
	"encoding/json"
	"github.com/techoc/bili/model"
	"github.com/techoc/bili/util"
	"io"
	"net/http"
	"strconv"
)

func GetData(uid int64) {
	var user model.User
	user.GetInfoData(uid)
	user.GetStatData(uid)
	user.GetUpStatData(uid)
	// 查找用户
	if _, err := model.FindUserByUid(uid); err != nil {
		user.AddUser()
		util.Log().Info("添加成功")
	} else {
		util.Log().Info("%v已存在,更新数据", user.Uid)
		user.UpdateUser()
	}

}

func GetUid(uid int64) bool {
	info, err := http.Get("https://api.bilibili.com/x/space/acc/info?mid=" + strconv.FormatInt(uid, 10) + "&jsonp=jsonp")
	if err != nil {
		util.Log().Panic("获取info信息失败, %s", err)
		return false
	}
	body, err := io.ReadAll(info.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(info.Body)
	if err != nil {
		util.Log().Error("获取info信息失败, %s", err)
		return false
	}
	var data model.Info
	err = json.Unmarshal(body, &data)
	if err != nil {
		util.Log().Error("获取info信息失败, %s", err)
		return false
	}
	if data.Code == -404 || data.Code == -412 {
		util.Log().Info("%v获取id失败啦,原因&v", uid, data.Message)
		return false
	} else {
		util.Log().Info("%v获取id成功", uid)
		return true
	}
}
