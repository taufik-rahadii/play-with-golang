package serializers

import commondto "taufikRahadi/fiber-boilerplate/common/dto"

func BuildResponseSingle(responseCode string, responseMessage string, data interface{}) *commondto.BaseResponseSingle {
	res := commondto.BaseResponseSingle{
		ResponseSchema: struct {
			ResponseCode    string "json:\"response_code\""
			ResponseMessage string "json:\"response_message\""
		}{ResponseCode: responseCode, ResponseMessage: responseMessage},
		ResponseOutput: struct {
			Detail interface{} "json:\"detail\""
		}{
			Detail: data,
		},
	}

	return &res
}
