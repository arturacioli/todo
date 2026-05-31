package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/arturacioli/todo/internal/storage"
	"github.com/arturacioli/todo/internal/task"
)



func (cli *Cli)HandlerAdd(args [] string){
	if len(args) < 1 {
		fmt.Println("Usage: add <task description>")
		return
	}
	nextId, err := storage.GetNextId(cli.Tasks)
	if err != nil{
		fmt.Printf("Erro ao gerar id %v\n",err)
		return
	}

	row := task.Task{
		Id: strconv.Itoa(nextId),
		Task: args[0],
		Complete: false,
		AddedAt: time.Now(),
	}

	cli.Tasks = append(cli.Tasks, row)
	err = storage.WriteTasks(cli.Tasks)
	if err != nil {
		fmt.Printf("Erro ao salvar tasks: %v\n", err)
		return
	}

	fmt.Println("Task added!")

}


