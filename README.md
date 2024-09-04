# Todo Web App

A comprehensive Todo web application built using Go, MongoDB, and JWT for authentication and role-based authorization. This application provides CRUD operations for todos, user registration, login, and role-based access control.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete todos.
- **User Authentication**: Sign up and log in users with hashed passwords.
- **Role-Based Authorization**: Protect routes based on user roles (e.g., user, admin).
- **JWT Token Management**: Secure endpoints with JWT tokens.
- **Rate Limiting**: Protect the API from abuse.
- **CORS Handling**: Enable cross-origin requests.

## Requirements

- Go 1.16+
- MongoDB
- Git

## Getting Started

### 1. Clone the Repository

```bash
git clone https://github.com/chmuhammadasim/Todo-WebApp-APIs-Golang.git
cd todo-webapp
```

### 2. Set Up MongoDB

Make sure you have MongoDB running locally or on a remote server. Update the MongoDB connection string in the `db` package if necessary.

### 3. Install Dependencies

Navigate to your project directory and install Go dependencies:

```bash
go mod tidy
```

### 4. Configure Environment Variables

Create a `.env` file in the root of the project with the following variables:

```
MONGO_URI=mongodb://localhost:27017/todoapp
JWT_SECRET=your_secret_key
```

Replace `your_secret_key` with a strong secret key for signing JWT tokens.

### 5. Run the Application

To start the application, use:

```bash
go run main.go
```

The server will start on port `8080` by default.

## API Endpoints

### User Authentication

- **Sign Up**
  - **Endpoint**: `/signup`
  - **Method**: `POST`
  - **Request Body**:
    ```json
    {
      "username": "your_username",
      "password": "your_password",
      "role": "user" // or "admin"
    }
    ```
  - **Response**:
    ```json
    {
      "message": "User created successfully"
    }
    ```

- **Login**
  - **Endpoint**: `/login`
  - **Method**: `POST`
  - **Request Body**:
    ```json
    {
      "username": "your_username",
      "password": "your_password"
    }
    ```
  - **Response**:
    ```json
    {
      "token": "your_jwt_token"
    }
    ```

### Todo Operations

Authentication is required for all todo operations. Use the `Authorization` header with the value `Bearer your_jwt_token`.

- **Create Todo**
  - **Endpoint**: `/api/todos`
  - **Method**: `POST`
  - **Request Body**:
    ```json
    {
      "title": "Todo Title",
      "completed": false
    }
    ```
  - **Response**:
    ```json
    {
      "id": "todo_id",
      "title": "Todo Title",
      "completed": false,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
    ```

- **Get All Todos**
  - **Endpoint**: `/api/todos`
  - **Method**: `GET`
  - **Response**:
    ```json
    [
      {
        "id": "todo_id",
        "title": "Todo Title",
        "completed": false,
        "created_at": "2024-01-01T00:00:00Z",
        "updated_at": "2024-01-01T00:00:00Z"
      }
    ]
    ```

- **Get Todo By ID**
  - **Endpoint**: `/api/todos/{id}`
  - **Method**: `GET`
  - **Response**:
    ```json
    {
      "id": "todo_id",
      "title": "Todo Title",
      "completed": false,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
    ```

- **Update Todo**
  - **Endpoint**: `/api/todos/{id}`
  - **Method**: `PUT`
  - **Request Body**:
    ```json
    {
      "title": "Updated Title",
      "completed": true
    }
    ```
  - **Response**:
    ```json
    {
      "id": "todo_id",
      "title": "Updated Title",
      "completed": true,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
    ```

- **Delete Todo**
  - **Endpoint**: `/api/todos/{id}`
  - **Method**: `DELETE`
  - **Response**:
    ```json
    {
      "message": "Todo deleted successfully"
    }
    ```

## Middleware

### Rate Limiter

Limits the number of requests from a single IP to prevent abuse.

### CORS

Handles cross-origin requests to ensure the API is accessible from different domains.

### Authentication

Protects routes by verifying JWT tokens and extracting user information.

## Contributing

Feel free to open issues or submit pull requests to contribute to this project. Ensure you follow best practices for Go and adhere to the project's coding style.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
