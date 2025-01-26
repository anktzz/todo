package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Tobedone  string     // `json:"tobedone"`
	Completed bool       // `json:"compleated"`
	CreatedAt time.Time  // `json:"CreatedAt"`
	DoneAt    *time.Time // `json:"DoneAt"`
}

type Task []Todo

func (td *Task) add(title string) { //td is reference to actual value of task.
	*td = append(*td, Todo{
		Tobedone:  title,
		Completed: false,
		DoneAt:    nil,
		CreatedAt: time.Now(),
	}) //*td points the type which was referenced by td
}

func (td *Task) delete(index int) {
	if index < 0 || index > len(*td) {
		fmt.Println("Error Invalid Index")
		return
	}
	*td = append((*td)[:index], (*td)[index+1:]...)
}

func (td *Task) change(index int, changed string) {
	if index < 0 || index > len(*td) {
		fmt.Println("Error Invalid Index")
		return
	}
	fmt.Println("Enter modifyied task")
	fmt.Scanln(&changed)
	(*td)[index].Tobedone = changed
}

func (td *Task) toggle(index int) {
	if index < 0 || index > len(*td) {
		fmt.Println("Error Invalid Index")
		return
	}
	isCompleated := (*td)[index].Completed
	if !isCompleated {
		compleationTime := time.Now()
		(*td)[index].DoneAt = &compleationTime
	}
	(*td)[index].Completed = !isCompleated
}

func (td *Task) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "CreatedAt", "DoneAt")
	for index, t := range *td {
		Completed := "❌"
		completedAt := ""
		if t.Completed {
			Completed = "✅"
			if t.DoneAt != nil {
				completedAt = t.DoneAt.Format(time.RFC1123)
			}
		}
		table.AddRow(strconv.Itoa(index), t.Tobedone, Completed, t.CreatedAt.Format(time.RFC1123), completedAt)
	}
	table.Render()
}
