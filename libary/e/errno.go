package e

import "github.com/gin-gonic/gin"

type Err struct {
	Data struct {
	} `json:"data"`
	Errmsg string `json:"errmsg"`
	Errno  int    `json:"errno"`
}

func (f *Err) MakeErr(errno int) Err {
	f.Errno = errno
	f.Errmsg = errMsg[errno]
	return *f
}

const (
	SUCCESS       = 1
	FAILD         = 2
	InvalidParams = 101
)

var errMsg = map[int]string{
	SUCCESS:       "success",
	FAILD:         "faild",
	InvalidParams: "Invalid Params",
}

func GetRrrReturn(errno int) gin.H {
	errReturn := gin.H{
		"errno":  errno,
		"errmsg": errMsg[errno],
		"data":   make(map[string]interface{}),
	}
	return errReturn
}
