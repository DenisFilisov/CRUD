# CRUD

**About:**
The main idea is to create a service with authorization by JWT token.
Also in this service we have News and Followers entities.
Users can create, read, update and delete the news. Also, users can subscribe to news and unsubscribe from news.


**What is contained:**
CRUD logic
AUTH by username/password
JWT token for all endpoints
Connection with DB
Graceful shutdown
Swagger for all endpoints
Migration tables and rollback tables
Docker Compose file
Make file

**How to run this application:**
1. Start docker
2. make run
3. make migrate 
4. go run cmd/main.go

