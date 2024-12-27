package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Task represents a single todo item
type Task struct {
	ID          int
	Description string
	CreatedAt   time.Time
	Completed   bool
}

// TodoDatabase manages database operations
type TodoDatabase struct {
	db *sql.DB
}

// InitDatabase creates and sets up the database connection
func InitDatabase() *TodoDatabase {
	// Retrieve database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construct connection string
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", 
		dbUser, dbPass, dbHost, dbPort, dbName)

	// Open database connection
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Verify connection
	if err := db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	// Create tasks table if not exists
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS tasks (
			id INT AUTO_INCREMENT PRIMARY KEY,
			description VARCHAR(255) NOT NULL,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			completed BOOLEAN DEFAULT FALSE
		)
	`)
	if err != nil {
		log.Fatalf("Error creating tasks table: %v", err)
	}

	return &TodoDatabase{db: db}
}

// AddTask inserts a new task into the database
func (td *TodoDatabase) AddTask(description string) error {
	_, err := td.db.Exec(
		"INSERT INTO tasks (description) VALUES (?)", 
		description,
	)
	return err
}

// ToggleTask updates the completion status of a task
func (td *TodoDatabase) ToggleTask(id int) error {
	_, err := td.db.Exec(
		"UPDATE tasks SET completed = NOT completed WHERE id = ?", 
		id,
	)
	return err
}

// GetTasks retrieves all tasks from the database
func (td *TodoDatabase) GetTasks() ([]Task, error) {
	// Query to get tasks, with incomplete tasks first
	rows, err := td.db.Query(`
		SELECT id, description, created_at, completed 
		FROM tasks 
		ORDER BY completed ASC, created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		if err := rows.Scan(
			&task.ID, 
			&task.Description, 
			&task.CreatedAt, 
			&task.Completed,
		); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// Close closes the database connection
func (td *TodoDatabase) Close() {
	td.db.Close()
}

// Global database manager
var todoDB *TodoDatabase

func main() {
	// Initialize database connection
	todoDB = InitDatabase()
	defer todoDB.Close()

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
	
	tasks, err := todoDB.GetTasks()
	if err != nil {
		http.Error(w, "Error fetching tasks", http.StatusInternalServerError)
		return
	}
	
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
		if err := todoDB.AddTask(taskDescription); err != nil {
			http.Error(w, "Error adding task", http.StatusInternalServerError)
			return
		}
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
		if err := todoDB.ToggleTask(id); err != nil {
			http.Error(w, "Error toggling task", http.StatusInternalServerError)
			return
		}
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
