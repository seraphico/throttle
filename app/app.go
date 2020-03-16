package main

import (
	"github.com/urfave/cli"
	"os"
	"throttle/modules"
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
			Action: func(c *cli.Context) (err error) {
				if err = InitThrottleProcess(); err != nil {
					return cli.NewExitError(err, 1)
				}

				if err = modules.GlobTcProMgr.Add(
					c.String(`i`),
					c.String(`r`),
					c.String(`C`),
					c.String(`t`),
					c.String(`d`),
				); err != nil {
					return cli.NewExitError(err, 1)
				}

				return
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "interface,i",
					Usage: `参数用于指定网络接口`,
					Value: "eth0",
				},
				&cli.StringFlag{
					Name:  "network,t",
					Usage: "用于指定带宽限制的目标地址",
					Value: "127.0.0.0/24",
				},
				&cli.StringFlag{
					Name:  "direction,d",
					Usage: "用于指定限制适用的方向",
					Value: "down",
				},
				&cli.StringFlag{
					Name:  "rate,r",
					Usage: "用于在add操作时指定带宽",
					Value: `100`,
				},
				&cli.StringFlag{
					Name:  "ceil,C",
					Usage: "用于指定突发带宽",
					Value: `100`,
				},
			},
		},
		{
			Name:  `ls`,
			Usage: `查看规则`,
			Action: func(c *cli.Context) (err error) {

				if err = InitThrottleProcess(); err != nil {
					return cli.NewExitError(err, 1)
				}

				if err = modules.GlobTcProMgr.Show(
					c.String(`i`),
					c.String(`t`),
				); err != nil {
					return cli.NewExitError(err, 1)
				}
				return
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  `interface,i`,
					Usage: `指定查询规则的设备`,
					Value: `eth0`,
				},
				&cli.StringFlag{
					Name:  "network,t",
					Usage: "指定查询的IP地址段",
					Value: "all",
				},
			},
		},
		{
			Name:  `rm`,
			Usage: `删除规则`,
			Action: func(c *cli.Context) (err error) {
				if err = InitThrottleProcess(); err != nil {
					return cli.NewExitError(err, 1)
				}
				if err = modules.GlobTcProMgr.Delete(
					c.String(`i`),
					c.String(`t`),
				); err != nil {
					return cli.NewExitError(err, 1)
				}
				return
			},
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "interface,i",
					Usage: "指定网络接口",
					Value: "eth0",
				},
				&cli.StringFlag{
					Name:  "network,t",
					Usage: "指定网段或者ip",
					Value: "127.0.0.1",
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
	if err = modules.InitTcProMgr(); err != nil {
		return
	}
	if err = modules.InitTableMgr(); err != nil {
		return
	}
	return
}
