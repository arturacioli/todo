package storage

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/arturacioli/todo/internal/task"
)

func LoadTasks()([]task.Task,error){

	file,err := os.OpenFile("todo.csv", os.O_RDWR|os.O_CREATE, 0664)
	if err != nil{
		fmt.Printf("Ocorreu um erro: %v\n", err)
		return nil,err
	}
	defer file.Close()


	reader := csv.NewReader(file)
	tasks := []task.Task{}
	for {
		record, err := reader.Read()
		if errors.Is(err, io.EOF){
			break
		}
		if err != nil{
			return nil, err
		}
		if len(record) < 4 {
		continue
		}
		t, err := time.Parse(time.RFC3339, record[3])
		if err != nil{
			return nil,err
		}

		task := task.Task{
			Id: record[0],
			Task: record[1],
			Complete: record[2] == "true",
			AddedAt: t,
		}

		tasks = append(tasks, task)
	}
	return tasks,nil

}

func WriteTasks(tasks []task.Task) error{
	
	file, err := os.Create("todo.csv")
	if err != nil{
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _,t := range tasks{
		record := []string{
			t.Id,
			t.Task,
			strconv.FormatBool(t.Complete),
			t.AddedAt.Format(time.RFC3339),
		}

		if err := writer.Write(record); err != nil{
			return err
		}
	}

	return nil
}

func GetNextId(tasks []task.Task) (int,error){
	lastID := 0
	for _, t := range tasks {
		id, err := strconv.Atoi(t.Id)
		if err != nil {
			continue
		}

		if id > lastID {
			lastID = id
		}
	}

	return lastID + 1, nil
}
