package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMoneyCollectionTable(ctx *context.Context) table.Table {

	moneyCollection := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := moneyCollection.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldSortable()
	info.AddField("Station_id", "station_id", db.Int4).
		FieldFilterable().
		FieldSortable()
	info.AddField("Money", "money", db.Int4)
	info.AddField("Ctime", "ctime", db.Timestamp)

	info.SetTable("money_collection").SetTitle("MoneyCollection").SetDescription("MoneyCollection")

	formList := moneyCollection.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Station_id", "station_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Money", "money", db.Int4, form.Currency).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Ctime", "ctime", db.Timestamp, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("money_collection").SetTitle("MoneyCollection").SetDescription("MoneyCollection")

	return moneyCollection
}
