# Exercise 4: Interacting with REST APIs (cURL & Go)

This exercise requires **no coding work**. It is meant to be a walkthrough of how to interact with external APIs from
your terminal using `curl`, and then how to do the same thing using Go's standard library.

We will be exploring the Studio Ghibli API (`https://ghibliapi.vercel.app`), which is a free and open read-only REST
API.

## Part 1: Terminal Exploration with `curl`

`curl` is a command-line tool used for transferring data to and from a server. It is one of the most common tools for
testing APIs.

Open your terminal and try running the following commands:

### 1. Fetching all films

```bash
curl https://ghibliapi.vercel.app/films
```

### 2. Fetching a specific film by ID

```bash
curl https://ghibliapi.vercel.app/films/58611129-2dbc-4a81-a72f-77ddfc1b1b49
```

### 3. Using Query Parameters

You can filter or limit results using the `?` syntax in the URL. Be sure to wrap the URL in quotes so your terminal
doesn't misinterpret special characters!

```bash
curl "https://ghibliapi.vercel.app/people?limit=3"
```

### 4. Fetching ONLY the HTTP Headers

Sometimes you just want to see the metadata (status code, content type) without downloading the entire JSON body. Use
`-I` (capital i) to fetch headers only.

```bash
curl -I https://ghibliapi.vercel.app/films
```

---

## Part 2: Fetching data in Go

Now that we know how the API behaves in the terminal, let's write a Go program that makes the exact same request.

In Go, we use the built-in `net/http` package to make requests, and `encoding/json` to parse the JSON response into Go
structs.

Open `example.go` in this directory to see a complete working example.

You can run the example yourself by navigating to this directory in your terminal and running:

```bash
go run example.go
```
