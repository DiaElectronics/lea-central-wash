package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetStationTable(ctx *context.Context) table.Table {

	station := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := station.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldFilterable().
		FieldSortable()
	info.AddField("Name", "name", db.Text)
	info.AddField("Deleted", "deleted", db.Bool)

	info.SetTable("station").SetTitle("Station").SetDescription("Station")

	formList := station.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Name", "name", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Deleted", "deleted", db.Bool, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("station").SetTitle("Station").SetDescription("Station")

	return station
}
