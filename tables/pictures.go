package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types/form"
)

func GetPicturesTable(ctx *context.Context) table.Table {

	pictures := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := pictures.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("ItemId", "itemId", db.Int)
	info.AddField("PicUrl", "picUrl", db.Varchar)
	info.AddField("IsDel", "isDel", db.Tinyint)
	info.AddField("CreateTime", "createTime", db.Int)
	info.AddField("UpdateTime", "updateTime", db.Int)

	info.SetTable("pictures").SetTitle("Pictures").SetDescription("Pictures")

	formList := pictures.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("ItemId", "itemId", db.Int, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("PicUrl", "picUrl", db.Varchar, form.Text).
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

	formList.SetTable("pictures").SetTitle("Pictures").SetDescription("Pictures")

	return pictures
}
