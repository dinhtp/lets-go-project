# Let's GO Project

## What to do
In this service, you will develop a backend server side functions that handle the **project** and **task** API
resources.This service will contain 2 sub-services which are the Rest and gRPC service.
Rest service will handle the http requests while gRPC will handle the rpc request to the API resources.

The project and task ERD can be referred below: \
![project-task ERD](./asset/project-task.png)

## API Output
### Project API Output
#### Get a project by ID.
    - URL: [GET] {project_url}/go/project/{id}
    - Response:
        {
            "id": "string",
            "company_id": "string",
            "name": "string",
            "code": "string",
            "status": "string",
            "description": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Create a project for a specific company.
    - URL: [POST] {project_url}/go/company/{company_id}/project
    - Payload:
        {
            "company_id": "string",
            "name": "string",
            "code": "string",
            "status": "string", // allowed values: "active", "inactive"
            "description": "string",
        }
    - Response:
        {
            "id": "string",
            "company_id": "string",
            "name": "string",
            "code": "string",
            "status": "string",
            "description": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Update a project by ID.
    - URL: [PUT] {project_url}/go/project/{id}
    - Payload:
        {
            "id": "string",
            "company_id": "string",
            "name": "string",
            "code": "string",
            "status": "string", // allowed values: "active", "inactive"
            "description": "string",
        }
    - Response:
        {
            "id": "string",
            "company_id": "string",
            "name": "string",
            "code": "string",
            "status": "string",
            "description": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Delete a project by ID.
    - URL: [DELETE] {project_url}/go/project/{id}
    - Status: 200
#### list project by company id, page, limit, status and filter by "name", "code".
    - URL: [GET] {project_url}/go/company/{company_id}/projects
    - Query: ?page=0&limit=0&status=string&search_value=string&search_fields=name,code
    - Response:
        {
            "items": [
                {
                    "id": "string",
                    "company_id": "string",
                    "name": "string",
                    "code": "string",
                    "status": "string",
                    "description": "string",
                    "created_at": "string",
                    "updated_at": "string"
                },
                ...
            ]
            "max_page": 0,
            "total_count": 0,
            "page": 0,
            "limit": 0,
        }
#### Assign an employee to a project.
    - URL: [POST] {project_url}/go/project/{project_id}/assign/{employee_id}
    - Status: 200
    - Assign Condition: project must has status of "active"
#### Dismiss an employee to a project.
    - URL: [DELETE] {project_url}/go/project/{project_id}/dismiss/{employee_id}
    - Status: 200
    - Dismiss Condition: employee must not have any "active" task in project


### Task API Output
#### Get a task by ID.
    - URL: [GET] {project_url}/go/task/{id}
    - Response:
        {
            "id": "string",
            "project_id": "string",
            "name": "string",
            "status": "string",
            "description": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Create a task for a specific project.
    - URL: [POST] {project_url}/go/project/{project_id}/task
    - Payload:
        {
            "project_id": "string",
            "name": "string",
            "status": "string", allowed values: "to_do", "doing", "done"
            "description": "string",
        }
    - Response:
        {
            "id": "string",
            "project_id": "string",
            "name": "string",
            "status": "string",
            "description": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Update a task by ID.
    - URL: [PUT] {project_url}/go/task/{id}
    - Payload:
        {
            "id": "string",
            "project_id": "string",
            "name": "string",
            "status": "string", allowed values: "to_do", "doing", "done"
            "description": "string",
        }
    - Response:
        {
            "id": "string",
            "project_id": "string",
            "name": "string",
            "status": "string",
            "description": "string",
            "created_at": "string",
            "updated_at": "string"
        }
#### Delete a task by ID.
    - URL: [DELETE] {project_url}/go/task/{id}
    - Status: 200
#### list task by project id, page, limit, status and filter by "name".
    - URL: [GET] {project_url}/go/project/{project_id}/tasks
    - Query: ?page=0&limit=0&status=string&search_value=string&search_fields=name
    - Response:
        {
            "items": [
                {
                    "id": "string",
                    "project_id": "string",
                    "name": "string",
                    "status": "string",
                    "description": "string",
                    "created_at": "string",
                    "updated_at": "string"
                },
                ...
            ]
            "max_page": 0,
            "total_count": 0,
            "page": 0,
            "limit": 0,
        }
> NOTE: DO NOT commit changes directly into the master branch.