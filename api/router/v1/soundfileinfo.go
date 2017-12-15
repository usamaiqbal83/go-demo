package apirouterv1

import "github.com/labstack/echo"

// imports have been intentionally removed

func RouteSoundFileInfoApis(acc *echo.Group, res *echo.Group, session *mgo.Session) {

	// initialize service of sound file
	soundFileService := service.SoundFileInfoService{Service: service.Service{Session: session}}

	// initialize service of user
	userService := service.UserService{Service: service.Service{Session: session}}

	// initialize s3 service
	s3Service := service.S3Service{}

	// initialize user controller
	soundFileController := controllerv1.SoundFileInfoController{SoundFileInfoService: &soundFileService,
		UserService: &userService,
		S3Service:   &s3Service}

	// upload sound file
	res.POST("/user/:id/soundfile", soundFileController.UploadSoundFile)

	// update sound file name for resource
	res.PATCH("/user/:id/soundfile/:sid", soundFileController.SoundFileNameUpdate)

	// delete sound file
	res.DELETE("/user/:id/soundfile/:sid", soundFileController.SoundFileDelete)

	// get sound file list
	res.GET("/user/:id/soundfile", soundFileController.SoundFilesGetList)
}
