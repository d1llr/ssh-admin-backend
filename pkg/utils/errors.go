package utils

import "errors"

type ErrorText string

const (
	AccessTokenExpiredError  ErrorText = "Access token expired"
	RefreshTokenExpiredError ErrorText = "Refresh token expired"
	UserNotFoundError        ErrorText = "User not found"
)

func CreateError(text ErrorText) error {
	return errors.New(string(text))
}
