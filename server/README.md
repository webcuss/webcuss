# webcuss | server

## Stack & Libs
1. Web framework - [Gin](https://github.com/gin-gonic/gin)
2. Postgres driver - [Pgx](https://github.com/jackc/pgx)
3. Postgres' cryptography extension - [pgcrypto](https://www.meetspaceapp.com/2016/04/12/passwords-postgresql-pgcrypto.html)

## Requirements
1. GoLang 1.19+

## Migrate
Create db tables
```shell
export DATABASE_URL=postgres://postgres:postgres@localhost:5432/mydb

go run main.go migrate
# go run main.go migrate clear # delete all tables & indexes
```

## Run
```shell
export APP_SECRET=*************
export DATABASE_URL=postgres://postgres:postgres@localhost:5432/mydb

go get ./...
go run main.go
```

## Test
```shell
export APP_SECRET=*************
export DATABASE_URL=postgres://postgres:postgres@localhost:5432/mydb

go test
```

## Database
### Table `avatar`
| column     | type           | index | fk  |
|------------|----------------|-------|-----|
| id         | string         | y     |     |
| uname      | string         | y     |     |
| pword      | string/hash    |       |     |
| createdOn  | timestamp      | y     |     |
| pebbles    | int            |       |     |
| verifiedOn | timestamp/null |       |     |
| email      | string/null    |       |     |

### Table `topic`
| column      | type      | index | fk      |
|-------------|-----------|-------|---------|
| id          | string    | y     |         |
| scheme      | string    |       |         |
| hostname    | string    | y     |         |
| path        | string    | y     |         |
| query       | string    | y     |         |
| querySearch | string    | y     |         |
| title       | string    |       |         |
| likes       | int       | y     |         |
| createdOn   | timestamp | y     |         |
| userId      | string    | y     | user.id |

### Table `comment`
| column    | type        | index | fk         |
|-----------|-------------|-------|------------|
| id        | string      | y     |            |
| topicId   | string/null | y     | topic.id   |
| commentId | string/null | y     | comment.id |
| userId    | string      | y     | user.id    |
| content   | string      |       |            |
| createdOn | timestamp   | y     |            |

### Table `reaction`
| column    | type        | index | fk         | desc       |
|-----------|-------------|-------|------------|------------|
| id        | string      | y     |            |            |
| userId    | string      | y     | user.id    |            |
| commentId | string      | y     | comment.id |            |
| reaction  | int         |       |            | `0`=👍     |
| createdOn | timestamp   | y     |            |            |
**Unique constraint**: `[userId, commentId, reaction]`

## API
* `/sup` - Sign up
    * POST
        * body
            ```json
            {
                "uname": "string",
                "pword": "string"
            }
            ```
        * returns 201 Created
            ```json
            {
                "token": "string"
            }
            ```
* `/sin` - Sign in
    * POST
        * body
            ```json
            {
                "uname": "string",
                "pword": "string"
            }
            ```
        * return 200 Ok
            ```json
            {
                "token": "string"
            }
            ```
* `/tpc` - Topic
    * GET - get top topics
        * qs
            * `&pg` - page number
        * returns 200 Ok
            ```json
            {
                "data": [
                    {
                        "id": "string",
                        "hostname": "string",
                        "path": "string",
                        "query": "string",
                        "title": "string",
                        "commentsCount": "number",
                        "likes": "number",
                        "user": {
                            "id": "string",
                            "uname": "string"
                        }
                    }
                ]
            }
            ```
    * POST - create a topic
        * body
            ```json
            {
                "url": "string",
                "title": "string",
                "comment": "string|null"
            }
            ```
        * returns 201 Created
            ```json
            {
                "id": "string",
                "commentId": "string|null"
            }
            ```
* `tpc/{topicId}/cmt`
    * POST - add comment to a topic
        * body
            ```json
            {
                "comment": "string"
            }
            ```
        * returns 201 created
            ```json
            {
                "id": "string"
            }
            ```
    * GET - get the comments of the topic
        * qs
            * `&pg` - page number
        * returns 200 Ok
            ```json
            {
                "data": [
                    {
                        "id": "string",
                        "content": "string",
                        "createdOn": "string",
                        "user": {
                            "id": "string",
                            "uname": "string"
                        }
                    }
                ]
            }
            ```
* `/cmt/{commentId}`
    * POST - reply to a comment
        * body
            ```json
            {
                "comment": "string"
            }
            ```
        * returns 201 Created
            ```json
            {
                "id": "string"
            }
            ```
    * GET - get the replies of the comment
        * qs
            * `&pg` - page number
        * returns 200 Ok
            ```json
            {
                "data": [
                    {
                        "id": "string",
                        "content": "string",
                        "createdOn": "string",
                        "user": {
                            "id": "string",
                            "uname": "string"
                        }
                    }
                ]
            }
            ```
* `/rctn/{commentId}`
    * POST - add reaction to a comment
        * body
            ```json
            {
                "reaction": "number"
            }
            ```
        * returns 201 Created
            ```json
            {
                "id": "string"
            }
            ```
    * GET - get the reactions of a comment
        * returns 200 Ok
            ```json
            {
                "data": [
                    {
                        "reaction": "number",
                        "count": "number"
                    }
                ]
            }
            ```
