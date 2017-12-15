package apirouterv1

import "github.com/labstack/echo"

// imports have been intentionally removed

func RouteRootApis(acc *echo.Group, res *echo.Group, session *mgo.Session) {

	// instantiate new root controller
	rootController := controllerv1.RootController{}

	// get resources home address
	acc.GET("/assethome", rootController.GetAssetHomePath)
}
