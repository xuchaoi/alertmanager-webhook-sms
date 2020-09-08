package main

import (
	"github.com/spf13/pflag"
	"github.com/xuchaoi/alertmanager-webhook-sms/cmd/sms-sender/app"
	"github.com/xuchaoi/alertmanager-webhook-sms/cmd/sms-sender/app/option"
	"k8s.io/component-base/cli/flag"
	"k8s.io/klog"
	"os"
)

func main() {
	klog.InitFlags(nil)

	s := option.NewSMSSenderOption()
	s.AddFlags(pflag.CommandLine)
	flag.InitFlags()

	if err := app.Run(s); err != nil {
		os.Exit(1)
	}
}
