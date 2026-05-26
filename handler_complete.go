package main

import (
	"fmt"

	"github.com/arturacioli/todo/internal/storage"
)

func (cli *Cli)HandlerComplete(args []string){
	if len(args) < 1{
		fmt.Println("Not enough arguments!")
		return
	}
	id := args[0]

	for idx := range cli.Tasks{
		if cli.Tasks[idx].Id == id{
			cli.Tasks[idx].Complete = true 	
			storage.WriteTasks(cli.Tasks)
			fmt.Printf("Task marked as complete!\n")
			return
		}
	}

}
