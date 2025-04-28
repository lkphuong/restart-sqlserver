package main

import (
	"log"
	"os/exec"

	sqlserver "github.com/lkphuong/restart-sqlserver/configs/databases"
	"github.com/robfig/cron"
)

func main() {

	job := cron.New()
	job.AddFunc("0 10 * * *", func() {
		db := sqlserver.ConnectionSqlServer()

		err := db.Ping()
		log.Default().Println("Ping database success! ", err)

		cmd := exec.Command("cmd", "/C", "restart-sqlservice.bat")
		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Default().Println("Error executing command: ", err)
		}
		log.Default().Println("Command output: ", string(output))

	})
	job.Start()

	select {}
}
