# project-go-task_manager

docker run --name "postgres-task_manager-container" -e POSTGRES_USER=userTest -e POSTGRES_PASSWORD=task123456 -e POSTGRES_DB="go-task_manager" -p 5432:5432 -d postgres

asignar nuevos roles a usuarios
desasignar roles a usuarios
asignar un usuario a un proyecto
desasignar un usuario a un proyecto
asignar una tarea a un proyecto
desasignar una tarea a un proyecto
implementar docker compose 