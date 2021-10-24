package function

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"goDemo/config"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"
)

func Print(i interface{}) {
	fmt.Println("---")
	fmt.Println(i)
	fmt.Println("---")
}

func RetJson(code, msg string, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
	c.Abort()
}

func GetCurrentTimeStamp() int64 {
	return time.Now().Unix()
}

func GetCurrentTimeStr() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func CreateSign(params url.Values) string {
	var key []string
	for s := range params {
		if s != "sn" {
			key = append(key, s)
		}
	}
	sort.Strings(key)
	var str = ""
	for i := 0; i < len(key); i++ {
		if i == 0 {
			str += fmt.Sprintf("%v=%v", key[i], params.Get(key[i]))
		} else {
			str += fmt.Sprintf("&%v=%v", key[i], params.Get(key[i]))
		}
	}
	sign := MD5(MD5(str) + MD5(config.APP_NAME+config.APP_SECRET))
	return sign
}

func VerifySign(ctx *gin.Context) (res map[string]string, err error) {
	method := ctx.Request.Method
	var sn string
	var params url.Values
	var debug string
	ts, _ := strconv.ParseInt(ctx.Query("ts"), 10, 64)
	if method == "GET" {
		params = ctx.Request.URL.Query()
		sn = ctx.Query("sn")
		debug = ctx.Query("debug")
	} else if method == "POST" {
		ctx.Request.ParseForm()
		params = ctx.Request.PostForm
		sn = ctx.PostForm("sn")
		debug = ctx.PostForm("debug")
	} else {
		err = errors.New("非法请求")
		return
	}
	if debug == "1" {
		res = map[string]string{
			"ts": strconv.FormatInt(GetCurrentTimeStamp(), 10),
			"sn": CreateSign(params),
		}
		return
	}

	exp, _ := strconv.ParseInt(config.APP_EXPIRE, 10, 64)
	if ts > GetCurrentTimeStamp() || GetCurrentTimeStamp()-ts >= exp {
		err = errors.New("TS error")
		return
	}
	if sn == "" || sn != CreateSign(params) {
		err = errors.New("SN error")
		return
	}
	return
}
