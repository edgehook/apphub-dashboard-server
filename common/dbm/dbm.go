package dbm

import (
	"context"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"k8s.io/klog/v2"
)

// Connect the sqlite3
func GormSQLite(dbPath string) *gorm.DB {
	klog.Infof("Connecting to sqlite3 %s", dbPath)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		Logger: &CustomLogger{},
	})
	if err != nil {
		klog.Fatalf("Connect to sqlite3 failed! %s", err.Error())
		return nil
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	return db
}

type CustomLogger struct {
}

func (ct *CustomLogger) LogMode(level logger.LogLevel) logger.Interface {
	return ct
}

func (ct *CustomLogger) Info(context context.Context, val string, i ...interface{}) {

}

func (ct *CustomLogger) Warn(context context.Context, val string, i ...interface{}) {

}

func (ct *CustomLogger) Error(context context.Context, val string, i ...interface{}) {
	// klog.Errorf("gorm db error: %v", val)
}
func (ct *CustomLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	// if err != nil {
	// 	sql, _ := fc()
	// 	klog.Errorf("sql: %v, error: %v", sql, err.Error())
	// }
}
