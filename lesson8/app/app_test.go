package app_test

import (
	program "lesson8/app"
	"lesson8/config"
	"lesson8/files"
	"log"
	"testing"
)

func Example() {
	cnfg, err := config.NewAppConfig()
	if err != nil {
		log.Fatal(err)
	}
	uniqueFiles := files.NewUniqueFilesMap()

	program := program.NewProgram(cnfg, uniqueFiles)
	err = program.Start()
	if err != nil {
		log.Fatal(err)
	}
	// Output:
	//Program starts searching for duplicate files in "."...
	//Found 2 unique files and 0 duplicates:
	//program.go
	//program_test.go
}

func BenchmarkProgram_Start(b *testing.B) {
	for j := 0; j < b.N; j++ {
		cnfg, err := config.NewAppConfig()
		if err != nil {
			log.Fatal(err)
		}
		cnfg.PrintResult = false
		uniqueFiles := files.NewUniqueFilesMap()

		program := program.NewProgram(cnfg, uniqueFiles)
		err = program.Start()
		if err != nil {
			log.Fatal(err)
		}
	}
}
