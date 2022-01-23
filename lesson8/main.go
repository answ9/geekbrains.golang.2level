package main

import (
	p "lesson8/app"
	"lesson8/config"
	f "lesson8/files"
	"log"
)

func main() {
	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	uniqueFiles := f.NewUniqueFilesMap()

	program := p.NewProgram(cnfg, uniqueFiles)
	err = program.Start()
	if err != nil {
		log.Fatal(err)
	}
}
