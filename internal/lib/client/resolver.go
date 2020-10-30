package client

import (
	"fmt"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func PathResolver(p string, f func(p string) ([]string, error)) ([]string, error) {
	re := regexp.MustCompile("\\[(\\*|[0-9]+)]")

	if !re.MatchString(p) {
		return f(p)
	}

	parts := strings.Split(p, "/")
	paths := []string{}
	dup := map[string]bool{}

	for i := 0; i < len(parts); i++ {
		if strings.HasPrefix(parts[i], "[") {
			s := re.FindStringSubmatch(parts[i])

			if len(s) != 2 {
				return nil, fmt.Errorf(
					"Invalid array expression: %s",
					parts[i],
				)
			}

			pwd := parts[0:i]
			list, err := f(strings.Join(pwd, "/"))

			if err != nil {
				return nil, err
			}

			if s[1] == "*" {
				for _, path := range list {
					parts[i] = filepath.Base(path)
					resolved, err := PathResolver(filepath.Join(parts...), f)

					if err != nil {
						return nil, err
					}

					for _, p := range resolved {
						if !dup[p] {
							paths = append(paths, p)
							dup[p] = true
						}
					}
				}
			} else {
				idx, err := strconv.Atoi(s[1])

				if err != nil {
					return nil, fmt.Errorf(
						"Invalid array expression: %s",
						parts[i],
					)
				}

				if len(list) <= idx {
					return nil, fmt.Errorf(
						"Array index %d not found in list",
						idx,
					)
				}

				parts[i] = filepath.Base(list[idx])
				resolved, err := PathResolver(filepath.Join(parts...), f)

				if err != nil {
					return nil, err
				}

				for _, p := range resolved {
					if !dup[p] {
						paths = append(paths, p)
						dup[p] = true
					}
				}
			}
		}
	}

	return paths, nil
}
