package pkg

import (
	"github.com/xuchaoi/alertmanager-webhook-sms/cmd/sms-sender/app/option"
	"net/http"
)

type WebhookServer struct {
	Server       *http.Server
	SMSSenderCfg option.SMSConfiguration
}

func (ws *WebhookServer) Handle(w http.ResponseWriter, r *http.Request)  {

}