# Authfer

Authfer is a Golang authentication service.

## Installation

```bash
go get -u github.com/Ifkarsyah/authfer
```

## Usage

### Login
```curl
curl --location --request POST 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
	"username": "username",
	"password": "password"
}'
```

### Refresh Token
```curl
curl --location --request POST 'http://localhost:8080/refresh' \
--header 'Authorization: Bearer {access token}' \
--header 'Content-Type: application/json' \
--data-raw '{
	"refresh_token": {refresh token}
}'
```

### Logout
```curl
curl --location --request POST 'http://localhost:8080/logout' \
--header 'Authorization: Bearer {access token}'
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
