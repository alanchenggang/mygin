package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)
import "gorm.io/driver/mysql"

var GLOBAL_DB *gorm.DB

func init() {

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       "root:coder4jdocker@tcp(112.126.71.240:3340)/gorm_class?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize:         256,                                                                                            // string 类型字段的默认长度
		DisableDatetimePrecision:  true,                                                                                           // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,                                                                                           // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,                                                                                           // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,                                                                                          // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{
		// 跳过默认事务
		SkipDefaultTransaction: false,
		// 是否复数表名
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_", // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: false, // 使用单数表名，启用该选项后，`User` 的表名应该是 `t_user`
		},
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束 在代码中体现外键
		PrepareStmt:                              true,
	})
	// 获取通用数据库对象 sql.DB ，然后使用其提供的功能
	sqlDB, err := db.DB()

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		panic(err)
	}
	//migrator := db.Migrator()
	//// 创建表
	//_ = migrator.CreateTable(&User{})
	//// 判断表是否存在
	//hasTable := migrator.HasTable(&User{})
	//fmt.Println(hasTable)
	// 删除表
	//_ = migrator.DropTable(&User{})
	//_ = db.AutoMigrate(&User{})

	GLOBAL_DB = db
}
