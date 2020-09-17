package option

import (
	"github.com/spf13/pflag"
	"strconv"
)

const (
	DefaultSenderPort       = 8060
	DefaultSMSNewStaffId    = "101"
	DefaultSMSEffectiveDate = ""
	DefaultMysqlUsername    = "root"
	DefaultMysqlPassword    = "1qaz!QAZ"
	DefaultMysqlPort        = 3306
	DefaultMysqlNetwork     = "tcp"
	DefaultMysqlDBName      = "mysql"
	DefaultSubPort          = "10086"
)

type SMSSenderOptions struct {
	SenderPort    int
	SMSCfg        SMSConfiguration
	MysqlCfg      MysqlConfiguration
}

type SMSConfiguration struct {
	Url           string
	Code          string
	Content       string
	NewStaffId    string
	EffectiveDate string
	SubPort       string
	CrmpfPubInfo  CrmpfPubInfo
}

type CrmpfPubInfo struct {
	CityCode      string
	CountryCode   string
	StaffId       string
	OrgId         string
}

type MysqlConfiguration struct {
	UserName      string
	Password	  string
	Network       string
	Server        string
	Port          int
	DBName        string
	InsertSql     string
}

func NewSMSSenderOption() *SMSSenderOptions {
	DefaultSMSCrmpfPubInfo := CrmpfPubInfo{}
	option := SMSSenderOptions{
		SenderPort: DefaultSenderPort,
		SMSCfg: SMSConfiguration{
			NewStaffId: DefaultSMSNewStaffId,
			EffectiveDate: DefaultSMSEffectiveDate,
			SubPort: DefaultSubPort,
			CrmpfPubInfo: DefaultSMSCrmpfPubInfo,
		},
		MysqlCfg: MysqlConfiguration{
			UserName: DefaultMysqlUsername,
			Password: DefaultMysqlPassword,
			Port: DefaultMysqlPort,
			Network: DefaultMysqlNetwork,
			DBName: DefaultMysqlDBName,
		},
	}
	return &option
}

func (o *SMSSenderOptions) AddFlags(fs *pflag.FlagSet)  {
	fs.IntVar(&o.SenderPort, "senderPort", o.SenderPort,
		"SMS sender's port, default: " + strconv.Itoa(DefaultSenderPort))
	fs.StringVar(&o.SMSCfg.Url, "smsUrl", o.SMSCfg.Url,
		"SMS API Url for sending SMS (Required). Such as: http://10.19.190.91:8080/ngcrmpfcore_sd/csf/standard/userSms")
	fs.StringVar(&o.SMSCfg.Code, "smsCode", o.SMSCfg.Code,
		"SMS to send user's mobile phone number (Required). Such as: 18312344321")
	fs.StringVar(&o.SMSCfg.Content, "smsContent", o.SMSCfg.Content, "SMS content, default: \"\"")
	fs.StringVar(&o.SMSCfg.NewStaffId, "newStaffId", o.SMSCfg.NewStaffId, "Employee ID, default: \"\"")
	fs.StringVar(&o.SMSCfg.EffectiveDate, "effectiveDate", o.SMSCfg.EffectiveDate,
		"Effective date, leave it blank for immediate delivery, default: \"\"")
	fs.StringVar(&o.SMSCfg.SubPort, "subPort", o.SMSCfg.SubPort, "SMS source (Required). Such as: 10086")
	fs.StringVar(&o.SMSCfg.CrmpfPubInfo.CityCode, "crmpfCityCode", o.SMSCfg.CrmpfPubInfo.CityCode,
		"SMS crmpfPubInfo cityCode")
	fs.StringVar(&o.SMSCfg.CrmpfPubInfo.CountryCode, "crmpfCountryCode", o.SMSCfg.CrmpfPubInfo.CountryCode,
		"SMS crmpfPubInfo countryCode")
	fs.StringVar(&o.SMSCfg.CrmpfPubInfo.StaffId, "crmpfStaffId", o.SMSCfg.CrmpfPubInfo.StaffId,
		"SMS crmpfPubInfo staffId")
	fs.StringVar(&o.SMSCfg.CrmpfPubInfo.OrgId, "crmpfOrgId", o.SMSCfg.CrmpfPubInfo.OrgId,
		"SMS crmpfPubInfo orgId")
	fs.StringVar(&o.MysqlCfg.UserName, "mysqlUserName", o.MysqlCfg.UserName,
		"The user name of the mysql which receives the alert SMS")
	fs.StringVar(&o.MysqlCfg.Password, "mysqlPassword", o.MysqlCfg.Password,
		"The user password of the mysql which receives the alert SMS")
	fs.StringVar(&o.MysqlCfg.Network, "mysqlNetwork", o.MysqlCfg.Network,
		"The network type of the mysql which receives the alert SMS")
	fs.StringVar(&o.MysqlCfg.Server, "mysqlServer", o.MysqlCfg.Server,
		"The server address of the mysql which receives the alert SMS")
	fs.IntVar(&o.MysqlCfg.Port, "mysqlPort", o.MysqlCfg.Port,
		"The port of the mysql which receives the alert SMS, default: " + strconv.Itoa(DefaultMysqlPort))
	fs.StringVar(&o.MysqlCfg.DBName, "mysqlDBName", o.MysqlCfg.DBName,
		"The db name of the mysql which receives the alert SMS")
	fs.StringVar(&o.MysqlCfg.InsertSql, "mysqlInsertSql", o.MysqlCfg.InsertSql,
		"The insert sql of the mysql which receives the alert SMS")
}