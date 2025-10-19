🗓️ Go Learning Plan — 30 Days (15 min/day)
🔹 Week 1 — Go Basics

Focus: syntax, variables, loops, functions

Day 1: Install Go

Install from https://go.dev/dl
Run: go version

Create a file main.go and print “Hello, Go!”

Day 2: Variables & types

Learn: var, :=, const

Practice: declare string, int, bool

Print them using fmt.Println

Day 3: Basic input/output

Use fmt.Scan for simple input

Try: read a name, print greeting

Day 4: Conditionals

Write a program that checks if a number is even or odd

Day 5: Loops

for is the only loop in Go

Print numbers 1 → 10 using for

Day 6: Arrays & slices

Create a slice of strings

Loop through it and print values

Day 7: Functions

Write and call add(a, b int) int

Bonus: return multiple values

🔹 Week 2 — Structs, Methods & Packages

Focus: organizing code like classes in Python

Day 8: Structs

Create a struct User with Name, Age

Initialize and print it

Day 9: Methods

Add a method Greet() for User that prints a message

Day 10: Pointers

Learn & (address) and * (dereference)

Modify struct fields using a pointer

Day 11: Packages & imports

Split your project: main.go and greet/greet.go

Import your custom package

Day 12: Error handling

Use errors.New() or fmt.Errorf()

Create a function that returns (int, error)

Day 13: Maps

Create a map [string]int for student scores

Add, delete, lookup keys

Day 14: Review day

Rebuild a small program combining structs, maps, and loops

🔹 Week 3 — Concurrency & APIs

Focus: goroutines, channels, HTTP servers

Day 15: Goroutines

Run two functions at once using go funcName()

Day 16: Channels

Create a channel to send and receive strings between goroutines

Day 17: WaitGroups

Use sync.WaitGroup to wait for multiple goroutines

Day 18: Build a tiny HTTP server

Use net/http

Example:
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Hello from Go server!")
})
http.ListenAndServe(":8080", nil)

Day 19: Handle routes

Add /about route returning “About page”

Day 20: JSON response

Return JSON with json.Marshal() and w.Header().Set("Content-Type", "application/json")

Day 21: Simple API project

Create a mini CRUD API with in-memory slice (no DB yet)

Week 4 — Databases, Testing & Final Project

Focus: PostgreSQL, project building

Day 22: Connect to PostgreSQL

Install lib/pq driver

Test connection using sql.Open

Day 23: Simple query

Create table users

Write code to INSERT and SELECT

Day 24: Build simple REST API (users CRUD)

Use net/http

Implement GET and POST routes

Day 25: Add PUT and DELETE endpoints

Update and delete user by ID

Day 26: Error handling & logging

Use log.Fatal(err)

Return proper HTTP status codes

Day 27: Unit testing basics

Create main_test.go

Write a simple test using testing package

Day 28: Final project polish

Add JSON responses, handle errors cleanly

Day 29: Build & run

go build . → run executable

Test all routes with curl or Postman

Day 30: Review & reflect

Summarize what you learned

Plan a next project (e.g., Go microservice for your Django ERP)

✅ After 30 days you’ll:

Understand Go syntax deeply

Know how to build APIs

Use goroutines/channels for concurrency

Connect Go with PostgreSQL

Have a working mini backend project 🎯