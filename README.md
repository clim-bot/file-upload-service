# File Upload Service

This project is a file upload service built with the Gin framework in Go. It allows users to upload, download, and manage files. The service uses GORM for ORM, SQLite3 for the database, JWT for authentication, and includes file type validation and unit tests.

## Features

- User authentication with JWT
- File upload with metadata storage
- File download
- List all files
- Delete files
- File type validation
- Unit tests

## Tech Stack

- [Go](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [GORM](https://gorm.io/)
- [SQLite3](https://www.sqlite.org/index.html)
- [godotenv](https://github.com/joho/godotenv)
- [JWT](https://github.com/golang-jwt/jwt/v5)
- [bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt)

## Folder Structure

```go
file-upload-service/
├── config/
│ └── config.go
├── controllers/
│ └── file_controller.go
├── database/
│ └── database.go
├── middleware/
│ └── auth.go
├── models/
│ └── file.go
├── routes/
│ └── routes.go
├── services/
│ └── file_service.go
├── utils/
│ └── validator.go
├── main.go
├── go.mod
├── go.sum
├── .env
```


## Setup

1. **Clone the repository:**

   ```bash
   git clone https://github.com/yourusername/file-upload-service.git
   cd file-upload-service
   ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Create a .env file:
    ```bash
    SECRET_KEY=your_secret_key
    ```

4. Run the application:
    ```bash
    go run main.go
    ```

## API Endpoints
- `POST /upload`: Upload a file (protected route).
- `GET /download/:id`: Download a file by ID (protected route).
- `GET /files`: List all files (protected route).
- `DELETE /files/:id`: Delete a file by ID (protected route).


## Example Requests
### Upload File
```bash
curl -X POST "http://localhost:8080/upload" -H "Authorization: Bearer <token>" -F "file=@path/to/your/file.jpg"
``` 

### Download File
```bash
curl -X GET "http://localhost:8080/download/1" -H "Authorization: Bearer <token>" --output downloaded_file.jpg
```

### List Files
```bash
curl -X GET "http://localhost:8080/files" -H "Authorization: Bearer <token>"
```

### Delete File
```bash
curl -X DELETE "http://localhost:8080/files/1" -H "Authorization: Bearer <token>"
```

## Running Tests
To run the unit tests:
```bash
go test ./...
```

## License
This project is licensed under the MIT License.