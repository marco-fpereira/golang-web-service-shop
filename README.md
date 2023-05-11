# About
This application is a MVP of a web service made using the following technologies:
    Backend: Golang
    Container: Docker
    Database: Postgresql (inside docker container)
    Frontend: HTML and Bootstrap

# Create database in docker

## Build docker image
```
cd docker
docker build -t postgres-web-shop-image ./
```

## Execute image above created
### If you want volumes, include: -v /tmp/database:/var/lib/postgresql/data 
```
docker run -e POSTGRES_PASSWORD=1234 --name web_shop -p 5432:5432 -d postgres-web-shop-image
```

# Verify if docker image was correctly created

### Enter into docker container
```
docker exec -it web_shop bash
```

### Checkout to postgres user
```
su - postgres
```

### Enter into sql table to execute queries
```
psql -d web_shop_db -U postgres -W
password: 1234
```

### list all databases
```
\l
```

### list tables from selected database
```
\dt
```

### show columns of table
```
\d products
```

### show previous inserted products
```
SELECT * FROM products;
```

### quit table
```
\q
```

### exit docker container
```
exit
exit
```

# Running application
## Download lib that helps communication between golang and postgresql:
### Using this doc: https://pkg.go.dev/github.com/lib/pq
```
cd ..
cd app
go mod init
go get -u -d github.com/lib/pq
```
## Running go
```
go run main.go
```