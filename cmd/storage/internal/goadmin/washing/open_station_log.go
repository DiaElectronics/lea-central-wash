package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetOpenStationLogTable(ctx *context.Context) table.Table {

	openStationLog := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := openStationLog.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldSortable()
	info.AddField("Station_id", "station_id", db.Int4).
		FieldFilterable().
		FieldSortable()
	info.AddField("Ctime", "ctime", db.Timestamp)

	info.SetTable("open_station_log").SetTitle("OpenStationLog").SetDescription("OpenStationLog")

	formList := openStationLog.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Station_id", "station_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Ctime", "ctime", db.Timestamp, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("open_station_log").SetTitle("OpenStationLog").SetDescription("OpenStationLog")

	return openStationLog
}
