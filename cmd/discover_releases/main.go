package main

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v45/github"
)

func main() {
	event := os.Getenv("DRONE_BUILD_EVENT")
	duration, err := time.ParseDuration(os.Getenv("DURATION"))
	if err != nil {
		panic(err)
	}

	client := github.NewClient(nil)
	releases, _, _ := client.Repositories.ListReleases(context.Background(), "drone", "drone", nil)

	// Get tag names
	var tagNames []string
	for _, release := range releases {
		releaseDuration := time.Since(release.PublishedAt.Time)
		if releaseDuration < duration && !*release.Prerelease {
			tagNames = append(tagNames, *release.TagName)
		}
	}

	// retain last tag if event isn't cron and tagNames is empty
	if len(tagNames) == 0 && event != "cron" {
		tagNames = append(tagNames, *releases[0].TagName)
	}

	// Write tag name to tags.txt
	if err := os.WriteFile("tags.txt", []byte(strings.Join(tagNames, "\n")), 0644); err != nil {
		panic(err)
	}
}
