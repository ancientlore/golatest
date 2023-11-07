package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

const dlURL = "https://go.dev/dl/"

func main() {
	resp, err := http.Get(dlURL + "?mode=json")
	if err != nil {
		log.Print(err)
		return
	}
	defer resp.Body.Close()

	var releases []release
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&releases)
	if err != nil {
		log.Print(err)
		return
	}

	goarch := strings.TrimSpace(os.Getenv("GOARCH"))
	if goarch == "" {
		var r bytes.Buffer
		cmd := exec.Command("go", "env", "GOARCH")
		cmd.Stdout = &r
		if err = cmd.Run(); err != nil {
			log.Print(err)
			return
		}
		goarch = strings.TrimSpace(r.String())
	}

	goos := strings.TrimSpace(os.Getenv("GOOS"))
	if goos == "" {
		var r bytes.Buffer
		cmd := exec.Command("go", "env", "GOOS")
		cmd.Stdout = &r
		if err = cmd.Run(); err != nil {
			log.Print(err)
			return
		}
		goos = strings.TrimSpace(r.String())
	}

	for i := range releases {
		if releases[i].Stable {
			lookFor := []string{"installer", "archive", "source"}
			found := false
			for j := range lookFor {
				if !found {
					for _, f := range releases[i].Files {
						if f.Arch == goarch && f.OS == goos && f.Kind == lookFor[j] {
							fmt.Printf("go version %s %s/%s %s\n", f.Version, f.OS, f.Arch, dlURL+f.Filename)
							found = true
							break
						}
					}
				}
			}
		}
	}
}
