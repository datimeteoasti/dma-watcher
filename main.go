package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/radovskyb/watcher"
)

const (
	pollingDurationMs = 1000
)

func main() {
	w := watcher.New()

	// SetMaxEvents to 1 to allow at most 1 event's to be received
	// on the Event channel per watching cycle.
	//
	// If SetMaxEvents is not set, the default is to send all events.
	w.SetMaxEvents(10)

	// Only notify for create events
	w.FilterOps(watcher.Create)

	// Only notify for json files
	r := regexp.MustCompile("[a-z0-9-_.]+.json$")
	w.AddFilterHook(watcher.RegexFilterHook(r, false))

	go func() {
		for {
			select {
			case event := <-w.Event:
				err := readWSFile(event.Path)
				if err != nil {
					log.Println(err)
				}
				log.Printf("%s\n", event)
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Watch test folder recursively for changes
	if err := w.AddRecursive("./test_folder"); err != nil {
		log.Fatalln(err)
	}

	// Print all files and folders currently being watched
	for path, f := range w.WatchedFiles() {
		fmt.Printf("%s: %s\n", path, f.Name())
	}

	// Start watching the provided files and folders
	// It'll check for changes every 100ms.
	log.Printf("Start watching files and folders every %d ms...", pollingDurationMs)
	if err := w.Start(time.Millisecond * pollingDurationMs); err != nil {
		log.Fatalln(err)
	}
}

// readWSFile read the content of the file uploaded by the weather station
func readWSFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	log.Println(b)

	return nil
}
