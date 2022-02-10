package service

import (
	"testing"
	"time"

	"github.com/NetEase-Media/ngo/client/db"
	"github.com/stretchr/testify/assert"
)

func TestDataFromDB(t *testing.T) {
	// 方式一 代码写死
	opts := []*db.Options{
		{
			Name:            "db01",
			Url:             "",
			MaxIdleCons:     10,
			MaxOpenCons:     10,
			ConnMaxLifetime: time.Second * 1000,
			ConnMaxIdleTime: time.Second * 60,
		},
	}

	// 方式二 从配置文件读取
	//config.Init("../../configs/app-local.yaml")
	//config.Unmarshal("db", opts)

	db.Init(opts)
	mysqlDB = db.GetMysqlClient("db01")

	id := 1
	car := GetDataFromDB(id, "")
	assert.Equal(t, car.ID, id)
}
