# jwt-test-server
Simple JWT Server made to understand how JWT works.  
By https://github.com/mmalessa  

Inspired by https://github.com/sohamkamani/jwt-go-example

**`In a production environment, use it only at your own risk!`**  

# Installation
```sh
go get github.com/mmalessa/jwt-test-server
go install github.com/mmalessa/jwt-test-server
```

# Run
```sh
~/go/bin/jwt-test-server /path/to/config/file.yaml
```

# Configuration
All you need is in config.yaml
```yaml
jwt:
    key: mySecretKey
    expirationtime: 30 # in minutes
server:
    port: 8000
credentials:
    user1: password1
    user2: password2
```

# API Methods
## POST http://localhost:8000/login

### Request Body
```json
{
    "username":"myUsername",
    "password":"myHiddenPassword"
}
```

### Response
200: application/json
```json
{
    "token": "eyJhbGciOiJ*********Tw"
}
```

## GET http://localhost:8000/refresh

### Header
`Authorization: "Bearer eyJhbGciOiJ*********Tw"`  

### Response
200: application/json
```json
{
    "token": "eyJhbGciOiJ*********Tw"
}
```

## GET http://localhost:8000/welcome

### Header
`Authorization: "Bearer >JWT Token string<`  

### Response
200: application/json
```json
{
    "username": "test",
    "exp": 1561991616
}
```

## Error Responses
400, 401, 500: application/json  
```json
{
    "Code": "500",
    "Message": "Internal Server Error"
}
```

