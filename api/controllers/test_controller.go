/*
 * @Description:
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 15:13:52
 * @LastEditTime: 2019-04-30 17:03:05
 */
package controllers

import (
	"base_core/config"
	"base_core/data"
	"base_core/log"
	"base_core/services"
	"encoding/json"

	"github.com/kataras/iris"
)

type TestController struct {
	Ctx         iris.Context
	TestService services.TestService
}

/**
 * @Description: /test/test for test
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @param {
	 condata.TestControllerInputData
 }
 * @return:
 * @Date: 2019-04-30 15:15:19
*/
func (c *TestController) TestController() {
	// Get the body parameter
	reqByte := c.Ctx.Values().Get("reqBody").([]byte)

	// Get token parameter
	// token := c.Ctx.Values().Get("Token").(jwt.MapClaims)

	// Parse the body parameter
	req := new(data.TestControllerInputData)
	if err := json.Unmarshal(reqByte, &req); err != nil {
		log.ControllerLogger.Error("json解析异常:", err.Error())
		c.Ctx.JSON(iris.Map{
			"Code": config.ERROR_CODE_CAN_NOT_SKIP,
			"Msg":  "test err",
			"Data": "",
		})
		return
	}

	// controller data out
	resp := c.TestService.TestServices(req)

	// handle the data for response
	resp.TestControllerOutputStr = resp.TestControllerOutputStr + " Controller Data Out"

	c.Ctx.JSON(iris.Map{
		"Code": config.SUCC_CODE,
		"Msg":  "test ok",
		"Data": resp,
	})
}
