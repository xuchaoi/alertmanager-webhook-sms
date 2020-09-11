package webhook

import (
	"encoding/json"
	"github.com/xuchaoi/alertmanager-webhook-sms/cmd/sms-sender/app/option"
	"github.com/xuchaoi/alertmanager-webhook-sms/pkg/webhook/handle"
	"io/ioutil"
	"k8s.io/klog"
	"net/http"
)

type Server struct {
	HttpServer   *http.Server
	SMSSenderCfg option.SMSConfiguration
	MysqlCfg     option.MysqlConfiguration
}

func (ws *Server) Handle(w http.ResponseWriter, r *http.Request)  {
	var body []byte
	if r.Body != nil {
		if data, err := ioutil.ReadAll(r.Body); err == nil {
			body = data
		}
	}
	if len(body) == 0 {
		klog.Error("alertManager data is empty")
		http.Error(w, "alertManager data is empty", http.StatusBadRequest)
	}

	if r.URL.Path == "sms" || r.URL.Path == "wechat" {
		var f interface{}
		unmarshalErr := json.Unmarshal(body, &f)
		if unmarshalErr != nil {
			klog.Error(unmarshalErr)
		}
		alertData := f.(map[string]interface{})
		alerts := alertData["alerts"].([]interface{})
		for _, alerti := range alerts {
			alert := alerti.(map[string]interface{})
			annotations := alert["annotations"].(map[string]interface{})
			smsContent := annotations["description"].(string)

			if r.URL.Path == "sms" {
				klog.Infof("Start to send sms!")
				info, err := handle.SMSHandle(smsContent, ws.SMSSenderCfg)
				if err != nil {
					klog.Errorf("Failed send SMS by SMS interface, err: %v", err)
				} else {
					klog.Infof("End to send sms, result: %s", info)
				}
			} else if r.URL.Path == "wechat" {
				klog.Infof("Start to send alert to wechat database!")
				info, err := handle.WechatHandle(smsContent, ws.MysqlCfg)
				if err != nil {
					klog.Errorf("Failed send SMS to mysql, err: %v", err)
				} else {
					klog.Infof("End to send sms to mysql, result: %s", info)
				}
			}
		}
	} else {
		klog.Infof("This url does not belong to the handle processing range, url: %s", r.URL.Path)
		return
	}

}