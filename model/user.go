package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	// User represents the structure of our resource
	User struct {
		ID             bson.ObjectId `json:"id" bson:"_id"`
		TelephonicId   int32         `json:"telephonicId" bson:"telephonicId"`
		TelephonicCode int32         `json:"telephonicCode" bson:"telephonicCode"`
		GMailAccountId string        `json:"gmailId" bson:"gmailId"`
		FirstName      string        `json:"firstName" bson:"firstName" validate:"required"`
		LastName       string        `json:"lastName" bson:"lastName" validate:"required"`
		CompanyName    string        `json:"companyName" bson:"companyName" validate:"required"`
		Email          string        `json:"email" bson:"email" validate:"required,email"`
		Phone          string        `json:"phone" bson:"phone" validate:"required"`
		Password       string        `json:"password" bson:"password" validate:"required,min=8"`
		Role           UserRole      `json:"role,omitempty" bson:"role,omitempty" validate:"required"`
		ParentID       bson.ObjectId `json:"parentid,omitempty" bson:"parentid,omitempty"`
		Account        Account       `json:"account,omitempty" bson:"account,omitempty"`
		Info           Info          `json:"info" bson:"info"`
		CreateDate     time.Time     `json:"createDate" bson:"createDate"`
		UpdateDate     time.Time     `json:"updateDate" bson:"updateDate"`
	}

	Account struct {
		SiteName    string `json:"siteName,omitempty" bson:"siteName,omitempty"`
		Phone       string `json:"phone,omitempty" bson:"phone,omitempty"`
		SalesPhone  string `json:"salesPhone,omitempty" bson:"salesPhone,omitempty"`
		SalesEmail  string `json:"salesEmail,omitempty" bson:"salesEmail,omitempty"`
		HomePageURL string `json:"homePageUrl,omitempty" bson:"homePageUrl,omitempty"`
	}

	Info struct {
		TotalNumbersUploaded int64 `json:"totalNumbersUploaded" bson:"totalNumbersUploaded"`
	}
)

type Users []User
