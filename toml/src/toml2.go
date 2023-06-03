package main

import (
	"os"
	"fmt"
	"github.com/BurntSushi/toml"
)

type tomlConfig struct {
	Conf map[string]any
}

	
func check(e error) {
    if e != nil {
        panic(e)
    }
}


func isType(a, b interface{}) bool {
    return fmt.Sprintf("%T", a) == fmt.Sprintf("%T", b)
}


func PrintDict(k1 string, v1 interface{}) {
    fmt.Printf("Key:%s\n",k1)
    if isType(v1,  make(map[string]any)) {
        data := v1.(map[string]interface{})
        for k2, v2 := range data {
          PrintDict(k2, v2)
        }
    } else {
       fmt.Printf("Value:%s\n",v1)
    }
}


func main() {
	var Conf map[string]any
        data, err := os.ReadFile("config.toml")
        if err != nil {
                fmt.Printf("ERROR: reading (config.toml).")
                os.Exit(1)
        }
        err=toml.Unmarshal(data, &Conf)

        for k, v := range Conf {
                PrintDict(k, v)
        }
}
