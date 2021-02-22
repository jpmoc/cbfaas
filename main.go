package main

import (
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type conf struct {
    Version int64 `yaml:"version"`
    Project string `yaml:"project"`
    App struct {
        name string `yaml:"name"`
        ports []int `yaml:"ports"`
    }
}

func (c *conf) getConf() *conf {

    yamlFile, err := ioutil.ReadFile("manifest.yml")
    if err != nil {
        log.Printf("yamlFile.Get err   #%v ", err)
    }
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        log.Fatalf("Unmarshal: %v", err)
    }

    return c
}

func main() {
    var c conf
    c.getConf()

    fmt.Println(c)
}
