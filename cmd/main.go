package main

import (
	"fmt"
	conf "github.com/kasiforce/trade/config"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/cache"
	"github.com/kasiforce/trade/repository/db/dao"
	"github.com/kasiforce/trade/routes"
	"github.com/kasiforce/trade/service/pay"
)

func main() {
	loading()
	r := routes.NewRouter()
	fmt.Println("启动成功...")
	_ = r.Run(conf.Config.System.HttpPort)
}

func loading() {
	conf.InitConfig() //配置文件初始化
	util.InitLog()    //日志文件初始化
	dao.InitMySQL()   //数据库初始化
	cache.InitCache() //redis初始化
	pay.InitAlipay()  //支付包SDK初始化
}
