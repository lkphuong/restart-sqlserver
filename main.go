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
	job.AddFunc("0 10 * * *", func() {
		err := db.Ping()
		log.Default().Println("Ping database success! ", err)

		if err != nil {
			log.Default().Println("Restarting SQL Server...")

			stopCmd := exec.Command("cmd", "/C", "net stop MSSQLSERVER")
			stopAgentCmd := exec.Command("cmd", "/C", "net stop SQLSERVERAGENT")
			startCmd := exec.Command("cmd", "/C", "net start MSSQLSERVER")
			startAgentCmd := exec.Command("cmd", "/C", "net start SQLSERVERAGENT")

			stopOutput, err := stopCmd.CombinedOutput()
			if err != nil {
				log.Printf("Error stopping MSSQLSERVER: %v\n", err)
			}
			log.Printf("Stop MSSQLSERVER: %s\n", stopOutput)

			stopAgentOutput, err := stopAgentCmd.CombinedOutput()
			if err != nil {
				log.Printf("Error stopping SQLSERVERAGENT: %v\n", err)
			}
			log.Printf("Stop SQLSERVERAGENT: %s\n", stopAgentOutput)

			startOutput, err := startCmd.CombinedOutput()
			if err != nil {
				log.Printf("Error starting MSSQLSERVER: %v\n", err)
			}
			log.Printf("Start MSSQLSERVER: %s\n", startOutput)

			startAgentOutput, err := startAgentCmd.CombinedOutput()
			if err != nil {
				log.Printf("Error starting SQLSERVERAGENT: %v\n", err)
			}
			log.Printf("Start SQLSERVERAGENT: %s\n", startAgentOutput)

			log.Default().Println("Done!!!")
		}

	})
	job.Start()

	select {}
}
