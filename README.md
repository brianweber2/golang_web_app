# golang_web_app
Web app built in Golang using the standard http library.

## Local Development

Run the web server with the command:

`go run cmd/web/*.go`

Goto `http://localhost:8080`

### Helpful Commands

`go get <path to package>`: to install new go dependencies

`go mod tidy`: to update the dependency lock file after adding imports in the project

## Project Layout

### cmd

Under `cmd/web` includes the following:

* `main.go` - entry point of the web app
* `middleware.go` - place for code that interacts with a request betweem the client and server
* `routes.go` - place where the router and routes are configured

### pkg

Under `pkg/config` includes the following:

* `config.go` - stores the application-wide configuration

Under `pkg/handlers` includes the following:

* `handlers.go` - defines the repository connection and the logic for each route

Under `pkg/models` includes the following:

* `templatedata.go` - defines the schema for data that gets injected into the HTML templates

Under `pkg/render` includes the following:

* `render.go` - defines how HTML templates are rendered by reading the filesystem or template cache

### templates

* `layout.html` - defines the content shared across the application
* `home.html` - defines the content a user will see on the application's home page
* `about.html` - defines the content a user will see on the application's about page

## Endpoints Available

* `/` - home page
* `/about` - about page
