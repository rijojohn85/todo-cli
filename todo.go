package todo

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type Todos []item

func (t *Todos) Add(task string) {
	todo := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}
	*t = append(*t, todo)
}

func (t *Todos) Toggle(index int) error {
	ls := *t
	err := ls.validateIndex(index)
	if err != nil {
		return err
	}
	if !ls[index-1].Done {
		ls[index-1].Done = true
		ls[index-1].CompletedAt = time.Now()
	} else {
		ls[index-1].Done = false
		ls[index-1].CompletedAt = time.Time{}
	}
	return nil
}

func (t *Todos) validateIndex(index int) error {
	ls := *t
	// new_t := make(Todos, len(ls)-1)
	if index <= 0 || index > len(ls) {
		return errors.New("invalid index")
	}
	return nil
}

func (t *Todos) Delete(index int) error {
	ls := *t
	err := ls.validateIndex(index)
	if err != nil {
		return err
	}
	*t = append(ls[:index-1], ls[index:]...)
	return nil
}

func (t *Todos) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return err
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todos) Store(filename string) error {
	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todos) Print() {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	colFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("#", "Task", "Done", "Created At", "Completed At")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(colFmt)

	for i, t := range *t {
		completed := "❌"
		completedAt := "-"

		if t.Done {
			completed = "✅"
			completedAt = t.CompletedAt.Format(time.RFC850)
		}
		tbl.AddRow(strconv.Itoa(i+1), t.Task, completed, t.CreatedAt.Format(time.RFC850), completedAt)
	}

	tbl.Print()
}

func (t *Todos) Edit(input string) error {
	ls := *t
	parts := strings.SplitN(input, ":", 2)

	if len(parts) != 2 {
		return errors.New("invalid format. please enter id:new_task")
	}
	index, con_err := strconv.Atoi(parts[0])
	if con_err != nil {
		return errors.New("invalid format. please enter id:new_task")
	}

	err := ls.validateIndex(index)
	if err != nil {
		return err
	}
	ls[index-1].Task = parts[1]
	return nil
}
