package utils

import "errors"

var (
	userIsExist = "User is exist in db"
)

func UserisExistError() error {
	return errors.New(userIsExist)
}
