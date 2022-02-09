package plugin

import "plugin"

func LoadPlugin(pluginPaths []string) ([]*plugin.Plugin, error) {
	pluginPaths, err := Find(&Filter{
		UseDeepTraversal: true,
		UseRegex:         true,
		IsFile:           false,
		BasePath:         "/Users/yui/IdeaProjects/logstash/lib/",
		Search:           "^.+\\.rb$",
	})
	if err != nil {
		panic("find plugin error!")
	}
	var plugins []*plugin.Plugin
	for _, path := range pluginPaths {
		plugin, err := plugin.Open(path)
		if err != nil {
			panic("load plugin error!")
		}
		plugins = append(plugins, plugin)
	}
	return plugins, nil
}
