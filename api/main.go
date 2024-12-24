package main

import (
	"github.com/lucapierini/project-go-task_manager/config"
)

func init(){
	config.LoadEnvVariables()
	config.ConnectDB()
	config.SyncDB()
}

func main() {

}