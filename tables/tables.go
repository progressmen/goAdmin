// This file is generated by GoAdmin CLI adm.
package tables

import "github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"

// The key of Generators is the prefix of table info url.
// The corresponding value is the Form and Table data.
//
// http://{{config.Domain}}:{{Port}}/{{config.Prefix}}/info/{{key}}
//
// example:
//
// "items" => http://localhost:9033/admin/info/items
// "pictures" => http://localhost:9033/admin/info/pictures
//
// example end
//
var Generators = map[string]table.Generator{

	"items":    GetItemsTable,
	"pictures": GetPicturesTable,

	// generators end
}