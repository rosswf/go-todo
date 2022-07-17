package todo_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/rosswf/go-todo-cli"
)

func TestGETTasks(t *testing.T) {
	data := []todo.Task{
		{
			Id:       1,
			Name:     "Task 1",
			Complete: false,
		},
		{
			Id:       2,
			Name:     "Task 2",
			Complete: false,
		},
	}

	storage := CreateMockStorage(data)
	taskList := todo.CreateTaskList(storage)
	server := todo.NewTaskServer(taskList)

	t.Run("test /tasks returns status 200", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Code
		want := http.StatusOK

		if got != want {
			t.Errorf("got status %d, want %d", got, want)
		}
	})

	t.Run("test /tasks returns a list of tasks", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/tasks", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []todo.Task
		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Could not decode json, %v", err)
		}

		want := data

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got response %+v, want %+v", got, want)
		}
	})
}
