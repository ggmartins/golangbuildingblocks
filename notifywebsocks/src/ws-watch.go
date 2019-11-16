package main

import (
	"os"
	"fmt"
	"net/http"
	"flag"
	"bufio"
	"strings"
	"github.com/fsnotify/fsnotify"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", ":7077", "http service address")
var upgrader = websocket.Upgrader{} 

func handleWS(w http.ResponseWriter, r *http.Request) {
        c, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
                fmt.Print("upgrade:", err)
                return
        }
        defer c.Close()
        for {
                mt, message, err := c.ReadMessage()
                if err != nil {
                        fmt.Println("read:", err)
                        break
                }
                fmt.Printf("recv: %s", message)
                err = c.WriteMessage(mt, message)
                if err != nil {
                        fmt.Println("write:", err)
                        break
                }
        }
}


func main() {
	flag.Parse()
	var f* os.File
	var errfile error
	var scanner* bufio.Scanner

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		fmt.Println("ERROR", err)
	}
	defer watcher.Close()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		handleWS(w, r)
	})
	go func() {
		for {
			select {
			case event := <-watcher.Events:
				fmt.Printf("EVENT! %#v\n", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					if strings.HasPrefix( event.Name, "/tmp/tmp.ta." ) {
						fmt.Println("WRITE:", event.Name)
						if f == nil {
							f, errfile = os.Open(event.Name) // os.OpenFile has more options if you need them
							defer f.Close()
							if errfile != nil { // error checking is good practice
								fmt.Println(errfile)
							}
						        scanner = bufio.NewScanner(f)
						}
						for scanner.Scan() {
							fmt.Println(scanner.Text())
						}
						if errscanner := scanner.Err(); errscanner != nil {
							fmt.Println(errscanner)
						}
				}
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					if strings.HasPrefix( event.Name, "/tmp/tmp.ta." ) {
						fmt.Println("CREATE:", event.Name)
					}
				}
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	if err := watcher.Add("/tmp/"); err != nil {
		fmt.Println("ERROR", err)
	}

	errhttp := http.ListenAndServe(*addr, nil)
	if errhttp != nil {
		fmt.Println("ListenAndServe: " + errhttp.Error())
	}
}

