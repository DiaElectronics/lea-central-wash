package washingadmin

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetStationHashTable(ctx *context.Context) table.Table {

	stationHash := table.NewDefaultTable(table.Config{
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

	info := stationHash.GetInfo().HideFilterArea()

	info.AddField("Station_id", "station_id", db.Int4).
		FieldSortable()
	info.AddField("Hash", "hash", db.Text)
	info.AddField("Ctime", "ctime", db.Timestamp)

	info.SetTable("station_hash").SetTitle("StationHash").SetDescription("StationHash")

	formList := stationHash.GetForm()
	formList.AddField("Station_id", "station_id", db.Int4, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Hash", "hash", db.Text, form.RichText).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Ctime", "ctime", db.Timestamp, form.Datetime).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("station_hash").SetTitle("StationHash").SetDescription("StationHash")

	return stationHash
}
