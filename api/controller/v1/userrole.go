package controllerv1

// imports have been intentionally removed

type UserRoleController struct {
	Controller
	UserRoleService *service.UserRoleService
	UserService     *service.UserService
}

func (urc *UserRoleController) GetAssignableUserRole(c echo.Context) error {

	// Grab id
	id := c.Param("id")

	user, err := urc.UserService.UserObject(id)

	if err != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, "User does not exist", "User does not exist")
	}

	userRoles, err := urc.UserRoleService.AssignableUserRolesForParent(user)

	if err != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, "Roles fetch failed", "Roles fetching failed")
	}

	// Marshal provided interface into JSON structure
	data, err2 := json.Marshal(userRoles)
	if err2 != nil {
		return httpresponse.CreateBadResponse(&c, http.StatusBadRequest, "Roles fetch failed", "Roles fetching failed")
	}

	return httpresponse.CreateSuccessResponse(&c, http.StatusOK, "Roles fetched Successfully", "Roles fetched Successfully", data)
}
