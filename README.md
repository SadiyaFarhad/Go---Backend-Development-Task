# Go Backend Development Task  
Build a RESTful API using Go to manage users, including their name and date of birth. The API should calculate and return each user’s age dynamically when retrieving user details.  
This lightweight RESTful backend is built with Go (Fiber) and PostgreSQL, and it's fully containerized with Docker Compose. It provides CRUD operations for user data and automatically calculates users' ages from their dates of birth.  
  
**Overview**  
  
Architecture:  
  
Go Fiber API, handles routes, validation, and responses.  
  
PostgreSQL, stores user data (id, name, dob).  
  
Computation, calculates user age dynamically.  
  
Docker Compose, runs and connects app and database containers.  
  
**Tech Stack**  
Component | Purpose  
--- | ---  
Go (v1.25) | Core backend language  
Fiber (v2) | Web framework  
PostgreSQL (v16) | Database  
Docker / Compose | Containerization  
Zap (Uber) | Logging  
Validator v10 | Input validation  
  
⚙️ Setup & Commands  
# Initialize Go module  
go mod init go-backend-user  

# Install dependencies  
go get <package>  
go mod tidy  

# Build and start containers  
docker-compose build  
docker-compose up  

# Stop and clean  
docker-compose down  
docker system prune -f  
  
🗃 Database Commands  
# Access PostgreSQL shell  
docker exec -it <db_container_name> psql -U user -d users_db  

\dt                         # List tables  
SELECT * FROM users;        # View users  
INSERT INTO users (name, dob) VALUES ('Alice', '1990-05-10');  
UPDATE users SET name='Alice Updated' WHERE id=1;  
DELETE FROM users WHERE id=1;  
\q                          # Exit  
  
**API Endpoints**  
Method | Endpoint | Description  
--- | --- | ---  
POST | /users | Create new user  
GET | /users | Get all users  
GET | /users/:id | Get user by ID  
PUT | /users/:id | Update user  
DELETE | /users/:id | Delete user  
  
**Execution Flow**  
  
docker-compose up starts the app and database containers.  
  
The Go app connects to PostgreSQL using environment variables.  
  
The API listens on localhost:3000.  
  
Logs are managed with Zap.  
  
All CRUD operations store data in PostgreSQL.  
  
**Outcome**  
  
This creates a functional backend service for user management.  
  
The modular REST API has a clear separation of layers.  
  
It is portable and easy to deploy with Docker.  
  
Structured logging and input validation are implemented.
