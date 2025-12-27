package main

import (
	"discord-quest-watcher/internal/browser"
	"discord-quest-watcher/internal/quests"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	token, webhook := os.Getenv("TOKEN"), os.Getenv("DISCORD_WEBHOOK_URL")
	rewardFilter := func() string {
		if value := os.Getenv("REWARD_FILTER"); value != "" {
			return value
		}
		return "all"
	}()

	checkInterval := func() int {
		if value := os.Getenv("FETCH_INTERVAL"); value != "" {
			if minutes, err := strconv.Atoi(value); err == nil && minutes > 0 {
				return minutes
			}
			log.Printf("❌ Invalid FETCH_INTERVAL=%s, using default 30", value)
		}
		return 30
	}()

	runOnce := os.Getenv("RUN_ONCE") == "true"

	if token == "" || webhook == "" {
		log.Fatal("🚫 TOKEN and DISCORD_WEBHOOK_URL required")
	}

	log.Printf("🏁 Starting Discord quest monitor with reward_filter=%s, check_interval=%d minutes, run_once=%t", rewardFilter, checkInterval, runOnce)

	// create browser and authenticate once
	br, err := browser.CreateBrowser()
	if err != nil {
		log.Fatalf("❌ Failed to create browser: %v", err)
	}
	defer br.MustClose()

	if err := browser.AuthenticateWithToken(br, token); err != nil {
		log.Fatalf("❌ Failed to authenticate: %v", err)
	}

	for {
		log.Println("🔍 Checking for new quests")
		if err := quests.CheckQuests(br, webhook, rewardFilter, runOnce); err != nil {
			log.Printf("❌ Quest check failed: %v", err)
			if runOnce {
				os.Exit(1)
			}
		}

		if runOnce {
			log.Println("✔ Check complete, exiting 🔚")
			break
		}
		time.Sleep(time.Duration(checkInterval) * time.Minute)
	}
}
