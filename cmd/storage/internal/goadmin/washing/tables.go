// This file is generated by GoAdmin CLI adm.
package washingadmin

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "station" => http://localhost:9033/admin/info/station
// "keypair" => http://localhost:9033/admin/info/keypair
// "money_report" => http://localhost:9033/admin/info/money_report
// "relay_report" => http://localhost:9033/admin/info/relay_report
// "relay_stat" => http://localhost:9033/admin/info/relay_stat
// "money_collection" => http://localhost:9033/admin/info/money_collection
// "station_hash" => http://localhost:9033/admin/info/station_hash
// "open_station_log" => http://localhost:9033/admin/info/open_station_log
// "station_program" => http://localhost:9033/admin/info/station_program
// "kasse" => http://localhost:9033/admin/info/kasse
//
// example end
var Generators = map[string]table.Generator{

	"station":          GetStationTable,
	"keypair":          GetKeypairTable,
	"money_report":     GetMoneyReportTable,
	"relay_report":     GetRelayReportTable,
	"relay_stat":       GetRelayStatTable,
	"money_collection": GetMoneyCollectionTable,
	"station_hash":     GetStationHashTable,
	"open_station_log": GetOpenStationLogTable,
	"station_program":  GetStationProgramTable,
	"kasse":            GetKasseTable,

	// generators end
}
