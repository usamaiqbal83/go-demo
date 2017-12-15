package apirouterv1

import "github.com/labstack/echo"

func RouteUserRoleApis(acc *echo.Group, res *echo.Group, session *mgo.Session) {

	// initialize service of user role
	userRoleService := service.UserRoleService{Service: service.Service{Session: session}}
	// initialize service of user
	userService := service.UserService{Service: service.Service{Session: session}}

	// initialize user role controller
	urc := controllerv1.UserRoleController{UserRoleService: &userRoleService, UserService: &userService}

	// Get a assignable user roles
	acc.GET("/user/:id/userrole", urc.GetAssignableUserRole)
}
