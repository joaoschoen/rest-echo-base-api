# Base Echo REST API

The objective of this project is to develop a fully functional API using all the functionalities required from Echo to comply to the best practices of REST API

The secondary objective is that this API is a boilerplate for any database implementation, so it cannot connect to any database, it is only concerned with the treatment of http requests, thus it knows how to use all 

# How to run

To run this API simply use:

```
go run .
```

Or alternatively for active development, install [Air](https://github.com/cosmtrek/air) and then use:

```
air
```

## Libs used

- [Echo](https://github.com/labstack/echo)  
    - Backend framework
- [Godotenv](https://github.com/joho/godotenv)
    - Environment variables loading

## Environment

This api uses a .env file for configuration, at the current moment here are the features that can be configured

- PORT

## Features

- MVC pattern for project structure
- Routing with multiple files and folders
- JSON format for data interchange

## Methods

### GET

- :param based endpoint to GET single objects

### PUT

- :param based endpoint to UPDATE an object with given param

### POST

- JSON body request treatment and response

### DELETE

- :param based endpoint to DELETE an object with given param