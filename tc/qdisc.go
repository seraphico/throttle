package tc

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func (self *TcMgr) Qdisc(dev string) (err error) {
	if err = self.qdiscRootHandleStatus(dev); err != nil {
		return
	}

	if err = self.qdiscRootHandle(dev); err != nil {
		return
	}

	return
}

func (self *TcMgr) qdiscRootHandle(dev string) (err error) {
	var (
		cmd     *exec.Cmd
		command string
		out     []byte
	)
	command = fmt.Sprintf(`tc qdisc add dev %s root handle %s:%s htb default %s`, dev, MAJOR, MINOR, DEFAULT)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if out, err = cmd.Output(); err != nil {
		return
	}

	fmt.Println(string(out))

	return
}

func (self *TcMgr) qdiscRootHandleStatus(dev string) (err error) {

	var (
		cmd     *exec.Cmd
		command string
		out     []byte
	)
	command = fmt.Sprintf(` tc  qdisc show dev %s`, dev)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if out, err = cmd.Output(); err != nil {
		return
	}
	for _, st := range strings.Split(string(out), " ") {
		if st == `htb` {
			err = errors.New(`htb root qdisc not found`)
			break
		}
	}

	return
}
