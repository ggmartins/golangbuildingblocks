package main

import (
  "fmt"
  "sync"
  "time"
)

type Data struct {
  comm chan string
  continue_comm int
  timeout int
  wg0 sync.WaitGroup
  wg1 sync.WaitGroup
}

func (d *Data) init () {
  d.comm = make(chan string)
  d.wg1.Add(4)
  d.wg0.Add(2)
  d.continue_comm = 1
}

func (d *Data)Worker(id int)  {
  var i int
  for{
    i++
    fmt.Println("ID: ", id, "  inc:", i)
    if d.continue_comm == 1 {
       d.comm <- fmt.Sprintf("inc:%d",i)
    } else {
      break //if discontinued, stop producing too
    }
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
  fmt.Println("Finished.")
}

func (d *Data)mainthread()  {
  go d.Worker(1)
  go d.Worker(2)
  go d.Worker(3)
  go d.Worker(4)

  d.WaitThreads()
}

func (d *Data) ReadChannel()  {

  for d.continue_comm == 1 {
    select {
      case v, ok := <-d.comm:
          if ok {
              fmt.Printf("%s was read.\n", v)
              d.timeout=0
          } else {
              fmt.Println("Quitting...")
              break;
          }
      case <-time.After(time.Second * 1):
         d.timeout++
         if d.timeout > 10 && d.timeout < 20{
           fmt.Println("Warning: idle input, quitting in ", 20 - d.timeout)
         } else if d.timeout > 20 {
           d.continue_comm = 0
         }
     }
  }

  d.wg0.Done()
}

func (d *Data) ReadInput()  {
  for{
    fmt.Println("cmd:")

    var s string
    fmt.Scanln(&s)
    fmt.Println(s)

    if(s=="quit"){
      fmt.Println("cmd: Quitting..")
      d.continue_comm = 0
      break
    } else {
      d.timeout = 0
    }
  }
}

func main()  {
  var d Data

  d.init()

  fmt.Println("Waiting... ")
  go d.mainthread()
  go d.ReadInput()
  go d.ReadChannel()

  d.WaitFinished()

}
