package handle

import (
	"github.com/gin-gonic/gin"
	"goAdmin/libary/e"
)

type Item struct {
	Id       int
	Name     string
	Icon     string
	Pictures []Picture
}

type Picture struct {
	Id     int
	ItemId int
	picUlr string
}

type inputGetPicList struct {
	ItemId string `form:"itemId" binding:"required"`
}

func GetPicList(c *gin.Context) {
	var res = e.GetRrrReturn(e.SUCCESS)
	var inputParams inputGetPicList
	if err := c.ShouldBind(&inputParams); err != nil {
		c.JSON(200, e.GetRrrReturn(e.InvalidParams))
		return
	}

	var where = map[string]interface{}{
		"items.id":       inputParams.ItemId,
		"items.isDel":    1,
		"pictures.isDel": 1,
	}
	rows, err := DbCon.Table("items").Select("items.id as itemId, pictures.id as picId, pictures.picUrl").Joins("inner join pictures on pictures.itemId = items.id").Where(where).Rows()

	if err != nil {
		c.JSON(200, e.GetRrrReturn(e.FAILD))
		return
	}
	var resItemId int
	var resPicId int
	var resPicUrl string

	var resDatas = make([]map[string]interface{}, 0)
	for rows.Next() {
		var resData = make(map[string]interface{}, 0)
		rows.Scan(&resItemId, &resPicId, &resPicUrl)
		resData["itemId"] = resItemId
		resData["picId"] = resPicId
		resData["picUrl"] = UplodPath + resPicUrl
		resDatas = append(resDatas, resData)
	}

	res["data"] = resDatas
	c.JSON(200, res)
}
