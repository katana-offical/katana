package utils

import (
	"fmt"
	"katana/core/utils/plugin"
	"testing"
)

func TestFilterFiles(t *testing.T) {
	files, _ := plugin.Find(&plugin.Filter{
		UseDeepTraversal: true,
		UseRegex:         true,
		IsFile:           false,
		BasePath:         "/Users/yui/IdeaProjects/logstash/lib/",
		Search:           "^.+\\.rb$",
	})
	fmt.Println(files)
}
