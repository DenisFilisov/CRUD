# CRUD

**About:**
The main idea is to create a service with authorization by JWT token.
Also in this service we have News and Followers entities.
Users can create, read, update and delete the news. Also, users can subscribe to news and unsubscribe from news.
All logs processed throw the Queue. (Kafka)


**What is contained:**
1. CRUD logic
2. AUTH by username/password
3. JWT token for all endpoints
4. Connection with DB
5. Graceful shutdown
6. Swagger for all endpoints
7. Migration tables and rollback tables
8. Docker Compose file
9. Make file
10. Cache Lib added
11. JWT tokens cashed
12. RefreshToken Realisation
13. Kafka implemented
14. Log Handler implemented

**How to run this application:**
1. Start docker
2. make run
3. make migrate 
4. go run cmd/main.go


For start application you should add to .env file some variables
```dotenv
DB_PASSWORD=postgres

SALT=dngkahfkglahlfanhfla
SIGNINGKEY=fpahfolkahoghalokghoa
```