package main

import (
	"errors"
	"os"
	"os/exec"
	"sort"
	"strings"
)

func GetLatestTag() (string, error) {
	_, err := exec.LookPath("git")
	if err != nil {
		return "", errors.New("Error found git")
	}

	cmd := exec.Command("git", "fetch", "--tag", "-v")

	for _, env := range os.Environ() {
		cmd.Env = append(cmd.Env, env)
	}

	err = cmd.Run()
	if err != nil {
		return "", errors.New("Error fetch tags")
	}

	output, err := exec.Command("git", "tag").Output()
	if err != nil {
		return "", err
	}

	output_str := strings.TrimSuffix(string(output), "\n")
	tags := strings.Split(output_str, "\n")

	if len(tags) == 0 {
		return "0.0.0", errors.New("No exist tag, use 0.0.0")
	}

	var versions VersionList
	for _, tag := range tags {
		tag = strings.TrimPrefix(tag, "v")
		v, _ := NewVersion(tag)
		versions = append(versions, v)
	}

	sort.Sort(versions)
	latest := len(versions)
	return versions[latest-1].String(), nil
}
