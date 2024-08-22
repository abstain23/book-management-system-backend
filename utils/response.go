package utils

func SuccessResponse(data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code": 0,
		"msg":  "success",
		"data": data,
	}
}
