package browser

import (
	"fmt"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor *int
	Build *int
	Patch *int
}

func VersionFromString(ver string) (Version, error) {
	parts := strings.Split(ver, ".")
	var v Version
	if len(parts) > 0 {
		major, err := strconv.Atoi(parts[0])
		if err != nil {
			return v, err
		}
		v.Major = major
	}
	if len(parts) > 1 {
		minor, err := strconv.Atoi(parts[1])
		if err != nil {
			return v, err
		}
		v.Minor = &minor
	}
	if len(parts) > 2 {
		build, err := strconv.Atoi(parts[2])
		if err != nil {
			return v, err
		}
		v.Build = &build
	}
	if len(parts) > 3 {
		patch, err := strconv.Atoi(parts[3])
		if err != nil {
			return v, err
		}
		v.Patch = &patch
	}
	return v, nil
}

func (v Version) String() string {
	out := fmt.Sprint(v.Major)
	if v.Minor != nil {
		out += fmt.Sprintf(".%d", *v.Minor)
	}
	if v.Build != nil {
		out += fmt.Sprintf(".%d", *v.Build)
	}
	if v.Patch != nil {
		out += fmt.Sprintf(".%d", *v.Patch)
	}
	return out
}
