package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetRelayStatTable(ctx *context.Context) table.Table {

	relayStat := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := relayStat.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldSortable()
	info.AddField("Relay_report_id", "relay_report_id", db.Int4)
	info.AddField("Relay_id", "relay_id", db.Int4)
	info.AddField("Switched_count", "switched_count", db.Int4)
	info.AddField("Total_time_on", "total_time_on", db.Int4)

	info.SetTable("relay_stat").SetTitle("RelayStat").SetDescription("RelayStat")

	formList := relayStat.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Relay_report_id", "relay_report_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Relay_id", "relay_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Switched_count", "switched_count", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Total_time_on", "total_time_on", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("relay_stat").SetTitle("RelayStat").SetDescription("RelayStat")

	return relayStat
}
