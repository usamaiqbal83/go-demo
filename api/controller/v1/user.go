package controllerv1

// imports have been intentionally removed

type UserController struct {
	Controller
	UserRoleService *service.UserRoleService
	UserService     *service.UserService
}

// data structure to get credentials from json object
type (
	Credential struct {
		Username string `json:"username" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
)

// CreateUser creates a new user resource
func (uc *UserController) CreateUser(c echo.Context) error {

	// Stub user to be populated from the body
	u := model.User{}

	if err := c.Bind(&u); err != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, cmlmessages.PleaseTryAgain, cmlmessages.UserRequestFormatIncorrect)
	}

	// validate input request body
	if err := c.Validate(u); err != nil {
		fmt.Println(err)
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, cmlmessages.PleaseTryAgain, err.Error())
	}

	// verify if entered user role is valid
	valid, message := uc.UserRoleService.CanRoleBeAssigned(u.Role.ID)

	if !valid { // role with user is being created is not valid or not allowed
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, cmlmessages.PleaseTryAgain, message)
	}

	// skipping the error catching as role validation is already checked in previous step
	// so there would be no error in this case
	roleObject, _ := uc.UserRoleService.RoleObjectWithId(u.Role.ID)
	u.Role = *roleObject

	isEmailAlreadyExist := uc.UserService.IsUserExistsWithEmailAddress(u.Email)
	if isEmailAlreadyExist { // email with which user is being created is already in use
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, cmlmessages.PleaseTryAgain, cmlmessages.UserWithEmailAlreadyExists)
	}
	// assign parent id if request role is not super
	// in case of super there would not be any parent
	if roleObject.Name != cmlconstants.RoleSuper {

		parentUser, err := uc.UserService.UserObject(u.ParentID.Hex())
		if err != nil {
			// parent not found return error
			return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, cmlmessages.PleaseTryAgain, cmlmessages.InvalidParentId)
		}

		u.ParentID = parentUser.ID
	} else {
		u.ParentID = ""
	}

	// create a new user in database
	if err := uc.UserService.CreateNewUser(&u); err != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, cmlmessages.PleaseTryAgain, cmlmessages.UserCreationFailed)
	}

	// truncate password
	u.Password = "******************************"

	// Marshal provided interface into JSON structure
	uj, _ := json.Marshal(u)

	// generate jwt token
	uc.assignJwtToken(c, u)

	return httpresponse.CreateSuccessResponse(&c, http.StatusCreated, "Success", "Success", uj)
}

func (uc *UserController) GetUserInfo(c echo.Context) error {

	// Grab user id for which file is being uploaded
	userID := c.Param("id")

	// get user object
	user, er1 := uc.UserService.UserObject(userID)
	if er1 != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, cmlmessages.PleaseTryAgain, cmlmessages.UserDoesNotExist)
	}

	// check if user is authorized to make this change
	// as contact list can only be uploaded by customer
	if user.Role.Name != cmlconstants.RoleCustomer {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, cmlmessages.PleaseTryAgain, cmlmessages.UnauthorizedForAction)
	}

	user.Password = "***********************"

	// Marshal provided interface into JSON structure
	userData, _ := json.Marshal(user)
	return httpresponse.CreateSuccessResponse(&c, http.StatusCreated, "Success", "Success", userData)
}

// helper
func (uc *UserController) assignJwtToken(c echo.Context, user model.User) error {

	token, err := auth.Token(user)
	if err != nil {
		return err
	}

	c.Response().Header().Set("x_auth_token", token)

	return nil
}
