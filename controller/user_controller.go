package controller

import (
	"API-ECHO/model"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type userResponse struct {
	User     model.User
	Response string
}

type userListResponse struct {
	UserList []model.User
	Response string
}

func GetUser(c echo.Context) error {
	// PARAM
	var id string
	id = c.Param("id")

	// DUMMY DATA
	user := model.User{
		Email:    "jon@doe.com",
		Password: "badPasswordExample",
	}

	// DATABASE REQUEST GOES HERE

	//BUILD RESPONSE
	response := userResponse{
		User:     user,
		Response: fmt.Sprint("id to get:", id, " WARN: this is dummy data"),
	}

	return c.JSON(http.StatusOK, response)
}

func GetUserList(c echo.Context) error {

	// QUERY
	email := c.QueryParam("email")
	// PAGING
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Page must be a number, received:%d", page))
	}
	if page < 0 {
		return c.JSON(http.StatusBadRequest, fmt.Sprintf("Page number must be positive, received:%d", page))
	}
	PAGE_SIZE := 2
	START := PAGE_SIZE * page
	END := START + PAGE_SIZE
	// DATABASE REQUEST GOES HERE

	// DUMMY DATA
	userList := []model.User{
		{
			Email:    "jon1@doe.com",
			Password: "badPasswordExample",
		},
		{
			Email:    "jon2@doe.com",
			Password: "badPasswordExample",
		},
		{
			Email:    "dave1@doe.com",
			Password: "badPasswordExample",
		},
		{
			Email:    "dave2@doe.com",
			Password: "badPasswordExample",
		},
		{
			Email:    "dave3@doe.com",
			Password: "badPasswordExample",
		},
	}

	var filteredList []model.User

	for i := range userList {
		if strings.Contains(userList[i].Email, email) {
			filteredList = append(filteredList, userList[i])
		}
	}
	if END > len(filteredList) {
		END = len(filteredList)
	}
	filteredList = filteredList[START:END]
	//BUILD RESPONSE
	response := userListResponse{
		UserList: filteredList,
		Response: fmt.Sprint("WARN: this is dummy data"),
	}

	return c.JSON(http.StatusOK, response)
}

func PostUser(c echo.Context) error {
	// BODY
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	// DATABASE REQUEST GOES HERE

	// BUILD RESPONSE
	response := userResponse{
		User:     user,
		Response: "user to create",
	}
	return c.JSON(http.StatusOK, response)
}

func PutUser(c echo.Context) error {
	// PARAM
	var id string
	id = c.Param("id")

	// BODY
	var user model.User
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	// DATABASE REQUEST GOES HERE

	//BUILD RESPONSE
	response := userResponse{
		User:     user,
		Response: fmt.Sprint("id to update:", id),
	}
	return c.JSON(http.StatusOK, response)
}

func DeleteUser(c echo.Context) error {
	// PARAM
	var id string
	id = c.Param("id")

	// DATABASE REQUEST GOES HERE

	//BUILD RESPONSE
	response := fmt.Sprint("id to delete:", id)
	return c.JSON(http.StatusOK, response)
}
