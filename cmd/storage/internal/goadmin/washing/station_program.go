package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetStationProgramTable(ctx *context.Context) table.Table {

	stationProgram := table.NewDefaultTable(table.Config{
		Driver:     "postgresql",
		CanAdd:     true,
		Editable:   true,
		Deletable:  true,
		Exportable: true,
		Connection: "default",
		PrimaryKey: table.PrimaryKey{
			Type: db.Int4,
			Name: "station_id",
		},
	})

	info := stationProgram.GetInfo().HideFilterArea()

	info.AddField("Station_id", "station_id", db.Int4).
	FieldFilterable().
	FieldSortable()
	info.AddField("Program_id", "program_id", db.Int4)
	info.AddField("Name", "name", db.Text)
	info.AddField("Relays", "relays", db.Text)

	info.SetTable("station_program").SetTitle("StationProgram").SetDescription("StationProgram")

	formList := stationProgram.GetForm()
	formList.AddField("Station_id", "station_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Program_id", "program_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Name", "name", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Relays", "relays", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("station_program").SetTitle("StationProgram").SetDescription("StationProgram")

	return stationProgram
}
