# Go CLI CRUD

Este es un programa simple de línea de comandos (CLI) en Go para realizar operaciones CRUD (Crear, Leer, Actualizar, Eliminar) en una lista de tareas almacenada en un archivo JSON.

## Requisitos

- Go instalado en tu sistema. Puedes descargarlo [aquí](https://golang.org/dl/).

## Uso

1. Clona o descarga este repositorio en tu máquina.
2. Abre una terminal y navega al directorio del proyecto.
3. Ejecuta el programa con el siguiente comando:

```bash
   go run main.go [comando]
```

Sustituye [comando] con uno de los siguientes:

- list: Lista todas las tareas.
- add: Agrega una nueva tarea.
- complete: Marca una tarea como completa o incompleta.
- delete: Elimina una tarea por su ID.

Ejemplo:

```bash
go run main.go list
```

Ejemplos de Comandos

- Listar todas las tareas:

```bash
go run main.go list
```

- Agregar una nueva tarea:

```bash
go run main.go add
```

Se te pedirá ingresar el nombre de la tarea.

- Marcar una tarea como completa o incompleta:

```bash
go run main.go complete [ID]
```

Sustituye [ID] con el ID de la tarea que deseas marcar.

- Eliminar una tarea por su ID:

```bash
go run main.go delete [ID]
```

Sustituye [ID] con el ID de la tarea que deseas eliminar.

## Contribuir

¡Siéntete libre de contribuir! Cualquier mejora o corrección de errores es bienvenida. Abre un issues o envía una pull request.

## Licencia

Este proyecto está bajo la Licencia MIT.
