package handle

import (
	"github.com/gin-gonic/gin"
	"goAdmin/libary/e"
)

func GetItemList(c *gin.Context) {
	var res = e.GetRrrReturn(e.SUCCESS)

	var where = map[string]interface{}{
		"isDel": 1,
	}
	rows, err := DbCon.Table("items").Select("id,name,icon").Where(where).Rows()

	if err != nil {
		c.JSON(200, e.GetRrrReturn(e.FAILD))
		return
	}
	var resItemId int
	var resName string
	var resIcon string

	var resDatas = make([]map[string]interface{}, 0)
	for rows.Next() {
		var resData = make(map[string]interface{}, 0)
		rows.Scan(&resItemId, &resName, &resIcon)
		resData["itemId"] = resItemId
		resData["itemName"] = resName
		resData["itemIcon"] = UplodPath + resIcon
		resDatas = append(resDatas, resData)
	}

	res["data"] = resDatas
	c.JSON(200, res)
}
