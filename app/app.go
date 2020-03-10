package main

import (
	"github.com/urfave/cli"
	"os"
	"throttle/tc"
)

var app *cli.App

func init() {
	app = cli.NewApp()

	app.Name = "throttle"
	app.Usage = "throttle -h,--help"
	app.Authors = []cli.Author{
		{
			Name:  "seraphic",
			Email: "dongdong1260@gmail.com",
		},
	}
	app.Copyright = "©2010-2020 Seraphic Corporation,All Rights Reserved"
	app.Version = "release 1.0.0"
}

func main() {

	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "创建规则",
			Action: func(context *cli.Context) (err error) {
				if err = InitThrottleProcess(); err != nil {
					return cli.NewExitError(err, 1)
				}
				if err = tc.GTcMgr.Qdisc(context.String(`i`)); err != nil {
					return cli.NewExitError(err, 1)
				}
				return
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "interface,i",
					//Aliases: []string{`i`},
					Usage:   `参数用于指定网络接口，如 "eth0"`,
					Value:   "eth0",
				},
			},
		},
	}
	app.Run(os.Args)
}

func InitThrottleProcess() (err error) {
	if err = tc.InitTcMgr(); err != nil {
		return
	}

	return
}
