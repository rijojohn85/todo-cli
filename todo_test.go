package todo

import (
	"os"
	"testing"
)

const filename = "todo_test.json"

func deleteFile() {
	os.Remove(filename)
}

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
		deleteFile()
		t.Fatal(error)
	}
	ls := *read_todos
	if task != ls[0].Task {
		t.Errorf("Expcted %v, got %v", task, ls[0].Task)
	}
	deleteFile()
}

func TestAdd(t *testing.T) {
	t.Helper()
	t.Run("Add Task Test", func(t *testing.T) {
		task := "hello=paul"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Store(filename)
		if err != nil {
			deleteFile()
			t.Fatal(err)
		}
		todos.Load(filename)
		ls := *todos
		if task != ls[0].Task {
			t.Errorf("Expcted %v, got %v", task, ls[0].Task)
		}
		deleteFile()
	})
}

func TestDelete(t *testing.T) {
	t.Helper()
	t.Run("expect error", func(t *testing.T) {
		task := "rijo john"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Delete(2)
		if err == nil {
			deleteFile()
			t.Fatal("Expcted error, did not get one")
		} else {
			if err.Error() != "invalid index" {
				t.Errorf("Expected %q, got %q", "invalid index", err.Error())
			}
		}
		deleteFile()
	})
	t.Run("delete task", func(t *testing.T) {
		task := "rijo john"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Delete(1)
		if err != nil {
			deleteFile()
			t.Fatal(err)
		}
		if len(*todos) != 0 {
			t.Errorf("should have got 0, got %v", len(*todos))
		}
		deleteFile()
	})
}

func TestComplete(t *testing.T) {
	t.Helper()
	t.Run("expect error: complete", func(t *testing.T) {
		task := "rijo john"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Toggle(2)
		if err == nil {
			deleteFile()
			t.Fatal("Expcted error, did not get one")
		} else {
			if err.Error() != "invalid index" {
				t.Errorf("Expected %q, got %q", "invalid index", err.Error())
			}
		}
		deleteFile()
	})
	t.Run("complete task", func(t *testing.T) {
		task := "rijo john"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Toggle(1)
		if err != nil {
			deleteFile()
			t.Fatal(err)
		}
		ls := *todos
		if !ls[0].Done {
			t.Error("Task has not been marked complete")
		}
		deleteFile()
	})
}

func TestEdit(t *testing.T) {
	t.Helper()
	t.Run("expect error: wrong index", func(t *testing.T) {
		task := "rijo john"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Edit("2:hello")
		if err == nil {
			deleteFile()
			t.Fatal("Expcted error, did not get one")
		} else {
			if err.Error() != "invalid index" {
				t.Errorf("Expected %q, got %q", "invalid index", err.Error())
			}
		}
		deleteFile()
	})
	t.Run("expect error: wrong format", func(t *testing.T) {
		task := "rijo john"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Edit("2-hello")
		if err == nil {
			deleteFile()
			t.Fatal("Expcted error, did not get one")
		} else {
			if err.Error() != "invalid format. please enter id:new_task" {
				t.Errorf("Expected %q, got %q", "invalid index", err.Error())
			}
		}
		deleteFile()
	})
	t.Run("expect error: wrong format2", func(t *testing.T) {
		task := "rijo john"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Edit("hello:2")
		if err == nil {
			deleteFile()
			t.Fatal("Expcted error, did not get one")
		} else {
			if err.Error() != "invalid format. please enter id:new_task" {
				t.Errorf("Expected %q, got %q", "invalid index", err.Error())
			}
		}
		deleteFile()
	})
	t.Run("edit: working", func(t *testing.T) {
		task := "rijo john"
		todos := &Todos{}
		todos.Add(task)
		err := todos.Edit("1:hello")
		if err != nil {
			deleteFile()
			t.Fatal(err)
		}
		todos.Store(filename)

		ls := *todos
		ls.Load(filename)
		if ls[0].Task != "hello" {
			t.Errorf("expected %q got %q", "hello", ls[0].Task)
		}
		deleteFile()
	})
}
