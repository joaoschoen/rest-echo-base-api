package controller

import (
	"API-ECHO/model"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// @Summary		Get user data
// @Description	Receives ID by request param and retreives user data
// @Tags			user
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	model.GetUserResponse
// @Failure		404	"User not found."
// @Router			/user/{id} [get]
func GetUser(c echo.Context) error {
	// PARAM
	var id string
	id = c.Param("id")

	// DUMMY DATA
	user := model.SafeUser{
		ID:    "someID",
		Email: "jon@doe.com",
	}

	/**
		DATABASE REQUEST GOES HERE
	**/

	if id == "404" {
		return c.JSON(http.StatusNotFound, "User not found.")
	}

	//BUILD RESPONSE
	response := model.GetUserResponse{
		Data: user,
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary		Get user list
// @Description	Can receive email as a query filter
// @Description	This route is paged, it requrires a page number to operate if none is received, will return page 0
// @Tags			user
// @Produce		json
// @QueryParam			email	path		string	true	"Email filter"
// @QueryParam			page	path		int		true	"Page"
// @Success		200		{object}	model.GetUserListResponse
// @Failure		500	"Internal server error"
// @Router			/user/list [get]
func GetUserList(c echo.Context) error {
	// QUERY
	email := c.QueryParam("email")
	// PAGING
	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = 0
	}
	if page < 0 {
		page = 0
	}

	PAGE_SIZE := 2
	START := PAGE_SIZE * page
	END := START + PAGE_SIZE
	// DATABASE REQUEST GOES HERE

	// DUMMY DATA
	userList := []model.SafeUser{
		{
			Email: "jon1@doe.com",
			ID:    "someID1",
		},
		{
			Email: "jon2@doe.com",
			ID:    "someID2",
		},
		{
			Email: "dave1@doe.com",
			ID:    "someID3",
		},
		{
			Email: "dave2@doe.com",
			ID:    "someID4",
		},
		{
			Email: "dave3@doe.com",
			ID:    "someID5",
		},
	}
	totalPages := int(math.Ceil(float64(len(userList) / 2)))
	var filteredList []model.SafeUser
	if email != "" {
		for i := range userList {
			if strings.Contains(userList[i].Email, email) {
				filteredList = append(filteredList, userList[i])
			}
		}
	}
	if END > len(filteredList) {
		END = len(filteredList)
	}
	filteredList = filteredList[START:END]
	//BUILD RESPONSE
	response := model.GetUserListResponse{
		Data: filteredList,
		Paging: model.Paging{
			Page:  page,
			Total: totalPages,
		},
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary		Create new user
// @Description	Receives user email and password, returns UUID
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			email	path		string	true	"User email"
// @Success		200		{object}	model.PostUserResponse
// @Failure		400 "Email already in use"
// @Failure		404	"User not found."
// @Failure		500	"Internal server error"
// @Router			/user [post]
func PostUser(c echo.Context) error {
	// BODY
	var user model.UnsafeUser
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error while parsing received data")
	}
	// DATABASE REQUEST GOES HERE

	// CHECK DUPLICATE
	if user.Email == "alreadyIn@use.com" {
		return c.JSON(http.StatusBadRequest, "Email already in use")
	}

	// BUILD RESPONSE
	response := model.PostUserResponse{
		ID: "somerandomid",
	}
	return c.JSON(http.StatusOK, response)
}

// @Summary		Update user
// @Description	Receives updated user object, returns updated object
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200		{object}	model.PostUserResponse
// @Failure		400 "Email already in use"
// @Failure		404	"User not found."
// @Failure		500	"Internal server error"
// @Router			/user [put]
func PutUser(c echo.Context) error {
	// PARAM
	var id string
	id = c.Param("id")

	// BODY
	var user model.UnsafeUser
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "Error while parsing received data")
	}
	// DATABASE REQUEST GOES HERE

	// SIMULATED NOT FOUND
	if id == "404" {
		return c.JSON(http.StatusNotFound, "User not found")
	}

	// SIMULATED DUPLICATE
	if user.Email == "alreadyIn@use.com" {
		return c.JSON(http.StatusBadRequest, "Email already in use")
	}

	// BUILD RESPONSE
	response := model.PutUserResponse{
		Data: model.SafeUser{
			ID:    id,
			Email: user.Email,
		},
	}

	return c.JSON(http.StatusOK, response)
}

// @Summary		Delete user
// @Description	Receives user ID, returns deleted ID
// @Tags			user
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200		{object}	model.DeleteUserResponse
// @Failure		404	"User not found."
// @Router			/user [delete]
func DeleteUser(c echo.Context) error {
	// PARAM
	var id string
	id = c.Param("id")

	// SIMULATED NOT FOUND
	if id == "404" {
		return c.JSON(http.StatusNotFound, "User doesn't exist")
	}

	// DATABASE REQUEST GOES HERE

	//BUILD RESPONSE
	response := model.DeleteUserResponse{
		ID: id,
	}

	return c.JSON(http.StatusOK, response)
}
