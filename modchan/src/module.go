package main

import ( 
  "fmt"
  "dataunit"
)
type module string 
func (m module) RunMod(in <-chan *dataunit.DataUnit) {
    fmt.Println("Running module. ")
    n := <-in
    fmt.Printf("out: %v\n", n)
}
var Module module
