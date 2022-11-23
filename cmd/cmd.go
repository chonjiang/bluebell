package cmd

import "flag"

type Flags struct {
	configFile string
}

var f = new(Flags)

func init() {
	flag.StringVar(&f.configFile, "config", "./config.yaml", "配置文件路径")
	flag.Parse()
}

func (f *Flags) configPath() string {
	return f.configFile
}

func GetConfigFileName() string {
	return f.configPath()
}
