package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func getStatusLsusb() map[string]string {
	result := make(map[string]string)
	//r, error := regexp.Compile("([0-9a-f]{4}:[0-9a-f]{4})(.*)(Serial|Hub|root)?.*")
	r, error := regexp.Compile(" ([0-9a-f]{4}:[0-9a-fA-Z]{4}) (.*)")
	if error != nil {
		log.Panic(error)
	}
	cmd := exec.Command("lsusb")
	var out io.Reader

	out, err := cmd.StdoutPipe()

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	defer cmd.Process.Kill()
	s := bufio.NewScanner(out)
	nstr := "  \t\n"
	for s.Scan() {
		s := s.Text()
		if len(s) > 0 {

			m := r.FindStringSubmatch(s)

			if len(m) > 2 {
				fmt.Printf("%s - %s\n", m[1],
					strings.Trim(strings.Split(m[2], "Serial")[0], nstr))
				result[m[1]] = strings.Trim(strings.Split(m[2], "Serial")[0], nstr)
			}
		}
	}
	log.Println("Done.")
	return result
}

func main() {
	lsusb := getStatusLsusb()
	b, err := json.Marshal(lsusb)
	if err != nil {
		fmt.Println("error:", err)
	}
	//https://blog.golang.org/go-maps-in-action
	os.Stdout.Write(b)
}
