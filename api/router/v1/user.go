package apirouterv1

import "github.com/labstack/echo"

// imports have been intentionally removed

func RouteUserApis(acc *echo.Group, res *echo.Group, session *mgo.Session) {

	// initialize service of user role
	userRoleService := service.UserRoleService{Service: service.Service{Session: session}}

	// initialize service of user
	userService := service.UserService{Service: service.Service{Session: session}}

	// initialize user controller
	urc := controllerv1.UserController{UserRoleService: &userRoleService, UserService: &userService}

	// Create a new user
	acc.POST("/user", urc.CreateUser)

	// GMail login
	acc.POST("/gmailSignIn", urc.SignInWithGMailAccount)

	// get user with domain
	acc.GET("/user", urc.GetOwnerOfDomain)

	// login customer which falls under a customer
	acc.POST("/user/:id/customerLogin", urc.CustomerLoginInfo)

	// login customer with his telephonic credentials
	acc.POST("/user/teleLogin", urc.TelephonicLogin)

	// get user info
	res.GET("/user/:id", urc.GetUserInfo)

	// update user info
	res.PUT("/user/:id", urc.UpdateUserInfo)

	// update user password
	res.PATCH("/user/:id/updatePassword", urc.UpdateUserPassword)
}
