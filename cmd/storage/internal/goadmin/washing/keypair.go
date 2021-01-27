package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetKeypairTable(ctx *context.Context) table.Table {

	keypair := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := keypair.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldSortable()
	info.AddField("Station_id", "station_id", db.Int4).
		FieldFilterable().
		FieldSortable()
	info.AddField("Key", "key", db.Text)
	info.AddField("Value", "value", db.Text)

	info.SetTable("keypair").SetTitle("Keypair").SetDescription("Keypair")

	formList := keypair.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Station_id", "station_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Key", "key", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Value", "value", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("keypair").SetTitle("Keypair").SetDescription("Keypair")

	return keypair
}
