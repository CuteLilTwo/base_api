/*
 * @Description: System parameters initialized
 * @Author: Young (hao_youngg@163.com)
 * @LastEditors: Young (hao_youngg@163.com)
 * @Date: 2019-04-30 13:35:52
 * @LastEditTime: 2019-04-30 13:55:39
 */
package init

import (
	"flag"
	"mcc-core/config"
	"mcc-core/connect"
	"path/filepath"
)

var (
	conf string
)

func init() {
	flag.StringVar(&conf, "c", "", "config-file :./config-test.json|./config-dev.json")
	flag.Parse()
	config_file, _ := filepath.Abs(conf)
	config.ConfInit(config_file) // Depending on the path, the configuration information is loaded
	connect.PsqlInit()           // Establish a PostgreSQL connection
	connect.RedisInit()          // Establish a Redis connection
}
