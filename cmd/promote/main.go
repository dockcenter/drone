package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/mod/semver"
)

type PromotionParams struct {
	TAG         string
	DOCKER_TAGS string
}

func main() {
	// Read tags.txt
	bytes, err := os.ReadFile("tags.txt")
	if err != nil {
		panic(err)
	}
	str := string(bytes)

	var commands []string
	for _, tag := range strings.Split(str, "\n") {
		if strings.TrimSpace(tag) == "" {
			continue
		}

		var param PromotionParams
		param.TAG = tag

		// Build docker tags
		var tags []string
		if semver.IsValid(tag) {
			tags = append(tags, slice(semver.Canonical(tag)))
			tags = append(tags, slice(semver.MajorMinor(tag)))
			tags = append(tags, slice(semver.Major(tag)))
			tags = append(tags, "latest")
		}
		param.DOCKER_TAGS = strings.Join(tags, ",")

		// Build drone promote command
		cmd := "drone build promote $DRONE_REPO $DRONE_BUILD_NUMBER $ENVIRONMENT --param=TAG=" + param.TAG + " --param=DOCKER_TAGS=" + param.DOCKER_TAGS
		commands = append(commands, cmd)
	}

	// write commands to promote.sh
	if err := os.WriteFile("scripts/promote.sh", []byte(strings.Join(commands, "\n")), 0644); err != nil {
		panic(err)
	}

	if len(commands) == 0 {
		fmt.Println("No tags to promote")
	}
}

func slice(str string) string {
	if len(str) > 0 {
		return str[1:]
	}
	return str
}
