package main

type TestPlugin struct {
}

func (t *TestPlugin) test(input []byte) (data string, err error) {
	return "plugin run!" + string(input), nil
}
