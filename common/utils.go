package common

import "strings"

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
