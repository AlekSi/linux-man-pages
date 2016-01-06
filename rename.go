package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	vF := flag.Bool("v", false, "be verbose")
	flag.Parse()

	filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		dir, file := filepath.Split(path)
		if !strings.HasPrefix(file, "linux_") {
			path2 := filepath.Join(dir, "linux_"+file)
			if *vF {
				log.Printf("%s -> %s", path, path2)
			}
			err = os.Rename(path, path2)
		}
		return err
	})
}
