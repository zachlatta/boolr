# Boolr

_Finally_, a web-scale boolean-switching service. Informally created for
[SOHacks](http://sohacks.com/).

## Endpoints

### POST /users

#### Request

```json
{
  "username": "zaphod",
  "password": "Betelgeuse123"
}
```

#### Response

```
HTTP/1.1 200 OK
Content-Length: 114
Content-Type: application/json
Date: Sat, 14 Jun 2014 19:13:29 GMT
```

```json
{
  "created": "2014-06-14T19:13:29.747225073Z",
  "id": 2, "updated": "2014-06-14T19:13:29.74722522Z",
  "username": "zaphod"
}
```

###  POST /users/login

Get a JWT token issued. Include the issued token in future requests in the
`Authorization` header to authenticate.

Example `Authorization` header:

    Authorization:Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0MDMwMzI0ODUsImlkIjoyfQ.pzIjSmZcyfoGy9YWd0q3eiM_TEPQ5_TqTZjr2LYR8U0

#### Request

```json
{
  "username": "zaphod",
  "password": "Betelgeuse123"
}
```

#### Response

```
HTTP/1.1 200 OK
Content-Length: 171
Content-Type: application/json
Date: Sat, 14 Jun 2014 19:14:45 GMT
```

```json
{
  "expires": "2014-06-17T19:14:45.074069349Z",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE0MDMwMzI0ODUsImlkIjoyfQ.pzIjSmZcyfoGy9YWd0q3eiM_TEPQ5_TqTZjr2LYR8U0"
}
```

### GET /users/:id

* Must be authenticated
* `me` can be used as the id to retrieve the current user

#### Response

```
HTTP/1.1 200 OK
Content-Length: 109
Content-Type: application/json
Date: Sat, 14 Jun 2014 19:18:28 GMT
```

```json
{
  "created": "2014-06-14T19:13:29.747225Z",
  "id": 2,
  "updated": "2014-06-14T19:13:29.747225Z",
  "username": "zaphod"
}
```

### POST /users/:id/booleans

* Must be authenticated
* `me` can be used as the id to retrieve the current user

#### Request

```json
{
  "label": "Special Boolean"
}
```

* `label` is optional
* A boolean will be created if the request body is empty

#### Response

```
HTTP/1.1 200 OK
Content-Length: 162
Content-Type: application/json
Date: Sat, 14 Jun 2014 19:25:32 GMT
```

```json
{
  "bool": false,
  "created": "2014-06-14T19:25:32.380204314Z",
  "id": 7,
  "label": "Special Boolean",
  "switch_count": 0,
  "updated": "2014-06-14T19:25:32.38020436Z",
  "user_id": 2
}
```

### GET /users/:id/booleans

* Must be authenticated
* `me` can be used as the id to retrieve the current user

#### Response

```
HTTP/1.1 200 OK
Content-Length: 289
Content-Type: application/json
Date: Sat, 14 Jun 2014 19:32:46 GMT
```

```json
[
  {
    "bool": false,
    "created": "2014-06-14T19:25:22.198659Z",
    "id": 6,
    "switch_count": 0,
    "updated": "2014-06-14T19:25:22.198659Z",
    "user_id": 2
  },
  {
    "bool": true,
    "created": "2014-06-14T19:25:32.380204Z",
    "id": 7,
    "label": "Special Boolean",
    "switch_count": 1,
    "updated": "2014-06-14T19:26:51.588526Z",
    "user_id": 2
  }
]
```

### GET /booleans/:id

* Must be authenticated

#### Response

```
HTTP/1.1 200 OK
Content-Length: 156
Content-Type: application/json
Date: Sat, 14 Jun 2014 19:28:44 GMT
```

```json
{
  "bool": false,
  "created": "2014-06-14T19:25:32.380204Z",
  "id": 7,
  "label": "Special Boolean",
  "switch_count": 1,
  "updated": "2014-06-14T19:25:32.380204Z",
  "user_id": 2
}
```

### PUT /booleans/:id/switch

* Must be authenticated

#### Request

* Empty request body is fine, takes no parameters

#### Response

```
HTTP/1.1 200 OK
Content-Length: 159
Content-Type: application/json
Date: Sat, 14 Jun 2014 19:26:51 GMT
```

```json
{
  "bool": true,
  "created": "2014-06-14T19:25:32.380204Z",
  "id": 7,
  "label": "Special Boolean",
  "switch_count": 1,
  "updated": "2014-06-14T19:26:51.588526076Z",
  "user_id": 2
}
```

## License

See the [LICENSE](LICENSE) file.
