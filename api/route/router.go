/*
 * @Description: Routing distribution
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 15:45:11
 * @LastEditTime: 2019-04-30 16:24:32
 */
package route

import (
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

var (
	router iris.Party
)

type TestRouter struct{}

// 设置User一级路由
func (c *TestRouter) SetTestRouter(app *iris.Application, path string) {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // cross-domain request accept
		AllowCredentials: true,
	})
	router = app.Party(path, crs).AllowMethods(iris.MethodOptions)
	// 路由分发
	SetTestPathRouter()
}
