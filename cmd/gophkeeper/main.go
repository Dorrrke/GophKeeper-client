package main

import (
	"github.com/Dorrrke/GophKeeper-client/cmd"
	_ "github.com/mattn/go-sqlite3"
)

var (
	// buildVersion - версия сборки.
	buildVersion string
	// buildDate - дата сборки.
	buildDate string
	// buildCommit - комментарии к сборке.
	buildCommit string
)

func main() {
	cmd.BuildVersion = buildVersion
	cmd.BuildDate = buildDate
	cmd.BuildCommit = buildCommit
	cmd.Execute()
}
