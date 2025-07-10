package main

import (
	"github.com/tsukinoko-kun/timy/cmd"
	"github.com/tsukinoko-kun/timy/db"
)

func main() {
	defer db.Close()
	cmd.Execute()
}
