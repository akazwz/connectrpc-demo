package database

import (
	"fmt"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"

	"opendrive/models"
)

var (
	db *gorm.DB
)

// 默认的数据库文件名
const defaultDBFile = "test.db"

func init() {
	if err := SetupDatabase(defaultDBFile, nil); err != nil {
		panic(fmt.Sprintf("数据库初始化失败: %v", err))
	}
	if err := db.AutoMigrate(models.User{}, models.RefreshToken{}); err != nil {
		panic(fmt.Sprintf("数据库自动迁移失败: %v", err))
	}
}

// GetDataBase 返回数据库连接实例
func GetDataBase() *gorm.DB {
	return db
}

// SetupDatabase 设置数据库连接，可自定义数据库文件名和GORM配置
func SetupDatabase(dbFile string, config *gorm.Config) error {
	if dbFile == "" {
		dbFile = defaultDBFile
	}

	if config == nil {
		config = &gorm.Config{}
	}

	var err error
	gormDB, err := gorm.Open(sqlite.Open(dbFile), config)
	if err != nil {
		return fmt.Errorf("打开数据库连接失败: %w", err)
	}

	// 检查数据库连接
	sqlDB, err := gormDB.DB()
	if err != nil {
		return fmt.Errorf("获取底层数据库连接失败: %w", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return fmt.Errorf("数据库连接测试失败: %w", err)
	}
	db = gormDB
	return nil
}
