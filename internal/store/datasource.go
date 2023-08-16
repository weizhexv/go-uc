package store

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rolandhe/daog"
	"go-uc/vconfig"
)

var db = initDB()

func initDB() daog.Datasource {
	config := daog.DbConf{
		DbUrl:    vconfig.MysqlURL(),
		Size:     vconfig.MysqlMaxOpenConns(),
		Life:     vconfig.MysqlMaxLifeTime(),
		IdleCons: vconfig.MysqlMaxIdleConns(),
		IdleTime: vconfig.MysqlMaxIdleTime(),
		LogSQL:   true,
	}
	fmt.Printf("init db config: %v\n", config)
	datasource, err := daog.NewDatasource(&config)
	if err != nil {
		panic(err)
	}
	return datasource
}

func DB() daog.Datasource {
	return db
}
