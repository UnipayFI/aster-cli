package common

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dateTimeLayout = "2006-01-02 15:04:05"

// FormatUnixTime renders an int64 timestamp as a UTC string. The unit (s / ms /
// μs / ns) is sniffed by magnitude so that endpoints returning different units
// for the same conceptual field (some Aster spot/futures account endpoints
// return ns, others ms) all render correctly. Zero or negative values produce
// an empty string.
func FormatUnixTime(v int64) string {
	if v <= 0 {
		return ""
	}
	var t time.Time
	switch {
	case v >= 1e18:
		t = time.Unix(0, v)
	case v >= 1e15:
		t = time.UnixMicro(v)
	case v >= 1e12:
		t = time.UnixMilli(v)
	default:
		t = time.Unix(v, 0)
	}
	return t.UTC().Format(dateTimeLayout)
}

// FormatTime renders a time.Time as a UTC string in the project layout. Used
// in tandem with FormatUnixTime so all CLI output is timezone-consistent.
func FormatTime(t time.Time) string {
	if t.IsZero() {
		return ""
	}
	return t.UTC().Format(dateTimeLayout)
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
