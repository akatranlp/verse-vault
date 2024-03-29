### Service for user account creation and authentication

Create User Account: This action allows new users to register and establish their presence on the platform by providing necessary information and credentials.

Log In: Logging in enables users to access their accounts and enjoy the platform's features and personalized content.

Log Out: This action allows users to securely exit their accounts, ensuring their privacy and data security.

## How to use the User-Service

- First you need to create a rsa-keypair

```bash
mkdir ./certs
cd ./certs
openssl genrsa -out access-key.pem 2048
openssl rsa -in access-key.pem -outform PEM -pubout -out access-public.pem
openssl genrsa -out refresh-key.pem 2048
openssl rsa -in refresh-key.pem -outform PEM -pubout -out refresh-public.pem
```

- Then the user-service needs the database to be running before you can start it.
- Then the following environment variables must be correctly set to point to the other services

```bash
JWT_ACCESS_PRIVATE_KEY_PATH=<path-to-rsa-key>
JWT_ACCESS_PUBLIC_KEY_PATH=<path-to-rsa-pub-key>
JWT_REFRESH_PRIVATE_KEY_PATH=<path-to-rsa-key>
JWT_REFRESH_PUBLIC_KEY_PATH=<path-to-rsa-pub-key>
JWT_ACCESS_TOKEN_EXPIRATION=15m
JWT_REFRESH_TOKEN_EXPIRATION=168h
AUTH_IS_ACTIVE=true
PORT=<HTTP-Port>
GRPC_PORT=<GRPC-Port>
POSTGRES_HOST=<db-hostname>
POSTGRES_PORT=<db-port>
POSTGRES_USER=<db-username>
POSTGRES_PASSWORD=<db-passwort>
POSTGRES_DB=<db-name>
```

- At last the service can be run with the following command: `go run main.go`

### Create Docker-Image

If you want to use an docker-image instead, the following commands must be executed from the root of this project:

```bash
docker build -t user-service -f ./src/user-service/Dockerfile .
docker run -dit -v <certs-folder>:<certs-folder> -p port:port -p grpcPort:grpcPort --envFile envFile user-service
```
