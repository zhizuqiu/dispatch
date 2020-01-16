package models

type Base struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Result  Result `json:"result"`
}

type Result interface {
}

func Success(result interface{}) Base {
	return Base{200, "", result}
}

func SuccessWithMessage(result interface{}, message string) Base {
	return Base{200, message, result}
}

func InternalServerError(message string) Base {
	return Base{500, message, nil}
}

func NotExistsError(message string) Base {
	return Base{404, message, nil}
}

func InternalServerErrorWithResult(message string, result Result) Base {
	return Base{500, message, result}
}

type Option interface {
}
