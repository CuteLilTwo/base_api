/*
 * @Description: The main file entry to start the project
 * @Author: young(hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 11:25:45
 * @LastEditTime: 2019-04-30 17:03:33
 */
package main

import (
	"base_api/api/route"
	_ "base_api/init"
	"base_core/config"

	"github.com/kataras/iris"
)

/**
 * @description: Project initiation method
 * @param {nil}
 * @return: {nil}
 */
func main() {
	app := iris.New()

	new(route.TestRouter).SetTestRouter(app, "/test")

	app.Run(iris.Addr(config.Conf.Env.Port))
}

// test Request json parameters (either get or post)
// {
// 	"TestControllerInputstr":"Controller Data Input ->"
// }
