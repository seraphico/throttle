package tc

import (
	"fmt"
	"testing"
)

//测试一个随机classful id
func TestInitTcMgr(t *testing.T) {
	var err error

	if err = InitTcMgr(); err != nil {
		t.Error(err.Error())
	}

	fmt.Println(GlobTcMgr.random())

	fmt.Println(GlobTcMgr.hexToIp(`aca80000`))
	fmt.Println(GlobTcMgr.hexToIp(`ffffff00`))
	fmt.Println(GlobTcMgr.ipSubNetMaskToInt(GlobTcMgr.hexToIp(`ffffff00`)))
}
