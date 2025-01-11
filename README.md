# project-go-task_manager

docker run --name "postgres-task_manager-container" -e POSTGRES_USER=userTest -e POSTGRES_PASSWORD=task123456 -e POSTGRES_DB="go-task_manager" -p 5432:5432 -d postgres

asignar nuevos roles a usuarios
desasignar roles a usuarios
asignar un usuario a un proyecto
desasignar un usuario a un proyecto
asignar una tarea a un proyecto
desasignar una tarea a un proyecto
implementar docker compose 

<!-- [GIN-debug] POST   /api/auth/register        --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).Register-fm (4 handlers) -->
<!-- [GIN-debug] POST   /api/auth/login           --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).Login-fm (4 handlers) -->
<!-- [GIN-debug] POST   /api/auth/refresh         --> github.com/lucapierini/project-go-task_manager/handlers.RefreshTokenHandler (4 handlers) -->
<!-- [GIN-debug] POST   /api/admin/roles/         --> github.com/lucapierini/project-go-task_manager/handlers.(*RoleHandler).CreateRole-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/roles/         --> github.com/lucapierini/project-go-task_manager/handlers.(*RoleHandler).ListRoles-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/roles/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*RoleHandler).GetRole-fm (5 handlers) -->
<!-- [GIN-debug] PUT    /api/admin/roles/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*RoleHandler).UpdateRole-fm (5 handlers) -->
<!-- [GIN-debug] DELETE /api/admin/roles/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*RoleHandler).DeleteRole-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/users/         --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).ListUsers-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/users/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).GetUser-fm (5 handlers) -->
<!-- [GIN-debug] PUT    /api/admin/users/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).UpdateUser-fm (5 handlers) -->
<!-- [GIN-debug] DELETE /api/admin/users/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).DeleteUser-fm (5 handlers) -->
<!-- [GIN-debug] POST   /api/admin/users/:id_user/:id_role --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).AddRoleToUser-fm (5 handlers) -->
<!-- [GIN-debug] DELETE /api/admin/users/:id/:id_role --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).RemoveRoleFromUser-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/projects/      --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).ListProjects-fm (5 handlers) -->
<!-- [GIN-debug] POST   /api/admin/projects/      --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).CreateProject-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/projects/:id   --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).GetProjectById-fm (5 handlers) -->
<!-- [GIN-debug] PUT    /api/admin/projects/:id   --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).UpdateProject-fm (5 handlers) -->
<!-- [GIN-debug] DELETE /api/admin/projects/:id   --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).DeleteProject-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/projects/user/:id --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).ListProjectsByUserId-fm (5 handlers) -->
<!-- [GIN-debug] POST   /api/admin/projects/:id/user/:userId --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).AddUserToProject-fm (5 handlers) -->
<!-- [GIN-debug] DELETE /api/admin/projects/:id/user/:userId --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).RemoveUserFromProject-fm (5 handlers) -->
<!-- [GIN-debug] POST   /api/admin/projects/:id/task --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).AddTaskToProject-fm (5 handlers) -->
<!-- [GIN-debug] DELETE /api/admin/projects/:id/task/:taskId --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).RemoveTaskFromProject-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/tasks/         --> github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).ListTasks-fm (5 handlers) -->
<!-- [GIN-debug] POST   /api/admin/tasks/         --> github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).CreateTask-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/admin/tasks/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).GetTaskById-fm (5 handlers) -->
<!-- [GIN-debug] PUT    /api/admin/tasks/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).UpdateTask-fm (5 handlers) -->
<!-- [GIN-debug] DELETE /api/admin/tasks/:id      --> github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).DeleteTask-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/users/:id            --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).GetUser-fm (6 handlers) -->
<!-- [GIN-debug] PUT    /api/users/:id            --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).UpdateUser-fm (6 handlers) -->
<!-- [GIN-debug] DELETE /api/users/:id            --> github.com/lucapierini/project-go-task_manager/handlers.(*UserHandler).DeleteUser-fm (6 handlers) -->
<!-- [GIN-debug] POST   /api/projects/            --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).CreateProject-fm (5 handlers) -->
<!-- [GIN-debug] GET    /api/projects/:id         --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).GetProjectById-fm (6 handlers) -->
<!-- [GIN-debug] PUT    /api/projects/:id         --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).UpdateProject-fm (6 handlers) -->
<!-- [GIN-debug] DELETE /api/projects/:id         --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).DeleteProject-fm (6 handlers) -->
<!-- [GIN-debug] GET    /api/projects/user/:id    --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).ListProjectsByUserId-fm (6 handlers) -->
<!-- [GIN-debug] POST   /api/projects/:id/user/:userId --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).AddUserToProject-fm (6 handlers) -->
<!-- [GIN-debug] DELETE /api/projects/:id/user/:userId --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).RemoveUserFromProject-fm (6 handlers) -->
<!-- [GIN-debug] POST   /api/projects/:id/task    --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).AddTaskToProject-fm (6 handlers) -->
<!-- [GIN-debug] DELETE /api/projects/:id/task/:taskId --> github.com/lucapierini/project-go-task_manager/handlers.(*ProjectHandler).RemoveTaskFromProject-fm (6 handlers) -->
<!-- [GIN-debug] POST   /api/tasks/               github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).CreateTask-fm (5 handlers) --> -->
<!-- [GIN-debug] GET    /api/tasks/:id            --> github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).GetTaskById-fm (6 handlers) -->
<!-- [GIN-debug] PUT    /api/tasks/:id            --> github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).UpdateTask-fm (6 handlers) -->
<!-- [GIN-debug] DELETE /api/tasks/:id            --> github.com/lucapierini/project-go-task_manager/handlers.(*TaskHandler).DeleteTask-fm (6 handlers) -->