package modules

import (
	"fmt"
	"github.com/olekukonko/tablewriter"
	"math/rand"
	"os"
	"strconv"
	"throttle/tc"
	"time"
)

type TcProMgr struct {
}

var GlobTcProMgr *TcProMgr

func InitTcProMgr() (err error) {
	return
}

func (tcpm *TcProMgr) Add(dev string, rate string, ceil string, dst string) (err error) {

	var classid string
	var qdiscid string
	if err = tc.GlobTcMgr.Qdisc(dev); err != nil {
		return
	}

	classid = fmt.Sprintf(`%s:%s`, tc.MAJOR, strconv.Itoa(tcpm.random()))
	if err = tc.GlobTcMgr.Classful(dev, classid, rate, ceil); err != nil {
		return
	}
	qdiscid = fmt.Sprintf(`%s:`, tc.MAJOR)
	if err = tc.GlobTcMgr.Filter(dev, qdiscid, dst, classid); err != nil {
		return
	}
	return
}

func (tcm *TcProMgr) random() (r int) {
	rand.Seed(time.Now().UnixNano())
	r = rand.Intn(200)
	return
}

func (tcm *TcProMgr) Show(dev string) (err error) {
/*	if err = tc.GlobTcMgr.ClassfulShow(dev); err != nil {
		return
	}*/
	var data [][]string
	if data, err = tc.GlobTcMgr.FilterShows(dev); err != nil {
		return
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{`qdiscId`, `classId`, `rate`, `ceil`, `Address segment`})
	table.SetFooter([]string{``,``,``, `Total Filters`, strconv.Itoa(len(data))})
	table.SetAutoMergeCells(true)
	table.SetRowLine(true)
	table.AppendBulk(data)
	table.Render()
	return
}
