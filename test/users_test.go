package test

import (
	"errors"
	"testing"
	"time"

	users "github.com/HongXiangZuniga/CrudGoExample/pkg/Users"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mockMongoRepo struct {
}

func (m *mockMongoRepo) GetAllUser() (*[]users.User, error) {
	users := []users.User{
		{Id: 1, Name: "Hong Xiang", Email: "hongxiang17@gmail.com", Age: 28, Country: "Chile", EntryDate: primitive.NewDateTimeFromTime(time.Now())},
		{Id: 2, Name: "Andres Luna", Email: "al@gmail.com", Age: 30, Country: "Chile", EntryDate: primitive.NewDateTimeFromTime(time.Now())},
	}
	return &users, nil
}
func (m *mockMongoRepo) FindUserById(id int) (*users.User, error) {
	if id == 1 {
		return &users.User{Id: 1, Name: "Hong Xiang", Email: "hongxiang17@gmail.com", Age: 28, Country: "Chile", EntryDate: primitive.NewDateTimeFromTime(time.Now())}, nil
	} else if id == 2 {
		return &users.User{Id: 2, Name: "Andres Luna", Email: "al@gmail.com", Age: 30, Country: "Chile", EntryDate: primitive.NewDateTimeFromTime(time.Now())}, nil
	}
	return nil, errors.New("User not found")
}
func (m *mockMongoRepo) FindUsersByStringField(field, value string) (*[]users.User, error) {
	var resultUsers []users.User
	userHong := &users.User{Id: 1, Name: "Hong Xiang", Email: "hongxiang17@gmail.com", Age: 28, Country: "Chile", EntryDate: primitive.NewDateTimeFromTime(time.Now())}
	userAndres := &users.User{Id: 2, Name: "Andres Luna", Email: "al@gmail.com", Age: 30, Country: "Chile", EntryDate: primitive.NewDateTimeFromTime(time.Now())}
	switch field {
	case "Name":
		if value == "Hong Xiang" {
			resultUsers = append(resultUsers, *userHong)
			return &resultUsers, nil
		} else if value == "Andres Luna" {
			resultUsers = append(resultUsers, *userAndres)
			return &resultUsers, nil
		} else {
			return nil, nil
		}
	case "Email":
		if value == "hongxiang17@gmail.com" {
			resultUsers = append(resultUsers, *userHong)
			return &resultUsers, nil
		} else if value == "al@gmail.com" {
			resultUsers = append(resultUsers, *userAndres)
			return &resultUsers, nil
		} else {
			return nil, nil
		}
	case "Age":
		if value == "28" {
			resultUsers = append(resultUsers, *userHong)
			return &resultUsers, nil

		} else if value == "30" {
			resultUsers = append(resultUsers, *userAndres)
			return &resultUsers, nil
		} else {
			return nil, nil
		}
	case "Country":
		if value == "Chile" {
			resultUsers = append(resultUsers, *userHong)
			resultUsers = append(resultUsers, *userAndres)
			return &resultUsers, nil
		}
	}
	return nil, nil
}

// NoTest
func (m *mockMongoRepo) CreateUser(newUser users.User) error {
	return nil
}

func (m *mockMongoRepo) DeleteUser(idUser int) error {
	return nil
}

func TestGetAllUsers(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file %s", err.Error())
	}
	userServices := users.NewUserServices(&mockMongoRepo{})
	type testCase struct {
		name    []string
		isFaill bool
	}
	testCases := []testCase{
		{[]string{"Hong Xiang", "Andres Luna"}, false},
		{[]string{"Faill 1", "Faill 2"}, true},
	}
	for _, tc := range testCases {
		if tc.isFaill == true {
			resultUsers, err := userServices.GetAllUser(0)
			if err != nil {
				t.Errorf("Error in GetAllUsers:%s", err.Error())
			}
			for i, user := range *resultUsers {
				if user.Name == tc.name[i] {
					t.Fatalf("Name is %s but excepted is %s", user.Name, tc.name[i])
				}
			}

		} else {
			resultUsers, err := userServices.GetAllUser(0)
			if err != nil {
				t.Errorf("Error in GetAllUsers:%s", err.Error())
			}
			for i, user := range *resultUsers {
				if user.Name != tc.name[i] {
					t.Fatalf("Name is %s but excepted is %s", user.Name, tc.name[i])
				}
			}
		}

	}
}

func TestFindUserById(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error loading .env file %s", err.Error())
	}
	userServices := users.NewUserServices(&mockMongoRepo{})
	type testCase struct {
		id      int
		name    string
		isFaill bool
	}
	testCases := []testCase{
		{1, "Hong Xiang", false},
		{2, "Andres Luna", false},
		{1, "Andres Luna", true},
	}
	for _, tc := range testCases {
		if tc.isFaill == true {
			result, err := userServices.GetUserById(tc.id)
			if err != nil {
				t.Errorf("Error in get user by id:%s", err.Error())
			}
			if result.Name == tc.name {
				t.Errorf("Name is %s but excepted not equals to %s", result.Name, tc.name)
			}

		} else {
			result, err := userServices.GetUserById(tc.id)
			if err != nil {
				t.Errorf("Error in get user by id:%s", err.Error())
			}
			if result.Name != tc.name {
				t.Errorf("Name is %s but excepted %s", result.Name, tc.name)
			}
		}
	}

}
