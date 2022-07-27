# goDb
Connects to MySQL db and does some CRUD ops.

## Instructions:

1. Create MySQL database `gobase` and a table `users`:
```
CREATE DATABASE gobase;
USE gobase;
CREATE TABLE users (
     id MEDIUMINT NOT NULL AUTO_INCREMENT PRIMARY KEY,
     username VARCHAR(255) NOT NULL,
);
```

2. Cli:
```
git clone https://github.com/ellieasager/goDb
cd goDb
```

3. Open file `server.go` and **modify** variable `connectionDeets` with your DB username and password.

4. Cli:
```
go mod init github.com/ellieasager/goDb
go install github.com/ellieasager/goDb
go build
go run server.go
```
