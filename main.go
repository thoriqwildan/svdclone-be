package main

import (
	"github.com/thoriqwildan/svdclone-be/pkg/config"
	"github.com/thoriqwildan/svdclone-be/pkg/database"
	"github.com/thoriqwildan/svdclone-be/pkg/helper"
	"github.com/thoriqwildan/svdclone-be/pkg/server"
)

func main() {
	config.LoadEnv()
	helper.InitValidator()
	database.InitDB()
	server.Serve()
}