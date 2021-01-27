package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetRelayReportTable(ctx *context.Context) table.Table {

	relayReport := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := relayReport.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldSortable()
	info.AddField("Station_id", "station_id", db.Int4).
		FieldFilterable().
		FieldSortable()
	info.AddField("Ctime", "ctime", db.Timestamp)

	info.SetTable("relay_report").SetTitle("RelayReport").SetDescription("RelayReport")

	formList := relayReport.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Station_id", "station_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Ctime", "ctime", db.Timestamp, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("relay_report").SetTitle("RelayReport").SetDescription("RelayReport")

	return relayReport
}
