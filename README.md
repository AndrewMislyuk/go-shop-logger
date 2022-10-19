# Logging service for https://github.com/AndrewMislyuk/go-shop-backend

### Stack
- go 1.19
- mongo 

### Setup
Create .env file with variables for application

```
export DB_URI="mongodb://mongo:27017"
export DB_USERNAME="admin"
export DB_PASSWORD="admin"
export DB_DATABASE="audit"

export SERVER_PORT=9000
```

### Build and Run
```
make build_up
```
