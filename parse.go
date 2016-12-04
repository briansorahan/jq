package jq

import "strings"

func Parse(in string) (Op, error) {
	segments := strings.Split(in, ".")

	ops := make([]Op, 0, len(segments))
	for _, segment := range segments {
		key := strings.TrimSpace(segment)
		if key == "" {
			continue
		}

		ops = append(ops, Dot(key))
	}

	return Chain(ops...), nil
}
