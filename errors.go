package newsapi

import "fmt"

type apiError struct {
	Code    string
	Message string
}

func (e *apiError) Error() string {
	return fmt.Sprintf("%s - %s", e.Code, e.Message)
}
