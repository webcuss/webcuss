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
* `/sout` - Sign out
    * returns 200 Ok