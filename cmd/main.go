package main

import (
	"fmt"
	_ "fmt"
	"github.com/CocaineCong/gin-mall/routes"
	conf "github.com/kasiforce/trade/config"
	"github.com/kasiforce/trade/pkg/util"
	"github.com/kasiforce/trade/repository/db/dao"
)

func main() {
	loading()
	r := routes.NewRouter()
	_ = r.Run(conf.Config.System.HttpPort)
	fmt.Println("启动成功...")
}

func loading() {
	conf.InitConfig() //配置文件初始化
	util.InitLog()    //日志文件初始化
	dao.InitMySQL()   //数据库初始化
}
