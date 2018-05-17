package main

import (
  "fmt"
  "os"
  "plugin"
  "dataunit"
)


type Module interface {
  RunMod(in <-chan *dataunit.DataUnit)
}


func main() {


  out := make(chan *dataunit.DataUnit, 2000)
  plug, err := plugin.Open("./module.so")
  if err != nil {
    fmt.Printf("FATAL (plugin.Open): " + err.Error())
    os.Exit(1)
  }

  symModule, err := plug.Lookup("Module")
  if err != nil {
    fmt.Printf(err.Error())
    os.Exit(1)
  }

  var module Module
  module, ok:= symModule.(Module)
  if !ok {
    fmt.Println("unexpected type from module symbol")
    os.Exit(1)
  }

  module.RunMod(out)
}

