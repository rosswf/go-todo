package todo

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type TaskServer struct {
	taskList *TaskList
	http.Handler
}

func NewTaskServer(taskList *TaskList) *TaskServer {
	p := new(TaskServer)
	p.taskList = taskList

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/tasks", func(r chi.Router) {
		r.Get("/", p.tasksHandler)
	})

	p.Handler = r
	return p
}

func (p *TaskServer) tasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := p.taskList.GetAll()
	if err != nil {
		log.Printf("Could not get tasks %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	encoder := json.NewEncoder(w)
	err = encoder.Encode(tasks)
	if err != nil {
		log.Printf("Could not encode json %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
