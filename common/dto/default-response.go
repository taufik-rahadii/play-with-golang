package commondto

type BaseResponseSingle struct {
	ResponseSchema struct {
		ResponseCode    string `json:"response_code"`
		ResponseMessage string `json:"response_message"`
	} `json:"response_schema"`
	ResponseOutput struct {
		Detail interface{} `json:"detail"`
	} `json:"response_output"`
}

type BaseResponseList struct {
	ResponseSchema struct {
		ResponseCode    string `json:"response_code"`
		ResponseMessage string `json:"response_message"`
	} `json:"response_schema"`
	ResponseOutput struct {
		List struct {
			Pagination struct {
				Page  int `json:"page"`
				Size  int `json:"size"`
				Total int `json:"total"`
			} `json:"pagination"`
			Content interface{} `json:"content"`
		} `json:"list"`
	} `json:"response_output"`
}
