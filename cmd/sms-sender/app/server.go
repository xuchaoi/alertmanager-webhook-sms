package app

import (
	"context"
	"fmt"
	"github.com/xuchaoi/alertmanager-webhook-sms/cmd/sms-sender/app/option"
	"github.com/xuchaoi/alertmanager-webhook-sms/pkg/webhook"
	"k8s.io/klog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Run(o *option.SMSSenderOptions) error {
	ws := &webhook.Server{
		HttpServer: &http.Server{
			Addr: fmt.Sprintf("0.0.0.0:%v", o.SenderPort),
		},
		SMSSenderCfg: o.SMSCfg,
		MysqlCfg: o.MysqlCfg,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", HealthHandler)
	mux.HandleFunc("/sms", ws.Handle)
	mux.HandleFunc("/wechat", ws.Handle)
	ws.HttpServer.Handler = mux

	go func() {
		if err := ws.HttpServer.ListenAndServe(); err != nil {
			klog.Errorf("Failed to listen and handle SMS-Sender server: %v", err)
		}
	}()

	klog.Info("SMS Sender started.")

	// listening OS shutdown singal
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	klog.Infof("Got OS shutdown signal, shutting down webhook server gracefully...")
	ws.HttpServer.Shutdown(context.Background())

	panic("unreachable")
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "ok")
}