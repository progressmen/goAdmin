package tables

import (
	"github.com/GoAdminGroup/go-admin/context"
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/GoAdminGroup/go-admin/plugins/admin/modules/table"
	"github.com/GoAdminGroup/go-admin/template/types"
	"github.com/GoAdminGroup/go-admin/template/types/form"
	"goAdmin/handle"
	"goAdmin/libary/util/tool"
)

func GetPicturesTable(ctx *context.Context) table.Table {

	pictures := table.NewDefaultTable(table.DefaultConfigWithDriver("mysql"))

	// 获取item信息
	fops := getItemInfo()

	info := pictures.GetInfo().HideFilterArea()

	info.AddField("Id", "id", db.Int).
		FieldFilterable()
	info.AddField("分类ID", "itemId", db.Int).FieldDisplay(func(model types.FieldModel) interface{} {
		for _, v := range fops {
			if model.Value == v.Value {
				return v.Text
			}
		}
		return model.Value
	})
	info.AddField("图片", "picUrl", db.Varchar).FieldDisplay(func(model types.FieldModel) interface{} {
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
	//info.AddField("CreateTime", "createTime", db.Int)
	//info.AddField("UpdateTime", "updateTime", db.Int)

	info.SetTable("pictures").SetTitle("Pictures").SetDescription("Pictures")

	formList := pictures.GetForm()
	formList.AddField("Id", "id", db.Int, form.Default).
		FieldDisableWhenCreate().
		FieldDisableWhenUpdate()
	formList.AddField("分类", "itemId", db.Int, form.SelectSingle).
		FieldOptions(fops)
	formList.AddField("图片", "picUrl", db.Varchar, form.File)
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

	formList.SetTable("pictures").SetTitle("Pictures").SetDescription("Pictures")

	return pictures
}

// 获取item表信息
func getItemInfo() types.FieldOptions {
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
	return fops
}
