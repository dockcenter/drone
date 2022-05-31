package main

import (
	"context"
	"os"
	"strings"
	"time"

	"github.com/google/go-github/v45/github"
)

func main() {
	client := github.NewClient(nil)
	releases, _, _ := client.Repositories.ListReleases(context.Background(), "drone", "drone", nil)

	// Get tag names
	var tagNames []string
	for _, release := range releases {
		duration := time.Since(release.PublishedAt.Time)
		hours := duration.Hours()
		if hours < 24 && !*release.Prerelease {
			tagNames = append(tagNames, *release.TagName)
		}
	}

	// Write tag name to tags.txt
	if err := os.WriteFile("tags.txt", []byte(strings.Join(tagNames, "\n")), 0644); err != nil {
		panic(err)
	}
}
