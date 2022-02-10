package service

import (
	"fmt"

	"github.com/NetEase-Media/ngo/client/db"

	"gorm.io/gorm"
)

// Test 定义Car的数据模型
type Test struct {
	ID   int    `gorm:"id"`
	Name string `gorm:"name"`
}

// TableName 重写表名, 给gorm 使用
func (t *Test) TableName() string {
	return "test"
}

// GetDataFromDB 在数据库获取数据
func GetDataFromDB(id int, from string) *Test {
	t := Test{}
	var stat *gorm.DB

	client := db.GetMysqlClient(from)
	if client == nil {
		client = mysqlDB
	}

	stat = client.Where("id", id).Find(&t)
	//stat = client.Raw("select * from test where id = ?", id).Find(&t)
	if stat.RowsAffected == 0 || stat.Error != nil {
		fmt.Print("we can not find nothing")
		return nil
	}
	return &t
}

/**
CREATE DATABASE test;

USE test;

CREATE TABLE test(
	id bigint PRIMARY KEY AUTO_INCREMENT,
    name varchar(20) not null
);

INSERT INTO test(name) values ("tesla");
**/
