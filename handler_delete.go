package main

import (
	"fmt"

	"github.com/arturacioli/todo/internal/storage"
)

func (cli *Cli)HandlerDelete(args []string){
	if len(args) < 1{
		fmt.Println("Not enough arguments!")
		return
	}
	id := args[0]
	found := false

	for idx := range cli.Tasks{
		if cli.Tasks[idx].Id == id{
			cli.Tasks = append(cli.Tasks[:idx], cli.Tasks[idx+1:]...)
			found = true
			break
		}
	}

	if !found{
		fmt.Printf("Task not found!\n")
		return
	}
	
	err := storage.WriteTasks(cli.Tasks)
	if err != nil{
		fmt.Printf("Error saving after deletion :%v\n",err)
		return
	}

	fmt.Println("Task deleted!")
}
