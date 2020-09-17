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
		} else {
			klog.Errorf("Read body error, err: %v", err)
			return
		}
	}
	if len(body) == 0 {
		klog.Error("AlertManager data is empty")
		return
	}

	klog.V(4).Infof("SMSSender Cfg: %v", ws.SMSSenderCfg)
	klog.V(4).Infof("Mysql cfg: %v", ws.MysqlCfg)

	if r.URL.Path == "/sms" || r.URL.Path == "/wechat" {
		var f interface{}
		unmarshalErr := json.Unmarshal(body, &f)
		if unmarshalErr != nil {
			klog.Error(unmarshalErr)
		}
		alertData := f.(map[string]interface{})
		if alertData["alerts"] == nil {
			klog.Error("There is no alerts object in the alarm data")
			return
		}
		alerts := alertData["alerts"].([]interface{})
		for _, alerti := range alerts {
			alert := alerti.(map[string]interface{})
			if alert["annotations"] == nil {
				klog.Error("There is no alerts annotations object in the alarm data")
				return
			}
			annotations := alert["annotations"].(map[string]interface{})
			if annotations["description"] == nil || annotations["description"] == "" {
				klog.Error("There is no alerts annotations desc in the alarm data")
				return
			}
			smsContent := annotations["description"].(string)

			if r.URL.Path == "/sms" {
				klog.V(2).Infof("Start to send sms, sms content: %s", smsContent)
				info, err := handle.SMSHandle(smsContent, ws.SMSSenderCfg)
				if err != nil {
					klog.Errorf("Failed send SMS by SMS interface, err: %v", err)
				} else {
					klog.V(2).Infof("End to send sms, result: %s", info)
				}
			} else if r.URL.Path == "/wechat" {
				klog.V(2).Infof("Start to send alert to wechat database, sms content: %s", smsContent)
				info, err := handle.WechatHandle(smsContent, ws.MysqlCfg)
				if err != nil {
					klog.Errorf("Failed send SMS to mysql, err: %v", err)
				} else {
					klog.V(2).Infof("End to send sms to mysql, result: %s", info)
				}
			}
		}
	} else {
		klog.V(2).Infof("This url does not belong to the handle processing range, url: %s", r.URL.Path)
	}
	return
}