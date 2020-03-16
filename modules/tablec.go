package modules

import (
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type TableMgr struct {
	table *tablewriter.Table
}

var GlobTableMgr *TableMgr

func InitTableMgr() (err error) {

	var tables *tablewriter.Table
	tables = tablewriter.NewWriter(os.Stdout)

	GlobTableMgr = &TableMgr{
		table: tables,
	}

	return
}

func (t *TableMgr) Out(data [][]string) (err error)  {

	t.table.SetHeader([]string{`qdiscId`, `classId`, `filterId`, `rate`, `ceil`, `Address segment`})
	t.table.SetFooter([]string{``,``,``,``, `Total Filters`, strconv.Itoa(len(data))})
	t.table.SetAutoMergeCells(true)
	t.table.SetRowLine(true)
	t.table.AppendBulk(data)
	t.table.Render()
	return
}
