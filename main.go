package main

import (
	"log"

	"github.com/yanggu0t/go-rdbms-practice/config"
	"github.com/yanggu0t/go-rdbms-practice/internal/database"
	"github.com/yanggu0t/go-rdbms-practice/internal/router"
)

func main() {
	// 加載配置
	cfg := config.Load()

	// 初始化數據庫
	db := database.InitDB(cfg)

	// 設置路由
	r := router.SetupRouter(db, cfg)

	// Start server
	log.Fatal(r.Run(":8080"))
}
