package handle

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/xuchaoi/alertmanager-webhook-sms/cmd/sms-sender/app/option"
	_ "github.com/go-sql-driver/mysql"
	"k8s.io/klog"
)

func WechatHandle(smsContent string, mysqlCfg option.MysqlConfiguration) (string, error) {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", mysqlCfg.UserName, mysqlCfg.Password, mysqlCfg.Network,
		mysqlCfg.Server, mysqlCfg.Port, mysqlCfg.DBName)
	klog.V(4).Infof("Connect to mysql, conn: %s", conn)
	db, err := sql.Open("mysql", conn)
	if err != nil {
		e := fmt.Sprintf("Failed to connect to mysql, err: %v", err)
		return "", errors.New(e)
	}
	defer db.Close()
	klog.V(4).Infof("The executed sql statement: %s", mysqlCfg.InsertSql)

	res, err := db.Exec(mysqlCfg.InsertSql, smsContent)
	if err != nil {
		e := fmt.Sprintf("Failed to exec sql, sql: %s, err: %v", mysqlCfg.InsertSql, err)
		return "", errors.New(e)
	} else {
		info := fmt.Sprintf("Successed to exec sql, sql: %s, res: %v", mysqlCfg.InsertSql, res)
		return info, nil
	}

}