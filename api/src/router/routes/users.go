package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		Uri:      "/users",
		Method:   http.MethodGet,
		Function: controllers.GetAllUsers,
		NeedAuth: false,
	},

	{
		Uri:      "/user/{userId}",
		Method:   http.MethodGet,
		Function: controllers.GetUser,
		NeedAuth: false,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodPut,
		Function: controllers.UpdateUser,
		NeedAuth: false,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodPut,
		Function: controllers.CreateUser,
		NeedAuth: false,
	},
	{
		Uri:      "/user/{userId}",
		Method:   http.MethodDelete,
		Function: controllers.DeleteUser,
		NeedAuth: false,
	},
}
