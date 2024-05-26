package pkg

import (
	"astro_drone/app/constant"
	"astro_drone/app/domain/dto"
)

func Null() interface{} {
	return nil
}

func BuildResponse[T any](responseStatus constant.ResponseStatus, data T) dto.ApiResponseDTO[T] {
	return BuildResponse_(responseStatus.GetResponseStatus(), responseStatus.GetResponseMessage(), data)
}

func BuildResponse_[T any](status string, message string, data T) dto.ApiResponseDTO[T] {
	return dto.ApiResponseDTO[T]{
		ResponseKey:     status,
		ResponseMessage: message,
		Data:            data,
	}
}
