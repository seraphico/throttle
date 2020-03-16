package tc

import (
	"fmt"
	"os/exec"
	"strings"
)

func (tcm *TcMgr) Qdisc(dev string) (err error) {

	if !tcm.qdiscRootHandleStatus(dev) {
		//return errors.New(`root HTB not found.`)
		if err = tcm.qdiscRootHandle(dev); err != nil {
			return
		}
	}
	return
}

func (tcm *TcMgr) qdiscRootHandle(dev string) (err error) {
	var (
		cmd     *exec.Cmd
		command string
	)
	command = fmt.Sprintf(`tc qdisc add dev %s root handle %s:%s htb default %s`, dev, DOWNMAJOR, DOWNMINOR, DOWNDEFAULT)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if _, err = cmd.Output(); err != nil {
		return
	}
	return
}

func (tcm *TcMgr) qdiscRootHandleStatus(dev string) bool {

	var (
		cmd     *exec.Cmd
		command string
		out     []byte
	)
	command = fmt.Sprintf(`tc qdisc show dev %s`, dev)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	out, _ = cmd.Output()
	return strings.Contains(string(out), `htb`)
}

/*func (tc *TcMgr) QdiscShow(dev string) (rootid string, err error)  {
	var (
		cmd     *exec.Cmd
		command string
		out     []byte
	)
	command = fmt.Sprintf(`tc qdisc show dev %s`, dev)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	out, _ = cmd.Output()

	return
}
*/
