package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetKasseTable(ctx *context.Context) table.Table {

	kasse := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := kasse.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldSortable()
	info.AddField("Ctime", "ctime", db.Timestamp)
	info.AddField("Receipt_item", "receipt_item", db.Text)
	info.AddField("Tax_type", "tax_type", db.Text)
	info.AddField("Cashier_full_name", "cashier_full_name", db.Text)
	info.AddField("Cashier_inn", "cashier_inn", db.Text)

	info.SetTable("kasse").SetTitle("Kasse").SetDescription("Kasse")

	formList := kasse.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Ctime", "ctime", db.Timestamp, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Receipt_item", "receipt_item", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Tax_type", "tax_type", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Cashier_full_name", "cashier_full_name", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Cashier_inn", "cashier_inn", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("kasse").SetTitle("Kasse").SetDescription("Kasse")

	return kasse
}
