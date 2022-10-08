# webcuss

## Run Locally
Requirements:
* GoLang 1.19+


Run:
## Development
```sh
go get ./...
go run main.go
```

## Database
### Table `avatar`
| column                | type                  | index     | fk            |
| ---                   | ---                   | ---       | ---           |
| id                    | string                | y         |               |
| uname                 | string                | y         |               |
| pword                 | string/hash           |           |               |
| createdOn             | timestamp             | y         |               |
| pebbles               | int                   |           |               |
| verifiedOn            | timestamp/null        |           |               |
| email                 | string/null           |           |               |

### Table `topic`
| column                | type                  | index     | fk            |
| ---                   | ---                   | ---       | ---           |
| id                    | string                | y         |               |
| url                   | string                | y         |               |
| search                | string                | y         |               |
| createdOn             | timestamp             | y         |               |
| userId                | string                | y         | user.id       |

### Table `comment`
| column                | type                  | index     | fk            |
| ---                   | ---                   | ---       | ---           |
| id                    | string                | y         |               |
| topicId               | string/null           | y         | topic.id      |
| commentId             | string/null           | y         | comment.id    |
| userId                | string                | y         | user.id       |
| content               | string                |           |               |
| createdOn             | timestamp             | y         |               |

## API Routes
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
                "pg": "number",
                "data": [
                    {
                        "id": "string",
                        "url": "string",
                        "title": "string",
                        "commentsCount": "number",
                        "likesCount": "number",
                        "creator": "string",
                        "createdOn": "string"
                    }
                ]
            }
            ```
    * POST - create a topic
        * body
            ```json
            {
                "url": "string",
                "comment": "string|null"
            }
            ```
        * returns 201 Created
* `tpc/{topicId}/cmt`
    * POST - add comment to a topic
        * body
            ```json
            {
                "comment": "string"
            }
            ```
    * GET - get the comments of the topic
        * qs
            * `&pg` - page number
        * returns 200 Ok
            ```json
            {
                "id": "string",
                "url": "string",
                "pg": "number",
                "data": [
                    {
                        "id": "string",
                        "comment": "string",
                        "user": {
                            "id": "string",
                            "name": "string"
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
                "message": "string"
            }
            ```
        * returns 201 Created
    * GET - get the replies of the comment
        * qs
            * `&pg` - page number
        * returns 200 Ok
            ```json
            {
                "pg": "number",
                "data": [
                    {
                        "comment": "string",
                        "user": { }
                    }
                ]
            }
            ```
## Reference
* [pgcrypto](https://www.meetspaceapp.com/2016/04/12/passwords-postgresql-pgcrypto.html)
* 