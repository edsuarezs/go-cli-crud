package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

)

// Definición de la estructura Task que representa una tarea
type Task struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}

// Función para listar las tareas
func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas")
		return
	}

	for _, task := range tasks {
		// Determina el estado de completitud de la tarea y muestra la información
		status := "  "
		if task.Complete {
			status = "✅"
		}
		fmt.Printf("[%s] %d %s\n", status, task.ID, task.Name)
	}
}

// Función para agregar una nueva tarea
func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:       GetNextID(tasks),
		Name:     name,
		Complete: false,
	}
	return append(tasks, newTask)
}

// Función para eliminar una tarea por su ID
func DeleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks
}

// Función para marcar una tarea como completa o incompleta
func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Complete = !task.Complete
			break
		}
	}
	return tasks
}

// Función para guardar las tareas en un archivo
func SaveTasks(file *os.File, tasks []Task) {
	// Serializa las tareas en formato JSON
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}

	// Reinicia la posición del puntero al inicio del archivo y lo trunca
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}

	// Escribe los bytes en el archivo y realiza un flush
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}

	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

// Función para obtener el próximo ID disponible
func GetNextID(tasks []Task) int {
	if len(tasks) == 0 {
		return 1
	}
	return tasks[len(tasks)-1].ID + 1
}
