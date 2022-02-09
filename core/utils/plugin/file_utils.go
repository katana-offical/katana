package plugin

import (
	"errors"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"regexp"
)

type Filter struct {
	UseDeepTraversal bool

	UseRegex bool

	IsFile bool

	BasePath string

	Search string
}

func Find(filter *Filter) (paths []string, err error) {
	var pluginPaths []string
	if filter.IsFile {
		//find single path
		pluginPaths = append(pluginPaths, filter.BasePath+filter.Search)
		return pluginPaths, nil
	} else {
		//find multi path
		reg, err := regexp.Compile(filter.Search)
		if err != nil {
			log.Err(errors.New("can not compile regex:" + filter.Search))
			return pluginPaths, nil
		}
		filepath.Walk(filter.BasePath,
			func(path string, f os.FileInfo, err error) error {
				if err != nil {
					log.Err(errors.New("bad basePath:" + filter.BasePath))
					return err
				}
				if f.IsDir() {
					//if use useDeepTraversal
					if filter.UseDeepTraversal {
						return nil
					} else {
						return filepath.SkipDir
					}
				}
				var matched = false
				//find multi path use traversal and regex
				if filter.UseRegex {
					matched = reg.MatchString(f.Name())
				} else {
					//find multi path use traversal and not use regex
					matched = f.Name() == filter.Search
				}
				if matched {
					pluginPaths = append(pluginPaths, path)
				}
				return nil
			})
		return pluginPaths, nil
	}

}
