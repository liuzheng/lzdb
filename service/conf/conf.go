package conf

import (
	"flag"
	log "github.com/liuzheng712/golog"
	"path/filepath"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"fmt"
	"os"
)

var (
	Yaml = flag.String("c", "example.yml", "Config file")
	GenYaml = flag.Bool("gen", false, "gen config file")
)

func init() {
	flag.Parse()
	log.Initial()

	if *GenYaml {
		GenerateYamlFile()
		os.Exit(0)
	}
	log.Info("config", "loading config %v", *Yaml)
	filename, err := filepath.Abs(*Yaml)
	if err != nil {
		log.Error("config", "filepath.Abs error: %v", err)
		panic(err)
	}
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Error("config", "ReadFile error: %v", err)
		panic(err)
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		log.Error("config", "yaml.Unmarshal : %v", err)
		panic(err)
	}

}

var Config ConfigStruct = ConfigStruct{
	Http:HttpConfig{
		Port:5000,
		Host:"127.0.0.1",
	},
	Grpc:GrpcConfig{
		Port:5001,
		Host:"127.0.0.1",
	},
}

type ConfigStruct struct {
	Http HttpConfig `yaml:"http"`
	Grpc GrpcConfig `yaml:"grpc"`
}
type HttpConfig  struct {
	Port uint16 `yaml:"port"`
	Host string `yaml:"host"`
}
type GrpcConfig  struct {
	Port uint16 `yaml:"port"`
	Host string `yaml:"host"`
}

func GenerateYamlFile() {
	out, err := yaml.Marshal(Config)
	if err != nil {
		log.Error("config", "yaml.Marshal : %v", err)
		panic(err)
	}
	fmt.Print(string(out))
}