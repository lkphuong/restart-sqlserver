package main

import (
	"log"
	"os/exec"
	"strings"
	"time"

	"github.com/robfig/cron"
)

func checkServiceRunning(serviceName string) (bool, error) {
	cmd := exec.Command("cmd", "/C", "sc", "query", serviceName)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return false, err
	}
	return strings.Contains(string(output), "RUNNING"), nil
}

func restartSqlServer() error {
	log.Println("[ALERT] MSSQLSERVER is not running. Attempting to restart...")

	stopAgentCmd := exec.Command("cmd", "/C", "net", "stop", "SQLSERVERAGENT", "/yes")
	stopCmd := exec.Command("cmd", "/C", "net", "stop", "MSSQLSERVER", "/yes")
	startCmd := exec.Command("cmd", "/C", "net", "start", "MSSQLSERVER")
	startAgentCmd := exec.Command("cmd", "/C", "net", "start", "SQLSERVERAGENT")

	if output, err := stopAgentCmd.CombinedOutput(); err != nil {
		log.Printf("Error stopping SQLSERVERAGENT: %v\nOutput: %s", err, output)
	}
	if output, err := stopCmd.CombinedOutput(); err != nil {
		log.Printf("Error stopping MSSQLSERVER: %v\nOutput: %s", err, output)
	}

	time.Sleep(5 * time.Second)

	if output, err := startCmd.CombinedOutput(); err != nil {
		log.Printf("Error starting MSSQLSERVER: %v\nOutput: %s", err, output)
	}
	if output, err := startAgentCmd.CombinedOutput(); err != nil {
		log.Printf("Error starting SQLSERVERAGENT: %v\nOutput: %s", err, output)
	}

	log.Println("[INFO] Restart sequence executed.")
	return nil
}

func main() {
	job := cron.New()

	job.AddFunc("0 10 * * *", func() {
		const maxRetries = 3
		for attempt := 1; attempt <= maxRetries; attempt++ {
			running, err := checkServiceRunning("MSSQLSERVER")
			if err != nil {
				log.Printf("[ERROR] Failed to check MSSQLSERVER status: %v", err)
				return
			}

			if running {
				log.Println("[OK] MSSQLSERVER is running. No action needed.")
				return
			}

			log.Printf("[WARN] MSSQLSERVER not running. Attempt %d/%d", attempt, maxRetries)

			if err := restartSqlServer(); err != nil {
				log.Printf("[ERROR] Failed to restart MSSQLSERVER on attempt %d: %v", attempt, err)
			}

			time.Sleep(10 * time.Second)
			runningAfterRestart, err := checkServiceRunning("MSSQLSERVER")
			if err != nil {
				log.Printf("[ERROR] Failed to check MSSQLSERVER after restart: %v", err)
				return
			}
			if runningAfterRestart {
				log.Println("[SUCCESS] MSSQLSERVER restarted successfully.")
				return
			}
		}

		log.Println("[FAILED] MSSQLSERVER could not be restarted after maximum retries.")
	})

	job.Start()
	select {}
}
