package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"jiuhuo/handle"
	"jiuhuo/libary/util/tool"
)

func GetPicturesTable(ctx *context.Context) table.Table {

	pictures := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	info := pictures.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("ItemId", "itemId", db.Int)
	info.AddField("PicUrl", "picUrl", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
		return "<img height=\"30px\" src=\"" + handle.UplodPath + model.Value + "\" />"
	})
	info.AddField("IsDel", "isDel", db.Tinyint).FieldDisplay(func(model types.FieldModel) interface{} {
		if model.Value == "1" {
			return "否"
		}
		if model.Value == "2" {
			return "是"
		}
		return "未知"
	})
	//info.AddField("CreateTime", "createTime", db.Int)
	//info.AddField("UpdateTime", "updateTime", db.Int)

	info.SetTable("pictures").SetTitle("Pictures").SetDescription("Pictures")

	// 获取item信息
	where := map[string]interface{}{
		"isDel": 1,
	}
	rows, _ := handle.DbCon.Table("items").Select("id,name").Where(where).Rows()
	var resItemId int
	var resName string
	var fops types.FieldOptions
	for rows.Next() {
		var fop types.FieldOption
		rows.Scan(&resItemId, &resName)
		fop.Text = resName
		fop.Value = tool.IntToStr(resItemId)
		fops = append(fops, fop)
	}

	formList := pictures.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("ItemId", "itemId", db.Int, form.SelectSingle).
		FieldOptions(fops)
	formList.AddField("PicUrl", "picUrl", db.Varchar, form.File)
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
