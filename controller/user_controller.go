package controller

import (
	"API-ECHO/model"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

// DUMMY DATA
var DummyUser = model.SafeUser{
	ID:    "someID",
	Email: "jon@doe.com",
}

// @Summary		Create new user
// @Description	Receives user email and password, returns UUID
// @Tags			user
// @Accept			json
// @Produce		json
// @Param			email	path		string	true	"User email"
// @Success		201		{object}	model.PostUserResponse
// @Failure		400 "Bad Request"
// @Failure		422 "Email already in use"
// @Failure		500	"Internal server error"
// @Router			/user [post]
func PostUser(context echo.Context) error {
	// BODY
	var user model.UnsafeUser
	err := context.Bind(&user)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "Error while parsing received data")
	}
	// Empty data
	if user.Email == "" || user.Password == "" {
		return context.JSON(http.StatusBadRequest, "Error while parsing received data")
	}

	// DATABASE REQUEST GOES HERE

	// CHECK DUPLICATE
	if user.Email == "alreadyIn@use.com" {
		return context.JSON(http.StatusUnprocessableEntity, "Email already in use")
	}

	// BUILD RESPONSE
	response := model.PostUserResponse{
		ID: "somerandomid",
	}
	return context.JSON(http.StatusCreated, response)
}

// @Summary		Get user data
// @Description	Receives ID by request param and retreives user data
// @Tags			user
// @Produce		json
// @Param			id	path		int	true	"User ID"
// @Success		200	{object}	model.GetUserResponse
// @Failure		404	"User not found."
// @Router			/user/{id} [get]
func GetUser(context echo.Context) error {
	// PARAM
	var id string
	id = context.Param("id")

	/**
		DATABASE REQUEST GOES HERE
	**/

	if id == "404" {
		return context.JSON(http.StatusNotFound, "User not found.")
	}

	//BUILD RESPONSE
	response := model.GetUserResponse{
		Data: DummyUser,
	}

	return context.JSON(http.StatusOK, response)
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
func GetUserList(context echo.Context) error {
	// QUERY
	email := context.QueryParam("email")
	// PAGING
	page, err := strconv.Atoi(context.QueryParam("page"))
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
	var totalPages int
	var responseList []model.SafeUser
	// IF FILTERS
	var filteredList []model.SafeUser
	if email != "" {
		for i := range userList {
			if strings.Contains(userList[i].Email, email) {
				filteredList = append(filteredList, userList[i])
			}
		}
		if END > len(filteredList) {
			END = len(filteredList)
		}
		responseList = filteredList[START:END]
		totalPages = int(math.Ceil(float64(len(filteredList) / 2)))
	} else {
		if END > len(userList) {
			END = len(userList)
		}
		responseList = userList[START:END]
		totalPages = int(math.Ceil(float64(len(userList) / 2)))
	}

	//BUILD RESPONSE
	response := model.GetUserListResponse{
		Data: responseList,
		Paging: model.Paging{
			Page:  page,
			Total: totalPages,
		},
	}

	return context.JSON(http.StatusOK, response)
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
func PutUser(context echo.Context) error {
	// PARAM
	var id string
	id = context.Param("id")

	// BODY
	var user model.UnsafeUser
	err := context.Bind(&user)
	if err != nil {
		return context.JSON(http.StatusBadRequest, "Error while parsing received data")
	}
	// Empty data
	if user.Email == "" || user.Password == "" {
		return context.JSON(http.StatusBadRequest, "Error while parsing received data")
	}

	// DATABASE REQUEST GOES HERE

	// SIMULATED NOT FOUND
	if id == "404" {
		return context.JSON(http.StatusNotFound, "User not found")
	}

	// SIMULATED DUPLICATE
	if user.Email == "alreadyIn@use.com" {
		return context.JSON(http.StatusUnprocessableEntity, "Email already in use")
	}

	// BUILD RESPONSE
	response := model.PutUserResponse{
		Data: model.SafeUser{
			ID:    id,
			Email: user.Email,
		},
	}

	return context.JSON(http.StatusOK, response)
}

// @Summary		Delete user
// @Description	Receives user ID, returns deleted ID
// @Tags			user
// @Produce		json
// @Param			id	path		string	true	"User ID"
// @Success		200		{object}	model.DeleteUserResponse
// @Failure		404	"User not found."
// @Router			/user [delete]
func DeleteUser(context echo.Context) error {
	// PARAM
	var id string
	id = context.Param("id")

	// SIMULATED NOT FOUND
	if id == "404" {
		return context.JSON(http.StatusNotFound, "User doesn't exist")
	}

	// DATABASE REQUEST GOES HERE

	//BUILD RESPONSE
	response := model.DeleteUserResponse{
		ID: id,
	}

	return context.JSON(http.StatusOK, response)
}
