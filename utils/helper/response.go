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

func Success_Login(msg string, data, data2, data3, data4, data5, data6 interface{}) map[string]interface{} {
	return map[string]interface{}{
		"message":    msg,
		"token":      data,
		"user_id":    data5,
		"user":       data3,
		"role":       data2,
		"foto_user":  data6,
		"user_owner": data4,
		"code":       http.StatusOK,
	}

}
