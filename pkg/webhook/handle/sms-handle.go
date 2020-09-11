package handle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xuchaoi/alertmanager-webhook-sms/cmd/sms-sender/app/option"
	"github.com/xuchaoi/alertmanager-webhook-sms/pkg/util"
	"io/ioutil"
)

type SMSParams struct {
	smsCode       string
	smsContent    string
	newStaffId    string
	effectiveDate string
	subPort       string
	crmpfPubInfo  option.CrmpfPubInfo
}

func SMSHandle(smsContent string, smsCfg option.SMSConfiguration) (string, error) {
	var smsParams = new(SMSParams)
	smsParams.smsCode       = smsCfg.Code
	smsParams.smsContent    = smsContent
	smsParams.newStaffId    = smsCfg.NewStaffId
	smsParams.effectiveDate = smsCfg.EffectiveDate
	smsParams.subPort       = smsCfg.SubPort
	smsParams.crmpfPubInfo  = smsCfg.CrmpfPubInfo

	buf, err := json.Marshal(smsParams)
	if err != nil {
		e := fmt.Sprintf("Failed to convert the smsParams to byte through json tool, err: %v", err)
		return "", errors.New(e)
	}

	res, err := util.HttpPost(smsCfg.Url, bytes.NewReader(buf))
	if err != nil {
		e := fmt.Sprintf("Failed to send SMS by SMS API, err: %v", err)
		return "", errors.New(e)
	}

	if res.Body == nil {
		e := fmt.Sprintf("Failed to get res body which SMS interface returns, smsContent: %s", smsContent)
		return "", errors.New(e)
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		e := fmt.Sprintf("Failed to convert SMS interface to return body, err: %v", err)
		return "", errors.New(e)
	}

	var f interface{}
	err = json.Unmarshal(resBody, &f)
	if err != nil {
		e := fmt.Sprintf("Failed to convert SMS interface to return body, err: %v", err)
		return "", errors.New(e)
	}

	resData := f.(map[string]interface{})
	resDataObj := resData["object"].(map[string]interface{})
	// Success
	if resDataObj["respCode"] == "0" {
		info := fmt.Sprintf("Successed to send SMS interface, response code: 0, desc: %s", resDataObj["respDesc"])
		return info, nil
	} else {
		e := fmt.Sprintf("Successed to send SMS interface, response code: %s, desc: %s", resDataObj["respCode"], resDataObj["respDesc"])
		return "", errors.New(e)
	}
}
