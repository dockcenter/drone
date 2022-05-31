package main

import (
	"os"
	"os/exec"
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

	for _, tag := range strings.Split(str, "\n") {
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

		// Run drone promote command
		cmd := exec.Command("drone", "build", "promote", "$DRONE_REPO", "$DRONE_BUILD_NUMBER", "$ENVIRONMENT", "--param=TAG="+param.TAG, "--param=DOCKER_TAGS="+param.DOCKER_TAGS)
		if _, err := cmd.Output(); err != nil {
			panic(err)
		}
	}
}

func slice(str string) string {
	if len(str) > 0 {
		return str[1:]
	}
	return str
}
