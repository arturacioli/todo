package task

import "time"

type Task struct{
	Id string
	Task string
	Complete bool
	AddedAt	time.Time 
}
