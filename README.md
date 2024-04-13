
## Run locally with Docker

### Build
```bash
  docker-compose up --build
```
App will be listening at `localhost:8080`

### Reset DB
```bash
docker-compose down --volumes
```
## Usage / Testing

### Get Access Token
```cURL
curl --location 'https://dev-waygvmnclik5ibp8.us.auth0.com/oauth/token' \
--header 'content-type: application/x-www-form-urlencoded' \
--data-urlencode 'grant_type=password' \
--data-urlencode 'username=user1' \
--data-urlencode 'password=Password123!' \
--data-urlencode 'audience=https://loyalid-test/get-token' \
--data-urlencode 'scope=read:sample' \
--data-urlencode 'client_id=VPmfCPIyGcPKUfWRMkZ5Un2pZxKnpxQ9' \
--data-urlencode 'client_secret=9kHz1P7_ESm7QYctzye8R4xNSE0SQObUZbefuTNwql-eH8_K_py7WUjQAVBRVxy1'
```
Note the token to be used at `access_token` field from the response.

### Current User API
```cURL
curl --location 'http://localhost:8080/user/current' \
--header 'Authorization: Bearer {TOKEN}' \
```
Replace {TOKEN} with obtained access token from `Get Access Token`.

### Create Product API
```cURL
curl --location 'http://localhost:8080/product' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer {TOKEN}' \
--data '{
    "name": "A",
    "price": 10
}'
```
Replace {TOKEN} with obtained access token from `Get Access Token`.

### List Products API
```cURL
curl --location 'http://localhost:8080/product?page=1&limit=10' \
--header 'Authorization: Bearer {TOKEN}'
```
Replace {TOKEN} with obtained access token from `Get Access Token`.