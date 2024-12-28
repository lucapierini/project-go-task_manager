# project-go-task_manager

docker run --name "postgres-task_manager-container" -e POSTGRES_USER=userTest -e POSTGRES_PASSWORD=task123456 -e POSTGRES_DB="go-task_manager" -p 5432:5432 -d postgres