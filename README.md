**Description**

This project is a simple note-making application built with Golang. It provides functionalities for user authentication and note management.

**Tech Stack**

Backend: Golang

**Run**

1. Clone the Repository:
    ```sh
    git clone https://github.com/yourusername/note-making-app.git
    cd note-making-app
    ```
2. Build and Run the Application using Docker Compose:
    ```sh
    docker-compose up --build
    ```
3. Access the app at:
    ```
    http://localhost:8000
    ```

**API Endpoints**

- **Register**
  - URL: `POST /auth/register`
  - Request Body:
    ```json
    {
      "email": "user@example.com",
      "password": "yourSecurePassword"
    }
    ```

- **Login**
  - URL: `POST /auth/login`
  - Request Body:
    ```json
    {
      "email": "user@example.com",
      "password": "yourSecurePassword"
    }
    ```

- **Get All Notes**
  - URL: `GET /notes` (requires `Authorization: Bearer <token>`)
  - Query Parameters:
    - `status` (optional): Filter notes by status (true/false)
  - Response:
    ```json
    {
      "notes": [...]
    }
    ```

- **Get Note by ID**
  - URL: `GET /notes/:id` (requires `Authorization: Bearer <token>`)
  - Response:
    ```json
    {
      "notes": {
        "id": 1,
        "title": "Note Title",
        "status": true
      }
    }
    ```

- **Create Note**
  - URL: `POST /notes` (requires `Authorization: Bearer <token>`)
  - Request Body:
    ```json
    {
      "title": "Note Title",
      "status": true
    }
    ```

- **Update Note**
  - URL: `PUT /notes` (requires `Authorization: Bearer <token>`)
  - Request Body:
    ```json
    {
      "id": 1,
      "title": "Updated Title",
      "status": false
    }
    ```

- **Delete Note**
  - URL: `DELETE /notes/:id` (requires `Authorization: Bearer <token>`)
  - Response:
    ```json
    {
      "notes": "Successfully Deleted"
    }
    ```

**Environment Variables**

Configured in `docker-compose.yaml`:
- `POSTGRES_USER`
- `POSTGRES_PASSWORD`
- `POSTGRES_DB`