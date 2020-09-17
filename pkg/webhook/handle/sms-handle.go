package handle

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xuchaoi/alertmanager-webhook-sms/cmd/sms-sender/app/option"
	"github.com/xuchaoi/alertmanager-webhook-sms/pkg/util"
	"io/ioutil"
	"k8s.io/klog"
)

type SMSParams struct {
	Params SMSParam `json:"params"`
}

type SMSParam struct {
	SmsCode       string `json:"smsCode"`
	SmsContent    string `json:"smsContent"`
	NewStaffId    string `json:"newStaffId"`
	EffectiveDate string `json:"effectiveDate"`
	SubPort       string `json:"subPort"`
	CrmpfPubInfo  option.CrmpfPubInfo `json:"crmpfPubInfo"`
}

func SMSHandle(smsContent string, smsCfg option.SMSConfiguration) (string, error) {
	smsParams := SMSParams{
		Params: SMSParam{
			SmsCode: smsCfg.Code,
			SmsContent: smsContent,
			NewStaffId: smsCfg.NewStaffId,
			EffectiveDate: smsCfg.EffectiveDate,
			SubPort: smsCfg.SubPort,
			CrmpfPubInfo: smsCfg.CrmpfPubInfo,
		},
	}

	buf, err := json.Marshal(smsParams)
	if err != nil {
		e := fmt.Sprintf("Failed to convert the smsParams to byte through json tool, err: %v", err)
		return "", errors.New(e)
	}
	klog.V(4).Infof("SMS sender request body: %s", string(buf))
	res, err := util.SMSHttpPost(smsCfg.Url, bytes.NewReader(buf))
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
	klog.V(4).Infof("SMS interface resCode: %d", res.StatusCode)
	klog.V(4).Infof("SMS interface resData: %v", resData)
	if resData["object"] == nil {
		e := fmt.Sprint("Failed to get SMS body.object, err: body.object is nil")
		return "", errors.New(e)
	}
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
