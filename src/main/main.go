package main

import (
	_ "routers"
	"os"
	"errors"
	"utils"
	"mirbase"
	"github.com/stainberg/koalart"
	"wx"
)

func main() {
	if len(os.Args) == 1 {
		utils.LoadConfig("")
	} else if len(os.Args) == 2 {
		utils.LoadConfig(os.Args[1])
	} else {
		panic(errors.New("params error"))
	}
	mirbase.InitClient()
	go func() {
		wx.WxClient.Start()
	}()
	koala.Run(utils.Conf.HttpConf.RestAPIPort)
}
