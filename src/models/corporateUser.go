package models

import "time"

type CorporateUser struct {
	Id               string    `json:"Id" bson:"Id"`
	CorporateId      string    `json:"CorporateId" bson:"CorporateId"`
	PhoneNumber      string    `json:"PhoneNumber" bson:"PhoneNumber"`
	Gender           string    `json:"Gender" bson:"Gender"`
	UserSubType      string    `json:"UserSubType" bson:"UserSubType"`
	FirstName        string    `json:"FirstName" bson:"FirstName"`
	LastName         string    `json:"LastName" bson:"LastName"`
	Rating           string    `json:"Rating" bson:"Rating"`
	Owner            string    `json:"Owner" bson:"Owner"`
	Email            string    `json:"Email" bson:"Email"`
	EmailVerified    string    `json:"EmailVerified" bson:"EmailVerified"`
	ProfilePicUrl    string    `json:"ProfilePicUrl" bson:"ProfilePicUrl"`
	TacVersion       string    `json:"TacVersion" bson:"TacVersion"`
	TacLanguage      string    `json:"TacLanguage" bson:"TacLanguage"`
	EmergencyContact string    `json:"EmergencyContact" bson:"EmergencyContact"`
	Status           string    `json:"Status" bson:"Status"`
	CreatedTime      time.Time `json:"CreatedTime" bson:"CreatedTime"`
	UpdatedTime      time.Time `json:"UpdatedTime" bson:"UpdatedTime"`
}
