# User Service

A service for managing user data (creating new user, get list user and get specific user)

## Prerequisite Requirements

Before running this project, ensure you have the following tools installed:

1. **Golang** (version 1.22.10 or higher) - for building the application.
2. **Git** - for version control.

---

## Project Setup and Usage

### 1. Clone the Repository

```bash
git clone https://github.com/okaaryanata/user-svc.git
cd user-svc
```

### 1. Install Dependencies

```bash
git mod tidy
```

### 2. Create File .env

- create file **.env** base on file **.env.example**

```
go run cmd/user/main.go
```

## API Endpoints

### Service Route

**`{{url}}/users`**

### User Endpoints

| Method | Endpoint                        | Description                                                                        |
| ------ | ------------------------------- | ---------------------------------------------------------------------------------- |
| `POST` | `/users`                        | Create a user                                                                      |
| `GET`  | `/users/{id}`                   | Get specific user by ID                                                            |
| `GET`  | `/users?page_num=1&page_size=1` | Get list user, optional: query param **page_num** and **page_size** for pagination |

## [Postman Document Link](https://documenter.getpostman.com/view/7748154/2sAYX3r3cs)

## Environment Variables

The application uses the following environment variables (defined in the `.env` file):

```plaintext
APP_HOST=localhost
APP_PORT=9090
DB_FILE=users.db
DB_MIGRATION=true
```

---

## License

This project is licensed under the MIT License.
