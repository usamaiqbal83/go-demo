package model

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type (
	// Campaign represents the structure of our resource
	Campaign struct {
		ID              bson.ObjectId   `json:"id" bson:"_id"`
		CampaignCode    int64           `json:"campaignCode" bson:"campaignCode"`
		Type            int8            `json:"type" bson:"type" validate:"required,gte=1,lte=5"`
		Status          int8            `json:"status" bson:"status"`
		Name            string          `json:"name" bson:"name" validate:"required"`
		Limit           Limit           `json:"limit,omitempty" bson:"limit,omitempty"`
		ContactListInfo ContactListInfo `json:"contactListInfo" bson:"contactListInfo"`
		UserID          bson.ObjectId   `json:"userid,omitempty" bson:"userid"`
		CampaignInfo    CampaignInfo    `json:"campaignInfo" bson:"campaignInfo" validate:"required"`
		CampaignStats   Stats           `json:"campaignStats" bson:"campaignStats"`
		CreateDate      time.Time       `json:"createDate" bson:"createDate"`
		UpdateDate      time.Time       `json:"updateDate" bson:"updateDate"`
		Error           Error           `json:"error" bson:"error"`
	}

	// structure represents the limits for the campaign
	Limit struct {
		ConcurrentCalls int `json:"concurrentCalls" bson:"concurrentCalls"`
		MaxLimit        int `json:"maxLimit,omitempty" bson:"maxLimit,omitempty"`
	}

	// structure represents the campaign info
	CampaignInfo struct {
		CallerId            string        `json:"callerId" bson:"callerId"`
		SoundFileId         bson.ObjectId `json:"soundFileId" bson:"soundFileId"`
		DNCSoundFileId      bson.ObjectId `json:"dncSoundFileId,omitempty" bson:"dncSoundFileId,omitempty"`
		DNCKey              string        `json:"dncKey,omitempty" bson:"dncKey,omitempty"`
		VMSoundFileId       bson.ObjectId `json:"vmSoundFileId,omitempty" bson:"vmSoundFileId,omitempty"`
		TransferSoundFileId bson.ObjectId `json:"transferSoundFileId,omitempty" bson:"transferSoundFileId,omitempty"`
		TransferNumber      string        `json:"transferNumber,omitempty" bson:"transferNumber,omitempty"`
		TransferKey         string        `json:"transferKey,omitempty" bson:"transferKey,omitempty"`
		ScrubNumbersFromDNC bool          `json:"scrubNumbersFromDNC" bson:"scrubNumbersFromDNC"`
		AddToDNCList        bool          `json:"addToDNCList" bson:"addToDNCList"`
		RemoveDuplicate     bool          `json:"removeDuplicate" bson:"removeDuplicate"`
	}

	// structure represents the campaign info
	ContactListInfo struct {
		ID   bson.ObjectId `json:"id" bson:"id"`
		Code int64         `json:"code" bson:"code"`
	}

	// structure represents the campaign stats
	Stats struct {
		Total     int64 `json:"total" bson:"total"`
		Dialed    int64 `json:"dialed" bson:"dialed"`
		Busy      int64 `json:"busy" bson:"busy"`
		Error     int64 `json:"error" bson:"error"`
		NoAnswer  int64 `json:"noAnswer" bson:"noAnswer"`
		Live      int64 `json:"live" bson:"live"`
		VoiceMail int64 `json:"voiceMail" bson:"voiceMail"`
		DNC       int64 `json:"dnc" bson:"dnc"`
		Transfer  int64 `json:"transfer" bson:"transfer"`
	}
)

type CampaingList []Campaign
