package datasource

// imports have been intentionally removed

type UserDataSource struct {
	DataSource
}

// this method returns user database object using object id
func (uds *UserDataSource) UserObject(objectID bson.ObjectId) (*model.User, error) {
	parent := model.User{}

	err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).FindId(objectID).One(&parent)

	if err != nil {
		return nil, err
	}
	return &parent, err
}

// this method returns user database object using object id
func (uds *UserDataSource) UserObjectWithCredentials(username string, password string) (*model.User, error) {
	user := model.User{}

	if err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).
		Find(bson.M{"$and": []bson.M{
			bson.M{"email": username},
			bson.M{"password": cmlutils.GetMD5Hash(password)}}}).One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// this method returns user database object using object id
func (uds *UserDataSource) UserObjectWithGMailAccount(username string, gmailId string) (*model.User, error) {
	user := model.User{}

	if err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).
		Find(bson.M{"$and": []bson.M{
			bson.M{"email": username},
			bson.M{"gmailId": gmailId}}}).One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// this method returns user database object using object id
func (uds *UserDataSource) UserObjectWithTelephonicCredentials(telephonicId int32, telephonicCode int32) (*model.User, error) {
	user := model.User{}

	if err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).
		Find(bson.M{"$and": []bson.M{
			bson.M{"telephonicId": telephonicId},
			bson.M{"telephonicCode": telephonicCode}}}).One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// this method returns super user if present
func (uds *UserDataSource) SuperUser() (*model.User, error) {
	superUser := model.User{}

	if err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).Find(bson.M{
		"role.name": bson.M{"$eq": cmlconstants.RoleSuper},
	}).One(&superUser); err != nil {
		return nil, err
	}
	return &superUser, nil
}

func (uds *UserDataSource) SaveUser(user *model.User) error {
	// if there is no user id assign one
	if user.ID == "" {
		user.ID = bson.NewObjectId()
	}

	existingUserObject, err := uds.UserObject(user.ID)
	if err != nil {
		// user doesn't exist, create new
		user.CreateDate = time.Now().UTC()
		user.UpdateDate = time.Now().UTC()
		// Write the user to mongo
		if err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).Insert(user); err != nil {
			return err
		}
	} else {
		// user exists
		user.UpdateDate = time.Now().UTC()
		if err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).UpdateId(existingUserObject.ID, user); err != nil {
			return err
		}
	}
	return nil
}

// this method returns the user using his email address
func (uds *UserDataSource) UserWithEmailAddress(emailAddress string) (*model.User, error) {

	user := model.User{}

	if err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).Find(bson.M{"email": emailAddress}).One(&user); err != nil {
		return nil, err
	}

	return &user, nil
}

// this method returns the owner of the domain
func (uds *UserDataSource) UserWithDomain(domain string) (*model.User, error) {
	user := model.User{}

	err := uds.DbSession().DB(cmlconstants.Database).C(cmlconstants.CollectionUsers).Find(bson.M{"account.homePageUrl": domain}).One(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
