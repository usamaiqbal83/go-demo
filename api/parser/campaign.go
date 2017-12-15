package parser

// imports have been intentionally removed

// return campaign object after parsing the input request
func ParseCampaignInfoFromContext(c echo.Context, userId string, campaignService *service.CampaignService) (*model.Campaign, error) {

	campaignType := c.FormValue("type")
	campaignName := c.FormValue("name")
	callerId := c.FormValue("callerId")
	soundFileId := c.FormValue("soundFileId")
	dncSoundFileId := c.FormValue("dncSoundFileId")
	dncKey := c.FormValue("dncKey")
	vmSoundFileId := c.FormValue("vmSoundFileId")
	transferSoundFileId := c.FormValue("transferSoundFileId")
	transferNumber := c.FormValue("transferNumber")
	transferKey := c.FormValue("transferKey")
	scrubNumbersFromDNC := c.FormValue("scrubNumbersFromDNC")
	addToDNCList := c.FormValue("addToDNCList")
	removeDuplicate := c.FormValue("removeDuplicate")

	campaignTypeInteger, err := strconv.Atoi(campaignType)
	if err != nil {
		return nil, errors.New(cmlmessages.CampaignRequestFormatIncorrect)
	}

	campaignObject := model.Campaign{}
	campaignObject.Type = int8(campaignTypeInteger)
	campaignObject.Name = campaignName
	// validate input request body
	if err := c.Validate(campaignObject); err != nil {
		return nil, err
	}

	// check input data as per contact id
	switch campaignTypeInteger {
	case cmlconstants.CampaignTypeVoiceOnly:
		campaignInfo, err := campaignService.CreateVoiceOnlyCampaignInfo(userId, callerId, soundFileId)
		if err != nil {
			return nil, err
		}

		campaignObject.CampaignInfo = *campaignInfo
	case cmlconstants.CampaignTypeLiveAnswerAndAnswerMachineNoTransfer:
		campaignInfo, err1 := campaignService.CreateLiveAnswerAndAnswerMachineNoTransferCampaignInfo(userId, callerId,
			soundFileId,
			dncSoundFileId, dncKey,
			vmSoundFileId)
		fmt.Print(err1)
		if err1 != nil {
			return nil, err1
		}

		campaignObject.CampaignInfo = *campaignInfo
	case cmlconstants.CampaignTypeLiveAnswerAndAnswerMachineWithTransfer:
		campaignInfo, err2 := campaignService.CreateLiveAnswerAndAnswerMachineWithTransferCampaignInfo(userId, callerId,
			soundFileId,
			dncSoundFileId, dncKey,
			transferSoundFileId, transferNumber, transferKey,
			vmSoundFileId)
		if err2 != nil {
			return nil, err2
		}

		campaignObject.CampaignInfo = *campaignInfo
		campaignObject.Limit.MaxLimit = cmlconstants.CampaignDefaultMaxTransferLimit
	case cmlconstants.CampaignTypeLiveAnswerOnlyNoTransfer:
		campaignInfo, err3 := campaignService.CreateLiveAnswerOnlyNoTransferCampaignInfo(userId, callerId,
			soundFileId,
			dncSoundFileId, dncKey)
		if err3 != nil {
			return nil, err3
		}

		campaignObject.CampaignInfo = *campaignInfo
	case cmlconstants.CampaignTypeLiveAnswerOnlyWithTransfer:
		campaignInfo, err4 := campaignService.CreateLiveAnswerOnlyWithTransferCampaignInfo(userId, callerId,
			soundFileId,
			dncSoundFileId, dncKey,
			transferSoundFileId, transferNumber, transferKey)
		if err4 != nil {
			return nil, err4
		}

		campaignObject.CampaignInfo = *campaignInfo
		campaignObject.Limit.MaxLimit = cmlconstants.CampaignDefaultMaxTransferLimit
	default:
		return nil, errors.New("Campaign type is not supported")
	}

	campaignObject.CampaignInfo.ScrubNumbersFromDNC = scrubNumbersFromDNC == "true" || scrubNumbersFromDNC == "True"
	campaignObject.CampaignInfo.AddToDNCList = addToDNCList == "true" || addToDNCList == "True"
	campaignObject.CampaignInfo.RemoveDuplicate = removeDuplicate == "true" || removeDuplicate == "True"
	campaignObject.Limit.ConcurrentCalls = cmlconstants.CampaignDefaultCCLimit
	campaignObject.Status = cmlconstants.CampaignStatusNew
	campaignObject.CampaignCode = cmlutils.EpochMilli()
	campaignObject.Error.Code = cmlcodes.NoError
	campaignObject.Error.Description = ""

	return &campaignObject, nil
}
