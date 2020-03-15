package tc

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

type TcMgr struct {
}

const (
	MAJOR       = `1`
	MINOR       = `0`
	DEFAULT     = `1000`
	CLASSFULID  = `10`
	DEFAULTRATE = `10` //MB
	DEFAULTCEIL = `12` //MB
)

var GlobTcMgr *TcMgr

func InitTcMgr() (err error) {
	return
}

func (tcm *TcMgr) random() (r int) {
	rand.Seed(time.Now().UnixNano())
	r = rand.Intn(200)
	return
}

//十六进制转换为IP地址
//aca80000 ---> 172.168.0.0
//ffffff00 ---> 255.255.255.0
func (tcm *TcMgr) hexToIp(hexstr string) (ip string) {

	var h []byte
	h, _ = hex.DecodeString(hexstr)
	ip = fmt.Sprintf("%v.%v.%v.%v", h[0], h[1], h[2], h[3])
	return
}
func (tcm *TcMgr) ipSubNetMaskToInt(netmask string) (imask int, err error) {
	var (
		ipNetMasks  []string
		ipv4MaskArr []byte
	)
	ipNetMasks = strings.Split(netmask, ".")
	ipv4MaskArr = make([]byte, 4)
	if len(ipNetMasks) != 4 {
		err = fmt.Errorf("netmask: %s is not valid. pattern should like: 255.255.255.0", netmask)
		return
	}

	for i, value := range ipNetMasks {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			return 0, fmt.Errorf("ipMaskToInt call strconv.Atoi error:[%v] string value is: [%s]", err, value)
		}
		if intValue > 255 {
			return 0, fmt.Errorf("netmask cannot greater than 255, current value is: [%s]", value)
		}
		ipv4MaskArr[i] = byte(intValue)
	}

	imask, _ = net.IPv4Mask(ipv4MaskArr[0], ipv4MaskArr[1], ipv4MaskArr[2], ipv4MaskArr[3]).Size()
	return
}
