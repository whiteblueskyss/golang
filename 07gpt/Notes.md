`go mod init go-learning`

go → you’re using the Go tool.
mod → module system (dependency + project management).
init → create a new module.
go-learning → the module name (a unique name for your project).
Go will create a go.mod file that tells Go:
“This folder is a standalone project with its own dependencies.”

Go is statically typed, meaning types are checked at compile time.

WHAT IS net/http?

net/http is Go’s built-in web server and HTTP client library.

You don’t install anything. It’s part of Go’s standard library.

It allows you to:

create a web server

handle routes/endpoints

read HTTP requests

send HTTP responses

work with GET/POST/PUT/DELETE

parse JSON

set headers

build middleware

create HTTP clients

Everything needed for REST APIs is possible with just this package.

Decode(&varialbe), endcode(varialbe)
Decode needs a pointer because it must fill your struct.
Passing a struct (value) creates a copy, so Decode cannot modify it.
Passing a pointer lets Decode write the JSON data into your actual variable.

Standard ServeMux is:

fast

stable

in the std library

But it has limitations:

No path parameters like /students/{id}

No middleware chaining

No groups like /api/v1

Harder to scale for large APIs

So most Go backend developers use routers like:

github.com/go-chi/chi (recommended)

gin

fiber

We will use chi because:

very similar to React Router

clean

fast

widely used

perfect for REST APIs

easy to learn

A mux is simply:
A component that matches the incoming request path and chooses the correct handler.

Chi supports:

r.Get()
r.Post()
r.Put()
r.Delete()
r.Patch()
