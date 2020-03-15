package tc

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func (tcm *TcMgr) Classful(dev, classid, rate, ceil string) (err error) {
	if !tcm.classFulStatus(dev, classid) {
		if err = tcm.classfulcreate(dev, classid, rate, ceil); err != nil {
			return
		}
	}
	return
}

func (tcm *TcMgr) ClassfulShow(dev string) (err error) {
	var (
		//readsize       int
		bufs           *bufio.Reader
		cmd            *exec.Cmd
		command        string
		out            io.ReadCloser //[]byte
		cmdoutputlines []string
	)
	command = fmt.Sprintf(`tc class show dev %s`, dev)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if out, err = cmd.StdoutPipe(); err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	bufs = bufio.NewReader(out)
	for {
		line, errread := bufs.ReadString('\n')
		if line != "" {
			cmdoutputlines = append(cmdoutputlines, line)
		}

		if errread != nil || line == "" {
			break
		}
	}
	fmt.Println(cmdoutputlines)
	return
}

func (tcm *TcMgr) classfulcreate(dev, classid, rate, ceil string) (err error) {
	var (
		cmd    string
		cmdOut *exec.Cmd
	)

	cmd = fmt.Sprintf(`tc class add dev %s parent %s: classid %s htb rate %smbit ceil %smbit`,
		dev,
		MAJOR,
		classid,
		rate,
		ceil,
	)
	cmdOut = exec.Command(`/usr/bin/sh`, `-c`, cmd)
	if _, err = cmdOut.Output(); err != nil {
		return
	}
	return
}

func (tcm *TcMgr) classFulStatus(dev, classid string) (e bool) {

	var (
		cmd     *exec.Cmd
		command string
		out     []byte
	)
	command = fmt.Sprintf(`tc class show dev %s`, dev)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	out, _ = cmd.Output()
	return strings.Contains(string(out), classid)
}

func (tcm *TcMgr) classRate(dev, classid string) (rs map[string]string, err error) {

	var (
		cmd        *exec.Cmd
		command    string
		out        []byte
		outstrlist []string
	)
	rs = make(map[string]string, 0)
	command = fmt.Sprintf(`tc class show dev %s classid %s`, dev, classid)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if out, err = cmd.Output(); err != nil {
		return
	}
	outstrlist = strings.Split(string(out), " ")

	rs[`rate`] = outstrlist[7]
	rs[`ceil`] = outstrlist[9]
	return
}
