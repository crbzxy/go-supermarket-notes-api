# Supermarket and Notes API

## Descripción

Este es un servidor de API para la gestión de listas de supermercado y notas personales. Está construido con Go y utiliza MongoDB como base de datos. La API incluye la documentación generada con Swagger.

## Características

- Crear, leer, actualizar y eliminar listas de supermercado.
- Crear, leer, actualizar y eliminar notas.
- Gestión de usuarios.
- Documentación de API generada con Swagger.

## Requisitos

- Go 1.16 o superior
- MongoDB (puedes usar MongoDB Atlas para una base de datos en la nube)
- Una cuenta en MongoDB Atlas (si usas MongoDB en la nube)

## Instalación

1. Clona el repositorio:

   ```bash
   git clone https://github.com/tu-usuario/go-supermarket-notes-api.git
   cd go-supermarket-notes-api
   ```

2. Instala las dependencias:

   ```bash
   go mod tidy
   ```

3. Crea un archivo `.env` en el directorio raíz del proyecto con el siguiente contenido:

   ```env
   MONGODB_URI=mongodb+srv://tu-usuario:tu-contraseña@tu-cluster.mongodb.net/?retryWrites=true&w=majority&appName=TuApp
   DB_NAME=nombre_de_tu_base_de_datos
   PORT=8080
   ```

   Asegúrate de reemplazar `tu-usuario`, `tu-contraseña`, `tu-cluster` y `nombre_de_tu_base_de_datos` con los valores correctos.

4. Genera la documentación de Swagger:

   ```bash
   swag init -g cmd/main.go
   ```

## Uso

1. Inicia el servidor:

   ```bash
   go run cmd/main.go
   ```

2. La API estará disponible en `http://localhost:8080`.

3. Puedes acceder a la documentación de Swagger en `http://localhost:8080/swagger/index.html`.

## Rutas de la API

### Listas de Supermercado

- `POST /lists`: Crear una nueva lista.
- `GET /lists`: Obtener todas las listas.
- `GET /lists/{id}`: Obtener una lista por ID.
- `PUT /lists/{id}`: Actualizar una lista por ID.
- `DELETE /lists/{id}`: Eliminar una lista por ID.

### Notas

- `POST /notes`: Crear una nueva nota.
- `GET /notes`: Obtener todas las notas.
- `GET /notes/{id}`: Obtener una nota por ID.
- `PUT /notes/{id}`: Actualizar una nota por ID.
- `DELETE /notes/{id}`: Eliminar una nota por ID.

### Usuarios

- `POST /users`: Crear un nuevo usuario.
- `GET /users`: Obtener todos los usuarios.
- `GET /users/{id}`: Obtener un usuario por ID.
- `PUT /users/{id}`: Actualizar un usuario por ID.
- `DELETE /users/{id}`: Eliminar un usuario por ID.

## Estructura del Proyecto

    ```bash



go-supermarket-notes-api/
├── cmd/
│ └── main.go
├── config/
│ └── db.go
├── controllers/
│ ├── listController.go
│ ├── noteController.go
│ └── userController.go
├── docs/
│ ├── docs.go
├── models/
│ ├── list.go
│ ├── note.go
│ └── user.go
├── routes/
│ └── routes.go
├── services/
│ ├── listService.go
│ ├── noteService.go
│ └── userService.go
├── utils/
│ └── response.go
├── .env
├── .gitignore
├── go.mod
└── go.sum
```

# go-supermarket-notes-api
