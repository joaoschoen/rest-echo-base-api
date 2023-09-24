package controller

import (
	"API-ECHO/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userResponse struct {
	User     model.User
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
