package util

import (
	"github.com/samber/lo"
)

func GetFullname(scope string, name string) string {
	if lo.IsNotEmpty(scope) {
		return scope + "/" + name
	} else {
		return name
	}
}
