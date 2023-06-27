## Backoffice

DarchLabs backoffice api

### API spec


#### **GET /api/v1/health**

**Response**

Status: 200

```json
{
	"status": "running"
}
```

#### **POST /api/v1/users/signup**

**Request**

```json
{
	"email": "jdoe@gmail.com",
	"name": "jon doe",
	"password": "password124"
}
```

**Response**

Status: 201

```json
null
```

#### **POST /api/v1/users/login**

**Request**

```json
{
	"email": "jdoe@gmail.com",
	"password": "password124"
}
```

**Response**

Status: 201

```json
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Impkb2VAZ21haWwuY29tIiwiZXhwIjoxNzE5MTU3MjczfQ.MfaT_5lrapX4sapI992uQodW0xHsbv4UeNf0guCUEaA"
}
```

#### **POST /api/v1/users/api-key**

**Request**

```json
{
	"days_interval": 30
}
```

**Response**

Status: 201

```json
{
	"api_key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImY1ZTNiNTQyLWMwY2ItNDJmMi1hMjJlLWNlOWNkMGQzMTk3YyIsImV4cCI6MTY4ODExMjY3OX0.VXjFWxHXlW_TlkZ4HN_n0PmGqiaC9-O38LNWKHN1e2A"
}
```

#### **POST /api/v1/users/tokens**

**Request**

```json
{
	"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6Impkb2VAZ21haWwuY29tIiwiZXhwIjoxNzE3MDU1MDIwfQ.1ifMxIZ1wnNAjU2i-Kx7iEALVoIFaUrAQfWLAiLFnds"
}
```

**Response**

Status: 200

```json
{
	"user_id": "f5e3b542-c0cb-42f2-a22e-ce9cd0d3197c"
}
```

#### **POST /api/v1/users/valid-api-key**

**Request**

```json
{
	"api_key": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImY1ZTNiNTQyLWMwY2ItNDJmMi1hMjJlLWNlOWNkMGQzMTk3YyIsImV4cCI6MTY4ODExMjY3OX0.VXjFWxHXlW_TlkZ4HN_n0PmGqiaC9-O38LNWKHN1e2A"
}
```

**Response**

Status: 200

```json
{
	"user_id": "f5e3b542-c0cb-42f2-a22e-ce9cd0d3197c"
}
```