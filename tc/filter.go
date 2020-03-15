package tc

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strconv"
	"strings"
)

func (tcm *TcMgr) Filter(dev string, qdiscid string, dst string, flowid string) (err error) {
	if !tcm.filterStatus(dev, flowid) {
		if err = tcm.filterCreate(dev, qdiscid, dst, flowid); err != nil {
			return
		}
	}
	return
}

func (tcm *TcMgr) filterCreate(dev string, qdiscid string, dst string, flowid string) (err error) {
	var (
		cmd     *exec.Cmd
		command string
	)
	//`tc filter add  dev $1 protocol ip parent 2:0  u32 match ip dst 192.168.1.0/24  flowid 2:10`
	command = fmt.Sprintf(`tc filter add  dev %s protocol ip parent %s  u32 match ip dst %s  flowid %s`,
		dev,
		qdiscid,
		dst,
		flowid,
	)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if _, err = cmd.Output(); err != nil {
		return
	}
	return
}

func (tcm *TcMgr) filterStatus(dev, classid string) (t bool) {
	var (
		cmd     *exec.Cmd
		command string
		out     []byte
	)
	command = fmt.Sprintf(`tc filter show dev %s`, dev)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	out, _ = cmd.Output()
	return strings.Contains(string(out), classid)
}

func (tcm *TcMgr) FilterShow(dev string) (err error) {
	var (
		//readsize       int
		bufs           *bufio.Reader
		cmd            *exec.Cmd
		command        string
		out            io.ReadCloser //[]byte
		cmdoutputlines [][]string
	)
	command = fmt.Sprintf(`tc filter show dev %s`, dev)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if out, err = cmd.StdoutPipe(); err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	bufs = bufio.NewReader(out)
	for {
		var cmdoutputline []string
		var rootHandlId, classId, ipaddrs, netmasks string
		line, errread := bufs.ReadString('\n')
		if line != " " {
			if strings.Contains(line, `flowid`) {
				lines := strings.Split(line, " ")
				//cmdoutputline = append(cmdoutputline, lines[2], lines[20])
				rootHandlId = lines[2]
				classId = lines[20]
			}

			if strings.Contains(line, `match`) {
				lines := strings.Split(line, " ")
				netiplines := strings.Split(lines[3], `/`)
				ipaddr := tcm.hexToIp(netiplines[0])
				netmaskint, e := tcm.ipSubNetMaskToInt(tcm.hexToIp(netiplines[1]))
				if e != nil {
					return
				}
				ipaddrs = ipaddr
				netmasks = strconv.Itoa(netmaskint)
				//cmdoutputline = append(cmdoutputline,rootHandlId, ipaddr, strconv.Itoa(netmaskint))
			}
			cmdoutputline = append(cmdoutputline, rootHandlId, classId, ipaddrs, netmasks)
			fmt.Println(cmdoutputline)
		}
		cmdoutputlines = append(cmdoutputlines, cmdoutputline)

		if errread != nil || line == " " {
			break
		}
	}
	fmt.Println(cmdoutputlines)
	return
}
func (tcm *TcMgr) FilterShows(dev string) (data [][]string, err error) {
	var (
		// readsize       int
		bufs           *bufio.Reader
		cmd            *exec.Cmd
		command        string
		out            io.ReadCloser // []byte
		cmdoutputlines []string
	)
	command = fmt.Sprintf(`tc filter show dev %s`, dev)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if out, err = cmd.StdoutPipe(); err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	bufs = bufio.NewReader(out)

	for {
		line, e := bufs.ReadString('\n')
		if strings.Contains(line, `flowid`) {
			nlines := strings.Split(line, ` `)
			cmdoutputlines = append(cmdoutputlines, nlines[6])
		}
		if e != nil || len(line) == 0 {
			break
		}
	}

	for _, pref := range cmdoutputlines {
		lindata, e := tcm.filterShow(dev, pref)
		if e != nil {
			err = e
			return
		}
		data = append(data, lindata)
	}

	return
}
func (tcm *TcMgr) filterShow(dev string, prefid string) (li []string, err error) {
	var (
		// readsize       int
		cmd     *exec.Cmd
		command string
		out     io.ReadCloser
		bufs    *bufio.Reader
	)
	command = fmt.Sprintf(`tc filter show dev %s pref %s`, dev, prefid)
	cmd = exec.Command(`/usr/bin/sh`, `-c`, command)
	if out, err = cmd.StdoutPipe(); err != nil {
		return
	}
	if err = cmd.Start(); err != nil {
		return
	}
	bufs = bufio.NewReader(out)
	for {
		line, e := bufs.ReadString('\n')
		if len(line) != 0 {
			if strings.Contains(line, `flowid`) {
				lis := strings.Split(line, " ")
				ratem, e := tcm.classRate(dev, lis[18])
				if e != nil {
					err = e
					return
				}
				li = append( li, lis[2],lis[18], prefid, ratem[`rate`],  ratem[`ceil`])

			}
			if strings.Contains(line, `match`) {
				mlis := strings.Split(line, " ")
				mlist := strings.Split(mlis[3], `/`)
				netmask, e := tcm.ipSubNetMaskToInt(tcm.hexToIp(mlist[1]))
				if e != nil {
					err = e
					return
				}
				li = append(li, tcm.hexToIp(mlist[0]) + `/` +strconv.Itoa(netmask))
			}

		}
		if e != nil || len(line) == 0 {
			break
		}
	}
	return
}
