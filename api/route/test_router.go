/*
 * @Description: The secondary routing
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 15:45:19
 * @LastEditTime: 2019-04-30 17:02:48
 */
package route

import (
	"base_api/api/controllers"
	"base_core/services"
	"base_core/utils"
	"time"

	"github.com/kataras/iris"
)

func SetTestPathRouter() {
	post := &controllers.TestController{
		TestService: services.NewTestService(),
	}
	router.Post("/test", DataGet, func(ctx iris.Context) {
		post.Ctx = ctx
		utils.ColorFmt("黑色", "白色", "闪烁", false, "API")
		utils.ColorFmt("黑色", "黄色", "默认", false, time.Now().Format("2006-01-02 15:04:05"))
		utils.ColorFmt("青蓝色", "白色", "默认", true, "HTTP POST:/test/test")
		post.TestController()
	})
	router.Get("/test", DataGet, func(ctx iris.Context) {
		post.Ctx = ctx
		utils.ColorFmt("黑色", "白色", "闪烁", false, "API")
		utils.ColorFmt("黑色", "黄色", "默认", false, time.Now().Format("2006-01-02 15:04:05"))
		utils.ColorFmt("青蓝色", "白色", "默认", true, "HTTP POST:/test/test")
		post.TestController()
	})
}
