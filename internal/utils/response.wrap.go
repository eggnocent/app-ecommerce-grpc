package utils

import "github/eggnocent/app-grpc-eccomerce/pb/common"

func SuccessResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 200,
		Message:    message,
	}
}

func BadRequestResponse(message string) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode: 400,
		Message:    message,
		IsError:    true,
	}
}

func ValidationErrorResponse(validationsErrors []*common.ValidationError) *common.BaseResponse {
	return &common.BaseResponse{
		StatusCode:       400,
		Message:          "validation error",
		IsError:          true,
		ValidationErrors: validationsErrors,
	}
}
