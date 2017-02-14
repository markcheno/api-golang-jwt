# API-Golang-jwt

#### Bateries of POC
* Golang 1.7.3
* JWT
* Linux Ubuntu
* HTTPie framework

Generate a valid token. Post `/auth` with fields ***email*** and ***password***
```sh
http --verbose post http://localhost:3333/auth email=tzilli@inviron.com.br password=123

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
    "password": "123"
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


Save your token on a txt file, like `token.txt` and use a `GET` on `/admin` endpoint to test token.
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

protected area. USER ID = id-of-user-nice
```



