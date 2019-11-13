package main

import (
	"bufio"
	"io"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("ping", "8.8.8.8")
	var out io.Reader
	/*{
	        stdout, err := cmd.StdoutPipe()
	        if err != nil {
	            log.Fatal(err)
	        }
	        stderr, err := cmd.StderrPipe()
	        if err != nil {
	            log.Fatal(err)
	        }
	        out = io.MultiReader(stdout, stderr)
	}*/
	out, err := cmd.StdoutPipe()

	if err = cmd.Start(); err != nil {
		log.Fatal(err)
	}
	defer cmd.Process.Kill()
	s := bufio.NewScanner(out)
	for s.Scan() {
		log.Println(s.Text())
	}
	log.Println("Done.")
}
