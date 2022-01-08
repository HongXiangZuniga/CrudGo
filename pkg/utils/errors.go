package utils

import "errors"

type Pagination struct {
	Last    int `json:"last"`
	Next    int `json:"next"`
	Current int `json:"current"`
	Total   int `json:"total"`
}

var (
	userIsExist    = "User is exist in db"
	pagateNotValid = "Page not Valid"
)

func UserisExistError() error {
	return errors.New(userIsExist)
}

func PageNotValid() error {
	return errors.New(pagateNotValid)
}
