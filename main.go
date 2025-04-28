package main

import (
	"log"
	"os/exec"

	sqlserver "github.com/lkphuong/restart-sqlserver/configs/databases"
	"github.com/robfig/cron"
)

func main() {
	db := sqlserver.ConnectionSqlServer()
	job := cron.New()
	job.AddFunc("0 * * * *", func() {
		err := db.Ping()
		log.Default().Println("Ping database success! ", err)

		if err != nil {
			log.Default().Println("Restarting SQL Server...")
			cmd := exec.Command("cmd", "/C", "restart-sqlserver.bat")
			_, err := cmd.CombinedOutput()
			log.Default().Println("Error executing command: ", err)
			log.Default().Println("Done!!!")
		}

	})
	job.Start()

	select {}
}
