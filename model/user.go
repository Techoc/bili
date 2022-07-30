package model

import (
	"encoding/json"
	"fmt"
	"github.com/techoc/bili/util"
	"gorm.io/gorm"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

type User struct {
	gorm.Model
	Uid       int64  //b站uid
	Name      string //昵称
	Likes     int64  //获赞数
	Archive   int64  //播放量
	Article   int64  //阅读数
	Black     int64  //保留
	Follower  int64  //粉丝数
	Following int64  //关注数
	Whisper   int64  //保留
	Sign      string //个性签名
	Sex       string //性别
}

func FindUserByUid(uid int64) (User, error) {
	var user User
	result := DB.Where("uid = ?", uid).First(&user)
	return user, result.Error
}

func (user *User) AddUser() {
	DB.Create(&user)
}
func (user *User) UpdateUser() {
	DB.Where("uid = ?", user.Uid).Updates(&user)
}

func (user *User) GetInfoData(uid int64) {
	info, err := http.Get("https://api.bilibili.com/x/space/acc/info?mid=" + strconv.FormatInt(uid, 10) + "&jsonp=jsonp")
	if err != nil {
		util.Log().Panic("获取info信息失败, %s", err)
		return
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
		return
	}
	var data Info
	err = json.Unmarshal(body, &data)
	if err != nil {
		util.Log().Error("获取info信息失败, %s", err)
		return
	}
	user.Uid = data.Data.Mid
	user.Name = data.Data.Name
	user.Sign = data.Data.Sign
	user.Sex = data.Data.Sex
}

func (user *User) GetStatData(uid int64) {
	stat, err := http.Get("https://api.bilibili.com/x/relation/stat?vmid=" + strconv.FormatInt(uid, 10) + "&jsonp=jsonp")
	if err != nil {
		util.Log().Error("%s获取stat信息失败, %s", uid, err)
		return
	}
	body, err := io.ReadAll(stat.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(stat.Body)
	if err != nil {
		util.Log().Error("%s获取stat信息失败, %s", uid, err)
		return
	}
	var data Stat
	err = json.Unmarshal(body, &data)
	if err != nil {
		util.Log().Error("%s获取stat信息失败, %s", uid, err)
		return
	}
	user.Follower = data.Data.Follower
	user.Whisper = data.Data.Whisper
	user.Following = data.Data.Following
}
func (user *User) GetUpStatData(uid int64) {
	url := "https://api.bilibili.com/x/space/upstat?mid=" + strconv.FormatInt(uid, 10) + "&jsonp=jsonp"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Cookie", "buvid3=F698BB7E-BBC7-7E52-AD38-1A98F797140121091infoc; i-wanna-go-back=-1; _uuid=EB228C106-138F-31C2-1DEA-A81025B86675722040infoc; buvid4=DD554E4B-D158-7827-0FCD-70FBFFB992E023918-022072412-AOc2/D4Q7+/BWdE8vMrxmg==; buvid_fp=e169cb2adb7cdd1498f5b633e2c5baa6; fingerprint=0655f77574f0a8731cb2e522d1766876; DedeUserID=6092912; DedeUserID__ckMd5=44c7ab312468d2b8; b_ut=5; CURRENT_BLACKGAP=0; blackside_state=0; CURRENT_QUALITY=80; rpdid=|(u)~J~lYkYY0J'uYl)YY~kJu; hit-dyn-v2=1; LIVE_BUVID=AUTO2116586718223815; nostalgia_conf=-1; bsource=search_baidu; SESSDATA=c6eb5a66,1674636743,a0fa0*71; bili_jct=33b060d927ae87bfee469c65fc1b144c; innersign=1; sid=58sf1s1l; b_lsid=C43A56A7_1824A498D90; PVID=1; CURRENT_FNVAL=4048; b_timer={\"ffp\"=<buvid3=F698BB7E-BBC7-7E52-AD38-1A98F797140121091infoc; i-wanna-go-back=-1; _uuid=EB228C106-138F-31C2-1DEA-A81025B86675722040infoc; buvid4=DD554E4B-D158-7827-0FCD-70FBFFB992E023918-022072412-AOc2/D4Q7+/BWdE8vMrxmg==; buvid_fp=e169cb2adb7cdd1498f5b633e2c5baa6; fingerprint=0655f77574f0a8731cb2e522d1766876; DedeUserID=6092912; DedeUserID__ckMd5=44c7ab312468d2b8; b_ut=5; CURRENT_BLACKGAP=0; blackside_state=0; CURRENT_QUALITY=80; rpdid=|(u)~J~lYkYY0J'uYl)YY~kJu; hit-dyn-v2=1; LIVE_BUVID=AUTO2116586718223815; nostalgia_conf=-1; bsource=search_baidu; SESSDATA=c6eb5a66,1674636743,a0fa0*71; bili_jct=33b060d927ae87bfee469c65fc1b144c; innersign=1; sid=58sf1s1l; b_lsid=C43A56A7_1824A498D90; PVID=1; CURRENT_FNVAL=4048; b_timer={\"ffp\">")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%v\n", string(body))
	var data UpStat
	err = json.Unmarshal(body, &data)
	if err != nil {
		util.Log().Error("%s获取upstat信息失败, %s", uid, err)
		return
	}
	user.Article = data.Data.Article.View
	user.Archive = data.Data.Archive.View
	user.Likes = data.Data.Likes
}
