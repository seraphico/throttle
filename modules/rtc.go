package modules

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
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

//ud //UP Or Down
func (tcpm *TcProMgr) Add(dev string, rate string, ceil string, d string, ud string) (err error) {

	var classid string
	var qdiscid string
	if err = tc.GlobTcMgr.Qdisc(dev); err != nil {
		return
	}

	classid = fmt.Sprintf(`%s:%s`, tc.DOWNMAJOR, strconv.Itoa(tcpm.random()))
	if err = tc.GlobTcMgr.Classful(dev, classid, rate, ceil); err != nil {
		return
	}
	qdiscid = fmt.Sprintf(`%s:`, tc.DOWNMAJOR)
	if err = tc.GlobTcMgr.Filter(dev, qdiscid, d, classid, ud); err != nil {
		return
	}
	return
}

func (tcm *TcProMgr) random() (r int) {
	rand.Seed(time.Now().UnixNano())
	r = rand.Intn(200)
	return
}

func (tcm *TcProMgr) Show(dev string, netstr string) (err error) {
	/*	if err = tc.GlobTcMgr.ClassfulShow(dev); err != nil {
		return
	}*/
	var data [][]string
	switch netstr {
	case `all`:
		if data, err = tc.GlobTcMgr.FilterShows(dev); err != nil {
			return
		}
	default:
		if data, err = tc.GlobTcMgr.FilterShowsWithString(dev, netstr); err != nil {
			return
		}
	}

	if err = GlobTableMgr.Out(data); err != nil {
		return
	}

	return
}

func (tcm *TcProMgr) Delete(dev string, netstr string) (err error) {
	var netip net.IP
	var datas [][]string
	if _, _, err = net.ParseCIDR(netstr); err != nil {
		if netip = net.ParseIP(netstr); netip == nil {
			return errors.New(`Illegal IP address`)
		}
		return
	}

	if datas, err = tc.GlobTcMgr.FilterShowsWithString(dev, netstr); err != nil {
		return
	}
	for _, data := range datas {
		if err = tc.GlobTcMgr.FilterDelete(dev, data[2]); err != nil {
			return
		}

		if err = tc.GlobTcMgr.ClassFulDelet(dev, data[1]); err != nil {
			return
		}
	}
	return
}
