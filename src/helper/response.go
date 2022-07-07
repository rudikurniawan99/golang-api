package helper

type (
	successJson struct {
		Status  string      `json:"status"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}
)

// func JsonSuccess(c gin.Context, data interface{}) error {
// 	res := successJson{
// 		Status:  "success",
// 		Message: "success",
// 		Data:    data,
// 	}

// 	return c.JSON(http.StatusOK, res)
// }
