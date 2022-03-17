package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Routes{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Riquisitions:          controllers.CreateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Riquisitions:          controllers.SearchUsers,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodGet,
		Riquisitions:          controllers.SearchUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Riquisitions:          controllers.UpdateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Riquisitions:          controllers.DeleteUser,
		RequestAuthentication: false,
	},
}
