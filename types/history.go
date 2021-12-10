package types

import (
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

type SSMHItem struct {
	StartDate    *time.Time
	EndDate      *time.Time
	Target       string
	SessionID    string
	Reason       string
	InstanceName string
}

type SSMHItemList []SSMHItem

const dateFmt = time.RFC3339

func (l SSMHItemList) Display() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"#", "Target", "Instance Name", "Session ID", "Reason", "Start Date", "End Date"})

	for i, h := range l {
		table.Append([]string{
			strconv.Itoa(i + 1),
			h.Target,
			h.InstanceName,
			h.SessionID,
			h.Reason,
			h.StartDate.In(time.Local).Format(dateFmt),
			h.EndDate.In(time.Local).Format(dateFmt),
		})
	}

	table.Render()
}
