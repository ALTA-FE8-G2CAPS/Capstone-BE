package helper

import "net/http"

func Fail_Resp(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Failed",
		"message": msg,
	}

}

func Success_Resp(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": msg,
	}

}

func Success_DataResp(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": msg,
		"data":    data,
		"code":    http.StatusOK,
	}

}

func Success_Login(msg string, data interface{}, data2 interface{}, data3 interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "Success",
		"message": msg,
		"token":   data,
		"role":    data2,
		"user":    data3,
		"code":    http.StatusOK,
	}

}
