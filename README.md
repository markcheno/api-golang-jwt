# API-Golang-jwt

### Bateries of POC
* Golang 1.7.3
* JWT
* Linux Ubuntu
* MongoDB / Redis
* HTTPie framework
* Glide

### Update project dependencies
Before you running the server. Use the `glide` for update all packages of project.

### DB Dispatch
We can use one or more database / cache in same time, just add the session of db on `db/dispatch.go`

### POC Running
POC of `API with jwt` more `Permission(bitwise)` middleware per route.

Generate a valid token. Post `/auth` with fields ***email*** and ***password***
```sh
http --verbose post http://localhost:3333/auth email=tzilli@inviron.com.br password=1233

POST /auth HTTP/1.1
Accept: application/json, */*
Accept-Encoding: gzip, deflate
Connection: keep-alive
Content-Length: 53
Content-Type: application/json
Host: localhost:3333
User-Agent: HTTPie/0.9.9

{
    "email": "tzilli@inviron.com.br", 
    "password": "1233"
}
```

We got the response.
```sh
HTTP/1.1 201 Created
Content-Length: 203
Content-Type: application/json; charset=utf-8
Date: Mon, 13 Feb 2017 12:22:17 GMT
Vary: Origin

{
    "expire": "2017-02-14 15:14:25 -0200 BRST", 
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoic3VwZXItaWQtb2YtbW9uZ29kYi11c2VyIiwiYWRtaW4iOnRydWUsImV4cCI6MTQ4NzA5MjQ2NSwiaXNzIjoibG9jYWxob3N0OjMzMzMifQ.lEy23l89sAe03g9Dg24FUiqUKEopSt61f-CE-1U6SpM"
}

```


Save your token on a txt file, like `token.txt` and use a `GET` on `/admin/:slug` endpoint to test token.
```sh
$ http --verbose --auth-type=jwt --auth=$(cat token.txt): get http://localhost:3333/admin
```

Request

```sh
GET /admin HTTP/1.1
Accept: */*
Accept-Encoding: gzip, deflate
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiaWQtb2YtdXNlci1uaWNlIiwiYWRtaW4iOnRydWUsImV4cCI6MTQ4NzA3NDkzNywiaXNzIjoibG9jYWxob3N0OjMzMzMifQ.6Xxg678o6WrhQULMtYA9Z7GXICsruFrXIcHPIqQy6cw
Connection: keep-alive
Host: localhost:3333
User-Agent: HTTPie/0.9.9


HTTP/1.1 200 OK
Content-Length: 41
Content-Type: text/plain; charset=utf-8
Date: Mon, 13 Feb 2017 12:35:09 GMT
Vary: Origin

protected area. USER ID = 58aadf2ce3bdea2e00f9563b
```

### Project tree
```sh
├── api
│   ├── config
│   ├── controllers
│   │   ├── admin.go
│   │   ├── auth.go
│   │   ├── permission.go
│   │   ├── project.go
│   │   └── user.go
│   ├── dbs
│   │   ├── dispatch.go
│   │   ├── logger.go
│   │   └── mongodb.go
│   ├── middlewares
│   │   ├── loggerrequest.go
│   │   ├── mongodb.go
│   │   ├── permission.go
│   │   ├── project.go
│   │   └── token.go
│   ├── models
│   │   ├── jwt.go
│   │   ├── permission.go
│   │   ├── project.go
│   │   └── user.go
│   ├── routes
│   │   ├── main.go
│   │   ├── protected.go
│   │   └── public.go
│   ├── server.go
│   ├── services
│   │   ├── jwtauth.go
│   │   └── user.go
│   └── shared
│       ├── bcrypt.go
│       ├── pathproject.go
│       └── split.go
├── bin
│   ├── api-server-arm
│   ├── api-server-x32
│   └── api-server-x64
├── docs
│   └── httpie.txt
├── glide.lock
├── glide.yaml
├── README.md
├── scripts
│   ├── build-all.sh
│   ├── build-arm.sh
│   ├── build-docker-image.sh
│   ├── build-linux-x32.sh
│   └── build-linux-x64.sh
└── token.txt
```

### MongoDB models

#### Permissons
```
{ 
    "_id" : ObjectId("58ab1f5d846dd72723990def"), 
    "permissions" : [
        {
            "method" : "ALL", 
            "value" : NumberInt(16)
        }
    ], 
    "owner" : ObjectId("58aadf2ce3bdea2e00f9563b"), 
    "endpoint" : "admin", 
    "project" : ObjectId("58ab26d3e3bdea342a9571d8"), 
    "created_at" : ISODate("2017-02-20T17:02:11.452+0000"), 
    "updated_at" : ISODate("2017-02-20T17:03:11.506+0000")
}
```
#### Projects
```
{ 
    "_id" : ObjectId("58ab26d3e3bdea342a9571d8"), 
    "label" : "Test Project", 
    "slug" : "test-project", 
    "description" : "Test Project crud", 
    "owner" : ObjectId("58aadf2ce3bdea2e00f9563b"), 
    "users" : [
        ObjectId("58aadf2ce3bdea2e00f9563b")
    ], 
    "created_at" : ISODate("2017-02-20T17:26:43.083+0000"), 
    "updated_at" : ISODate("2017-02-20T17:26:43.083+0000")
}
```
#### Users
```
{ 
    "_id" : ObjectId("58aadf2ce3bdea2e00f9563b"), 
    "email" : "tzilli@inviron.com.br", 
    "password" : "$2a$10$gMmv6mFAxcJwqT1fef6Uj.KB.A.FL5DTnn2ytrKzcgHc/GXGIxKhe", 
    "admin" : true, 
    "created_at" : ISODate("2017-02-20T12:21:00.333+0000"), 
    "updated_at" : ISODate("2017-02-20T12:21:00.333+0000")
}
```

### Routes
You can crate a new methods or distinct kind of routes, just edit the package on folder `routes`, file `main.go`.
For this project, I'm created two routes, one for public access and other need a `JWT` authentication. By default we need a route with a `slug` on the URL request, to check the users belongs to the project, and than check permission of owner. 
```
├── api
│   ├── routes
│   │   ├── main.go
│   │   ├── protected.go
│   │   └── public.go
```

### Middlewares
Early data check and others purpose.
```
├── api
│   ├── middlewares
│   │   ├── loggerrequest.go
│   │   ├── mongodb.go
│   │   ├── permission.go
│   │   ├── project.go
│   │   └── token.go
```

### Controllers
For more actions of endpoint just create a CRUD functions on `ROOT of folder` or you can create a `ditinct folder`.
```
├── api
│   ├── controllers
│   │   ├── admin.go
│   │   ├── auth.go
│   │   ├── permission.go
│   │   ├── project.go
│   │   └── user.go
```

### Models
Models, struct, inteface and wrapper.
```
├── api
│   ├── models
│   │   ├── jwt.go
│   │   ├── permission.go
│   │   ├── project.go
│   │   └── user.go
```

### DataBase Connections and loggers
In this section we concentrate a sessions for all dbs, all in wrapped by `dispatch.go`.
```
├── api
│   ├── dbs
│   │   ├── dispatch.go
│   │   ├── logger.go
│   │   └── mongodb.go
```


---

The MIT License (MIT)

Copyright (c) 2017 THIAGO ZILLI SARMENTO

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.



