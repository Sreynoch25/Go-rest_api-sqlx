package utils

type ResponseWithPaging struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Page   int         `json:"page"`
	PerPage int        `json:"per_page"`
	Total  int         `json:"total"`
}

func ApiResponseWithPagination (success bool, message string, data interface{}, page int, perPage int, total int) ResponseWithPaging {
	return ResponseWithPaging{
		Success: success,
		Message: message,
		Data: data,
		Page: page,
		PerPage: perPage,
		Total: total,
	}
}


type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	StatusCode int      `json:"status_code"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Errors interface{} `json:"errors"`
}

func ApiResponse (success bool, message string, statusCode int, data interface{}) Response {

	switch errData := data.(type) {
		 case error: // Handle standard errors
		   return Response{
			Success:   success,
			Message:   message,
			StatusCode: statusCode,
			Data:      errData.Error(),
		   }
		 default: // Handle other data types
		 return Response{
			Success: success,
			Message: message,
			StatusCode: statusCode,
			Data: data,
		}
	}
	
}

