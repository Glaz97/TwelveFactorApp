# :gear: TwelveFactorApp

TwelveFactorApp implements simple REST API server.

## Getting Started

### Dependencies

- MongoDB
- Docker
- Make

### Server

TwelveFactorApp exposes REST API.

- `localhost:8056` REST server

### Local Development

Running MongoDB and Docker is necessary for the application.. To start MongoDB run:

```shell
make mongo
```

To run `twelvefactorapp`:

```shell
go run cmd/twelvefactorapp/main.go
```

### Endpoints :

- to create article execute: `POST localhost:8056/article` 
###### with json body: { "title": "someTestName" }

- to get article execute: `GET localhost:8056/article/{articleId}`
###### where {articleId} - id field, of a response object of a POST call

- swagger: `GET localhost:8056/docs/index.html` 

- status `GET http://localhost:8056/status`
