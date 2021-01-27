package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetMoneyReportTable(ctx *context.Context) table.Table {

	moneyReport := table.NewDefaultTable(table.DefaultConfigWithDriver("postgresql"))

	info := moneyReport.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int4).
		FieldSortable()
	info.AddField("Station_id", "station_id", db.Int4).
		FieldFilterable().
		FieldSortable()
	info.AddField("Banknotes", "banknotes", db.Int4)
	info.AddField("Cars_total", "cars_total", db.Int4)
	info.AddField("Coins", "coins", db.Int4)
	info.AddField("Electronical", "electronical", db.Int4)
	info.AddField("Service", "service", db.Int4)
	info.AddField("Ctime", "ctime", db.Timestamp)

	info.SetTable("money_report").SetTitle("MoneyReport").SetDescription("MoneyReport")

	formList := moneyReport.GetForm()
	formList.AddField("Id", "id", db.Int4, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Station_id", "station_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Banknotes", "banknotes", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Cars_total", "cars_total", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Coins", "coins", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Electronical", "electronical", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Service", "service", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Ctime", "ctime", db.Timestamp, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("money_report").SetTitle("MoneyReport").SetDescription("MoneyReport")

	return moneyReport
}
