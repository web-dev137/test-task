package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"

	"github.com/natefinch/lumberjack"
	log "github.com/sirupsen/logrus"
	"github.com/web-dev137/test-task/handler"
	"github.com/web-dev137/test-task/repository"
)

type CfgDB struct {
	PgHost     string
	PgDbName   string
	PgUser     string
	PgPassword string
	PgPort     int
}

type CfgLogger struct {
	LogFilename   string
	LogMaxSize    int
	LogMaxBackups int
	LogMaxAge     int
}

/*
*Setting db connect for postgres
 */
func Init(cfg *CfgDB) (*sql.DB, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s port=%d dbname=%s sslmode=disable", cfg.PgHost, cfg.PgUser, cfg.PgPassword, cfg.PgPort, cfg.PgDbName)
	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Connected failed ping failed")
		return nil, err
	}
	return db, nil
}

/*
*Setting logger
 */
func SetLog(cfg *CfgLogger) {
	log.SetOutput(&lumberjack.Logger{
		Filename:   cfg.LogFilename,
		MaxSize:    cfg.LogMaxSize,
		MaxBackups: cfg.LogMaxBackups,
		MaxAge:     cfg.LogMaxAge,
	})
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.JSONFormatter{})
}
func main() {
	cfgDB := &CfgDB{
		"localhost",
		"test_task_db",
		"postgres",
		"postgres",
		5432,
	}
	logCfg := CfgLogger{"logs/server.log", 10, 10, 30}
	db, err := Init(cfgDB)
	SetLog(&logCfg)
	if err != nil {
		log.WithFields(log.Fields{"error": err}).Error("Connected failed")
	}
	fmt.Println("connect success")
	repo := *repository.NewRepo(db) //init repo
	h := handler.NewHandler(&repo)  //init handler
	router := gin.Default()
	router.POST("/get-items", h.GetItems)
	router.Run(":8080")
}
