package main

import (
	cmd "github.com/Dorrrke/GophKeeper-client/internal/commands"
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
