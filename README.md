# HTMX Test

Dummy project for testing out `HTMX` with a Golang backend. I'm using `Go Fiber` for creating the REST backend, and `Templ` for templatizing the generated HTML returned from the backend. `Air` is used for live reloading and generation of the templates.

## Running the service

### Using Air

> Air needs to be installed separately, using `go install github.com/cosmtrek/air@latest`

```Bash
./air
```

### Running service manually

```Bash
# compile and build binary
go build -o ./tmp/main ./...
# run
./tmp/main
```
