package main

// imports have been intentionally removed

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {

	// initialize logger
	logger.InitLogger()

	// general configurations
	configureGeneral()

	// configure new db session
	dbSession := getSession()
	defer dbSession.Close()

	// configure to serve WebServices
	configureHttpServer(dbSession)
}

func configureGeneral() {
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println(cmlutils.RandomString(10))
}

func configureHttpServer(dbSession *mgo.Session) {

	e := echo.New()

	// add validation
	e.Validator = &CustomValidator{validator: validator.New()}

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.BodyDump(bodyDumpHandler))

	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		ExposeHeaders: []string{"x_auth_token"},
	}))

	// perform routing for v1 version of web apis
	performV1APIRouting(e, dbSession)

	// Server
	e.Logger.Fatal(e.Start(":3000"))
}

func performV1APIRouting(echo *echo.Echo, dbSession *mgo.Session) {

	// accessible web services will fall in this group
	acc := echo.Group("/v1")

	// restricted web services will fall in this group
	res := echo.Group("/v1")

	// Configure middleware with the custom claims type
	config := middleware.JWTConfig{
		Claims:     &auth.JwtUserClaim{},
		SigningKey: []byte(auth.JwtSecret),
	}

	// add middleware for restricted services
	res.Use(middleware.JWTWithConfig(config))
	res.Use(auth.MiddlewareRes)

	// route user role apis
	apirouterv1.RouteUserRoleApis(acc, res, dbSession)

	// route user apis
	apirouterv1.RouteUserApis(acc, res, dbSession)

	// route root apis
	apirouterv1.RouteRootApis(acc, res, dbSession)

	// route sound file apis
	apirouterv1.RouteSoundFileInfoApis(acc, res, dbSession)

	// route contact list apis
	apirouterv1.RouteContactListApis(acc, res, dbSession)

	// route campaign apis
	apirouterv1.RouteCampaignApis(acc, res, dbSession)
}

// getSession creates a new mongo session and panics if connection error occurs
func getSession() *mgo.Session {
	// for debugging only
	// credentials will be taken from environment variables in production

	info := &mgo.DialInfo{
		Addrs:    []string{cmlconstants.Host1, cmlconstants.Host2},
		Timeout:  60 * time.Second,
		Database: cmlconstants.Database,
		Username: cmlconstants.Username,
		Password: cmlconstants.Password,
		Source:   cmlconstants.Authenticationdb,
	}

	session, err := mgo.DialWithInfo(info)

	if err != nil {
		panic(err)
	} else {
		fmt.Println("----------------------------------- \n")
		fmt.Println("*** Database Session created ***\n")
		fmt.Println("----------------------------------- \n")
	}

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	return session
}

// middle ware handler
func bodyDumpHandler(c echo.Context, reqBody, resBody []byte) {

	logger.PrintString("================================")

	logger.PrintRequest(c.Request())

	logger.PrintString("--------request body-------")
	logger.PrintBody(reqBody)
	logger.PrintString("---------------------------")

	logger.PrintString("-------- response body --------")
	logger.PrintBody(resBody)
	logger.PrintString("-------------------------------")
	logger.PrintString("=================================")
}
