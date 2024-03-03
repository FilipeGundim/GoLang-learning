package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Func:                  controllers.CreateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Func:                  controllers.SearchUsers,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodDelete,
		Func:                  controllers.DeleteUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPatch,
		Func:                  controllers.UpdateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{userId}",
		Method:                http.MethodPut,
		Func:                  controllers.UpdateUser,
		RequestAuthentication: false,
	},
}
