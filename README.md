# Base Echo REST API

The objective of this project is to develop a fully functional API using all the functionalities required from Echo to comply to the best practices of REST API

The secondary objective is that this API is a boilerplate for any database implementation, so it cannot connect to any database, it is only concerned with the treatment of http requests, thus it knows how to use all 

# **Important:** This is a warning message.

This code is not made with the intention of being safe and should not be used in a serious application as is, reason why all the mentions to passwords use the "BadExample" text, there is no authentication yet, I will be implementing it in the future though.

This is a learning project, by the end of it I want to have a fully functional and safe API that could be used as an example of how to do things in golang without having to search too much for it and also as a repository of my knowledge of how to do things in golang.

## Features

- MVC pattern for project structure
- Routing with multiple files and folders
- JSON format for data interchange
- Documentation using Swagger 2.0
- Unit testing with 100% coverage

## Libs used

- [Echo](https://github.com/labstack/echo)  
    - Backend framework
- [Godotenv](https://github.com/joho/godotenv)
    - Environment variables loading
- [Swaggo](https://github.com/swaggo/swag)
    - Documentation
- [Echo Swagger](https://github.com/swaggo/echo-swagger)
    - Serving the Swagger UI
- [Testify](https://github.com/stretchr/testify)
    - Unit testing 
    
# How to run

To run this API simply use:

```
go run .
```

Or alternatively for active development, install [Air](https://github.com/cosmtrek/air) and then use:

```
air
```

# Docs

This API project has documentation using Swagger 2.0 and the lib [Swaggo](https://github.com/swaggo/swag) with comment annotation, the user controller has examples of how to use them for the for main HTTP request verbs that you'll need. 

The latest version of the documentation will always bbe already computed. To access the documentation, run the project then open this link in your browser http://localhost:3000/swagger/index.html 

Do note that 3000 is the standard IP address if you haven't changed it through .env

To update the documentation, install the [swag](https://github.com/swaggo/swag) CLI and run the command: 
```
swag i
```

## Environment

This api uses a .env file for configuration, at the current moment here are the features that can be configured

- PORT
    - Number representing the port on which the API will be served
- DEBUG
    - If se to true, Debug mode enables the generation of a routes file and the serving of Swagger documentation

## Testing

Testing this API was done with mock tests and the [Testify](https://github.com/stretchr/testify) lib for assertions alongside Golang's standard testing library

The tests in this repository were made using the following structure:

- One function per handler function
    - For each test for a specific handler:
        - One function that sets the stage
        - One assert group that runs the handler with the test presets

To run all tests, run the following command: 
```
go test ./... 
```

You can also add the **-cover** flag at the end to see test coverage in all the packages

## Methods

### GET

- :id based endpoint to GET single objects
- list endpoint to GET a list of objects with query filters and paging

### PUT

- :id based endpoint to UPDATE an object with given param

### POST

- JSON body request treatment and response

### DELETE

- :id based endpoint to DELETE an object with given id

# Writing


One particular design choice I made for this project is to make the names of my variables and functions be more verbose and self descriptive then what I usually see in golang projects 

I much prefer to work with code like this: 
```
func DeleteUser(context echo.Context) error {
	var id string
	id = context.Param("id")

	response := model.DeleteUserResponse{
		ID: id,
	}

	return context.JSON(http.StatusOK, response)
}
```
Then this:
```
func DeleteUser(c echo.Context) error {
	var id string
	id = c.Param("id")

	r := model.DeleteUserResponse{
		ID: id,
	}

	return c.JSON(http.StatusOK, r)
}
```
