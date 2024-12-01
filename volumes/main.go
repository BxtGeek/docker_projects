package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"
)

// Task represents a single todo item
type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	Completed   bool
}

// TodoManager manages the collection of tasks
type TodoManager struct {
	mu      sync.RWMutex
	tasks   []Task
	nextID  int
}

// AddTask adds a new task to the list
func (tm *TodoManager) AddTask(description string) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	task := Task{
		ID:          tm.nextID,
		Description: description,
		CreatedAt:   time.Now(),
		Completed:   false,
	}
	tm.tasks = append(tm.tasks, task)
	tm.nextID++
}

// ToggleTask marks a task as complete or incomplete
func (tm *TodoManager) ToggleTask(id int) {
	tm.mu.Lock()
	defer tm.mu.Unlock()

	for i, task := range tm.tasks {
		if task.ID == id {
			tm.tasks[i].Completed = !tm.tasks[i].Completed
			break
		}
	}
}

// GetSortedTasks returns tasks sorted with incomplete tasks first
func (tm *TodoManager) GetSortedTasks() []Task {
	tm.mu.RLock()
	defer tm.mu.RUnlock()

	// Create a copy of tasks
	tasks := make([]Task, len(tm.tasks))
	copy(tasks, tm.tasks)

	// Custom sorting: 
	// 1. Incomplete tasks come first
	// 2. Within each group (incomplete/complete), maintain original order
	sort.Slice(tasks, func(i, j int) bool {
		if tasks[i].Completed != tasks[j].Completed {
			return !tasks[i].Completed
		}
		return i < j
	})

	return tasks
}

// Global todo manager
var todoManager = &TodoManager{}

func main() {
	// Route handlers
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/add-task", addTaskHandler)
	http.HandleFunc("/toggle-task", toggleTaskHandler)

	// Start the server
	fmt.Println("Server starting on http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

// homeHandler renders the main page with tasks
func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	
	tasks := todoManager.GetSortedTasks()
	tmpl.Execute(w, tasks)
}

// addTaskHandler handles adding new tasks
func addTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	taskDescription := r.FormValue("task")
	if taskDescription != "" {
		todoManager.AddTask(taskDescription)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// toggleTaskHandler handles marking tasks as complete/incomplete
func toggleTaskHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	taskID := r.FormValue("id")
	if taskID != "" {
		id := 0
		fmt.Sscanf(taskID, "%d", &id)
		todoManager.ToggleTask(id)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
