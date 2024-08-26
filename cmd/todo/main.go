package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/rijojohn85/todo-cli"
)

const todoFile = "todos.json"

func main() {
	add := flag.String("add", "", "add a new todo")
	complete := flag.Int("complete", 0, "mark a string as completed")
	del := flag.Int("delete", 0, "delete a task")
	flag.Parse()
	todos := &todo.Todos{}
	if err := todos.Load(todoFile); err != nil {
		print_error(err, 1)
		os.Exit(1)
	}

	switch {
	case len(*add) > 0:
		todos.Add(*add)
		store(*todos)
	case *complete > 0:
		err := todos.Complete(*complete)
		if err != nil {
			print_error(err, 1)
		}
		store(*todos)
	case *del > 0:
		err := todos.Delete(*del)
		if err != nil {
			print_error(err, 1)
		}
		store(*todos)
	default:
		print_error(errors.New("invalid option"), 1)
	}
}

func print_error(err error, code int) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(code)
}

func store(todos todo.Todos) {
	err := todos.Store(todoFile)
	if err != nil {
		print_error(err, 1)
	}
}
