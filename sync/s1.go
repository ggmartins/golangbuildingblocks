package main

import (
  "fmt"
  "sync"
  "time"
)

type Data struct {
  comm chan string
  wg0 sync.WaitGroup
  wg1 sync.WaitGroup
}

func (d *Data) init () {
  d.comm = make(chan string)
  d.wg1.Add(4)
  d.wg0.Add(2)
}

func (d *Data)Worker(id int)  {
  var i int
  for{
    i++
    fmt.Println("ID: ", id, "  inc:", i)
    time.Sleep(time.Second)
    if(i>=10) {break}
  }
  fmt.Println("ID: ", i, "  :quit")
  d.wg1.Done()
}

func (d *Data)WaitThreads()  {
  fmt.Println("Waiting 1...")
  d.wg1.Wait()
  fmt.Println("Done 1 ...")
  d.wg0.Done()
}

func (d *Data)WaitFinished()  {
  fmt.Println("Finishing..")
  d.wg0.Wait()
  fmt.Println("Finished..")
}

func (d *Data)mainthread()  {
  go d.Worker(1)
  go d.Worker(2)
  go d.Worker(3)
  go d.Worker(4)

  d.WaitThreads()
}

func (d *Data) Input()  {

  for{
    fmt.Println("cmd:")

    var s string
    fmt.Scanln(&s)
    fmt.Println(s)

    if(s=="quit"){
      fmt.Println("Quitting")
      break
    }
  }
  d.wg0.Done()

}

func main()  {
  var d Data

  d.init()

  fmt.Println("Waiting... ")
  go d.mainthread()
  go d.Input()

  d.WaitFinished()

}
