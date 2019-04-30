/*
 * @Description: The middleware
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 15:55:52
 * @LastEditTime: 2019-04-30 17:02:55
 */
package route

import (
	"base_core/config"
	"base_core/connect"
	"base_core/log"
	"base_core/utils"
	"fmt"
	"io/ioutil"

	"github.com/kataras/iris"
)

func DataGet(ctx iris.Context) {
	reqBody, _ := ioutil.ReadAll(ctx.Request().Body)
	ctx.Values().Set("reqBody", reqBody)
	ctx.Next()
}

// jwt auth before requset
func MustJWTBefore(ctx iris.Context) {
	// 获取AuthToken
	tokenString := ctx.GetHeader("Authorization")
	// Header头内无Token
	if tokenString == "" {
		log.ControllerLogger.Warn("Please login first")
		ctx.JSON(iris.Map{
			"Code": config.BACK_TO_LOGIN,
			"Msg":  "Please login first",
			"Data": "",
		})
		return
	}
	// 解析AuthToken
	token, err := utils.AuthToken(tokenString, config.Conf.Jwt.Pwd)
	if err != nil || token == nil {
		fmt.Println(token, "|", err)
		log.ControllerLogger.Warn("用户登陆信息已失效")
		ctx.JSON(iris.Map{
			"Code": config.BACK_TO_LOGIN,
			"Msg":  "The login information is invalid. Please login again",
			"Data": "",
		})
		return
	}
	// 单点登陆session设置
	Key := &utils.KeyData{
		UserId: int(token["UID"].(float64)),
	}
	authkey := Key.GetKey("USER_LOGIN_KEY")
	v, _ := connect.RedisEngine.Get(authkey).Result()
	if v == "" {
		log.ControllerLogger.Warn("User login information has expired")
		ctx.JSON(iris.Map{
			"Code": config.BACK_TO_LOGIN,
			"Msg":  "The login information is invalid. Please login again",
			"Data": "",
		})
		return
	}
	if v != tokenString {
		log.ControllerLogger.Warn("The current user is logged in on another client", int(token["UID"].(float64)))
		ctx.JSON(iris.Map{
			"Code": config.BACK_TO_LOGIN,
			"Msg":  "The login information is invalid. Please login again",
			"Data": "",
		})
		return
	}
	// if you want to auth the user information,you can get the token info to auth it

	ctx.Values().Set("Token", token)
	ctx.Next()
}
