package todo

import (
	"testing"
)

const filename = "todo_test.json"

func TestStore(t *testing.T) {
	task := "hello=rijo"
	todos := &Todos{}
	todos.Add(task)
	err := todos.Store(filename)
	if err != nil {
		t.Fatal(err)
	}

	read_todos := &Todos{}

	error := read_todos.Load(filename)
	if error != nil {
		t.Fatal(error)
	}
	ls := *read_todos
	if task != ls[0].Task {
		t.Errorf("Expcted %v, got %v", task, ls[0].Task)
	}
}

func TestAdd(t *testing.T) {
	t.Helper()
	t.Run("Add Task Test", func(t *testing.T) {
		task := "hello=paul"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Store(filename)
		if err != nil {
			t.Fatal(err)
		}
		todos.Load(filename)
		ls := *todos
		if task != ls[0].Task {
			t.Errorf("Expcted %v, got %v", task, ls[0].Task)
		}
	})
}
