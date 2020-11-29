package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetItemsTable(ctx *context.Context) table.Table {

	items := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := items.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("Name", "name", db.Varchar)
	info.AddField("Icon", "icon", db.Varchar)
	info.AddField("IsDel", "isDel", db.Tinyint)
	info.AddField("CreateTime", "createTime", db.Int)
	info.AddField("UpdateTime", "updateTime", db.Int)

	info.SetTable("items").SetTitle("Items").SetDescription("Items")

	formList := items.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Name", "name", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("Icon", "icon", db.Varchar, form.Text).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("IsDel", "isDel", db.Tinyint, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("CreateTime", "createTime", db.Int, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("UpdateTime", "updateTime", db.Int, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("items").SetTitle("Items").SetDescription("Items")

	return items
}