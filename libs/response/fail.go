package response

import (
	type_response "github.com/rodericusifo/football-club-profile/libs/response/types"
)

func ResponseFail(message string, err error) type_response.ResponseError {
	return type_response.ResponseError{
		Success: false,
		Message: message,
		Error:   err,
	}
}
