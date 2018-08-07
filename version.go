package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	major int64
	minor int64
	build int64
}

func NewVersion(v string) (*Version, error) {
	slices := strings.Split(v, ".")
	major, _ := strconv.ParseInt(slices[0], 10, 64)
	minor, _ := strconv.ParseInt(slices[1], 10, 64)
	build, _ := strconv.ParseInt(slices[2], 10, 64)
	return &Version{
		major: major,
		minor: minor,
		build: build,
	}, nil
}

func (v *Version) Compare(other *Version) int {
	if v.major > other.major {
		return 1
	} else if v.major < other.major {
		return -1
	} else {
		if v.minor > other.minor {
			return 1
		} else if v.minor < other.minor {
			return -1
		} else {
			if v.build > other.build {
				return 1
			} else if v.build < other.build {
				return -1
			} else {
				return 0
			}
		}
	}
}

func (v *Version) LessThan(other *Version) bool {
	return v.Compare(other) < 0
}

func (v *Version) String() string {
	return fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.build)
}
