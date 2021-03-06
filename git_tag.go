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

	old_tags, _ := exec.Command("git", "tag", "-l").Output()

	old_tags_str := strings.TrimSuffix(string(old_tags), "\n")

	for _, old_tag := range strings.Split(old_tags_str, "\n") {
		exec.Command("git", "tag", "-d", old_tag).Run()
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

	var raw_tags []string
	for _, tag := range tags {
		if tag != "" {
			raw_tags = append(raw_tags, tag)
		}
	}

	var versions VersionList
	for _, tag := range raw_tags {
		tag = strings.TrimPrefix(tag, "v")
		v, _ := NewVersion(tag)
		versions = append(versions, v)
	}

	if len(versions) == 0 {
		return "0.0.0", errors.New("No exist tag, use 0.0.0")
	}

	sort.Sort(versions)
	latest := len(versions)
	return versions[latest-1].String(), nil
}

func GetNewVersion() (string, error) {
	tag, err := GetLatestTag()
	if err != nil && tag == "" {
		return "", err
	}

	v, _ := NewVersion(tag)
	v.build += 1
	new_tag := v.String()
	return new_tag, nil
}
