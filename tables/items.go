package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"jiuhuo/handle"
)

func GetItemsTable(ctx *context.Context) table.Table {

	items := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := items.GetInfo().HideFilterArea()

	// 字段显示
	info.AddField("Id", "id", db.Int).FieldFilterable()
	info.AddField("名称", "name", db.Varchar)
	info.AddField("图标", "icon", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "" {
			return "暂无"
		} else {
			return "<img height=\"30px\" src=\"" + handle.UplodPath + model.Value + "\" />"
		}

	})
	info.AddField("是否禁用", "isDel", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "1" {
			return "否"
		}
		if model.Value == "2" {
			return "是"
		}
		return "未知"
	})
	//info.AddField("CreateTime", "createTime", db.Int).FieldDisplay(func(model types.FieldModel) interface{} {
	//	var strV = model.Value + "000"
	//	return timer.ConversionTimeMSToDate(tool.StrToInt64(strV), timer.DefaultDatetimeFormat)
	//})
	//info.AddField("UpdateTime", "updateTime", db.Int).FieldDisplay(func(model types.FieldModel) interface{} {
	//	if model.Value == "0" {
	//		return "/"
	//	}
	//	var strV = model.Value + "000"
	//	return timer.ConversionTimeMSToDate(tool.StrToInt64(strV), timer.DefaultDatetimeFormat)
	//})

	info.SetTable("items").SetTitle("Items").SetDescription("Items")

	formList := items.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("名称", "name", db.Varchar, form.Text)
	formList.AddField("图标", "icon", db.Varchar, form.File)
	formList.AddField("是否禁用", "isDel", db.Tinyint, form.SelectSingle).
		FieldOptions(types.FieldOptions{
			{Text: "否", Value: "1"},
			{Text: "是", Value: "2"},
		}).
		FieldDefault("1")
	formList.AddField("CreateTime", "createTime", db.Int, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("UpdateTime", "updateTime", db.Int, form.Number).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()

	formList.SetTable("items").SetTitle("Items").SetDescription("Items")

	return items
}
