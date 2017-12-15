package auth

// imports have been intentionally removed

// jwt token claims which contains info regarding user
type JwtUserClaim struct {
	ID        bson.ObjectId `json:"id" bson:"_id"`
	FirstName string        `json:"firstName" bson:"firstName"`
	LastName  string        `json:"lastName" bson:"lastName"`
	Email     string        `json:"email" bson:"email"`
	jwt.StandardClaims
}

func Token(user model.User) (string, error) {
	// Set custom claims
	claims := &JwtUserClaim{
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * TokenExpiryTime).Unix(),
		},
	}

	fmt.Println(claims)

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(JwtSecret))
	if err != nil {
		return "", err
	}

	return t, nil
}
