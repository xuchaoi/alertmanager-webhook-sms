package option

import (
	"github.com/spf13/pflag"
	"strconv"
)

const DefaultSenderPort = 8060

type SMSSenderOptions struct {
	SenderPort int
	SMSCfg     SMSConfiguration
}

type SMSConfiguration struct {
	Code          string
	Content       string
	NewStaffId    string
	EffectiveDate string
	SubPort       string
}

func NewSMSSenderOption() *SMSSenderOptions {
	option := SMSSenderOptions{
		SenderPort:    DefaultSenderPort,
	}
	return &option
}

func (o *SMSSenderOptions) AddFlags(fs *pflag.FlagSet)  {
	fs.IntVar(&o.SenderPort, "senderPort", o.SenderPort, "SMS sender's port, default: "+ strconv.Itoa(DefaultSenderPort))
	fs.StringVar(&o.SMSCfg.Code, "smsCode", o.SMSCfg.Code, "SMS to send user's mobile phone number (Required). Such as: 18312344321")
	fs.StringVar(&o.SMSCfg.Content, "smsContent", o.SMSCfg.Content, "SMS content, default: \"\"")
	fs.StringVar(&o.SMSCfg.NewStaffId, "newStaffId", o.SMSCfg.NewStaffId, "Employee ID, default: \"\"")
	fs.StringVar(&o.SMSCfg.EffectiveDate, "effectiveDate", o.SMSCfg.EffectiveDate, "Effective date, leave it blank for immediate delivery, default: \"\"")
	fs.StringVar(&o.SMSCfg.SubPort, "subPort", o.SMSCfg.SubPort, "SMS source (Required). Such as: 10086")
}