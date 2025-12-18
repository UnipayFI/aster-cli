package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dateTimeLayout = "2006-01-02 15:04:05"

func ParseArgs(args []string) (params map[string]string) {
	params = make(map[string]string, len(args))

	var key string
	for len(args) > 0 {
		arg := strings.TrimLeft(args[0], "-")
		kv := strings.Split(arg, "=")
		if len(kv) == 1 {
			if key == "" {
				key = kv[0]
			} else {
				params[key] = kv[0]
				key = ""
			}
		}
		if len(kv) == 2 {
			params[kv[0]] = kv[1]
		}

		args = args[1:]
	}
	return
}

func ParseTimeFlag(flagName string, value string) (time.Time, bool, error) {
	v := strings.TrimSpace(value)
	if v == "" || v == "0" {
		return time.Time{}, false, nil
	}

	if isAllDigits(v) {
		ms, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return time.Time{}, false, invalidTimeFlagError(flagName)
		}
		if ms < 0 {
			return time.Time{}, false, invalidTimeFlagError(flagName)
		}
		return time.UnixMilli(ms), true, nil
	}

	t, err := time.ParseInLocation(dateTimeLayout, v, time.Local)
	if err != nil {
		return time.Time{}, false, invalidTimeFlagError(flagName)
	}
	return t, true, nil
}

func ParseTimeFlagUnixMilli(flagName string, value string) (int64, bool, error) {
	t, ok, err := ParseTimeFlag(flagName, value)
	if err != nil || !ok {
		return 0, ok, err
	}
	return t.UnixMilli(), true, nil
}

func invalidTimeFlagError(flagName string) error {
	name := strings.TrimSpace(flagName)
	if name == "" {
		name = "time"
	}
	return fmt.Errorf("invalid %s: expected unix milliseconds timestamp (e.g. 1734495381000) or datetime \"YYYY-MM-DD HH:MM:SS\" (e.g. \"2025-12-18 04:16:21\")", name)
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
