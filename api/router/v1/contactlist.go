package apirouterv1

import "github.com/labstack/echo"

// imports have been intentionally removed

func RouteContactListApis(acc *echo.Group, res *echo.Group, session *mgo.Session) {

	// initialize service of contact list
	contactListService := service.ContactListService{Service: service.Service{Session: session}}

	// initialize service of contact group
	contactGroupService := service.ContactGroupService{Service: service.Service{Session: session}}

	// initialize service of user
	userService := service.UserService{Service: service.Service{Session: session}}

	// initialize s3 service
	s3Service := service.S3Service{}

	// initialize contact list controller
	contactListController := controllerv1.ContactListController{ContactListService: &contactListService,
		ContactGroupService: &contactGroupService,
		UserService:         &userService,
		S3Service:           &s3Service,
	}

	// csv/xls/xlxs file upload
	res.POST("/user/:id/contactlist", contactListController.UploadContactListForUser)

	// update contact list name for resource
	res.PATCH("/user/:id/contactlist/:cid", contactListController.UpdateContactNameForContactList)

	// get contact list
	res.GET("/user/:id/contactlist", contactListController.GetContactListsForUser)

	// delete contact list
	res.DELETE("/user/:id/contactlist/:cid", contactListController.ContactListDelete)
}
