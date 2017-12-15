package apirouterv1

import "github.com/labstack/echo"

// imports have been intentionally removed

func RouteCampaignApis(acc *echo.Group, res *echo.Group, session *mgo.Session) {
	// initialize service of campaign
	campaignService := service.CampaignService{Service: service.Service{Session: session}}

	// initialize service of user
	userService := service.UserService{Service: service.Service{Session: session}}

	// initialize s3 service
	s3Service := service.S3Service{}

	// initialize service of contact list
	contactListService := service.ContactListService{Service: service.Service{Session: session}}

	// initialize service of contact group
	contactGroupService := service.ContactGroupService{Service: service.Service{Session: session}}

	// initialize campaign controller
	campaignController := controllerv1.CampaignController{CampaignService: &campaignService,
		UserService:         &userService,
		S3Service:           &s3Service,
		ContactListService:  &contactListService,
		ContactGroupService: &contactGroupService,
	}

	// create campaign
	res.POST("/user/:id/campaign", campaignController.CreateCampaign)

	// get campaign list
	res.GET("/user/:id/campaign", campaignController.GetCampaignListForUser)

	// get resources home address
	res.GET("/user/:id/campaign/:cid", campaignController.CampaignAction)

	// test campaign by sending test call on given number
	res.POST("/user/:id/campaign/test", campaignController.TestCampaign)
}
