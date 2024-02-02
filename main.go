package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	task "github.com/edsuarezs/go-cli-crud/tasks"
)

func main() {
	// Abre el archivo tasks.json en modo lectura/escritura, creándolo si no existe
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close() // Cierra el archivo al finalizar la función

	var tasks []task.Task

	// Obtiene información sobre el archivo
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}

	// Si el archivo no está vacío, lee su contenido y deserializa en la estructura tasks
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}

		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{} // Inicializa la estructura tasks si el archivo está vacío
	}

	// Verifica si se proporciona mas de un argumento en la línea de comandos
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	// Realiza acciones según el comando proporcionado en la línea de comandos
	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Cual es tu tarea?")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)

		tasks = task.AddTask(tasks, name)
		task.SaveTasks(file, tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar un ID para eliminar")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un numero")
			return
		}
		tasks = task.DeleteTask(tasks, id)
		task.SaveTasks(file, tasks)
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Debes proporcionar un ID para completar")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("El ID debe ser un numero")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTasks(file, tasks)
	default:
		printUsage()
	}
}

// Imprime el uso del programa en la línea de comandos
func printUsage() {
	fmt.Println("Uso: go-clid-crud [list|add|complete|delete]")
}
