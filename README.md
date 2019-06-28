# jwt-test-server
Simple JWT Server made to understand how JWT works.  
Inspired by https://github.com/sohamkamani/jwt-go-example

**`In a production environment, use it only at your own risk!`**  

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
200: (string) ">JWT Token string<"  

## GET http://localhost:8000/refresh

### Header
`Authorization: "Bearer >JWT Token string<"`  

### Response
`200: (string) JWT Token string`  

## GET http://localhost:8000/welcome

### Header
`Authorization: "Bearer >JWT Token string<?`  

### Response
`200: (string) "Welcome {username}"`  

## Error Responses
`400: (json) {"Code":"400","Message":"Bad Request"}`  
`401: (json) {"Code":"401","Message":"Unauthorized"}`  
`500: (json) {"Code":"500","Message":"Internal Server Error"}`  
