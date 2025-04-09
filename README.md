
# 🛠️ CodeSnip-Go: Markdown Snippet Magic

Welcome to the **CodeSnip-Go** demo repo!
This is a quick, fun playground to show a CLI tool that **automagically injects code snippets** from Go files right into your Markdown files. No copy-pasta needed ✨

## Tags


* `<code>`
* `<go>`
* `<cmd>`


---

## 🐣 The Beginning: Hello World

Let's start with the obligatory classic:

```go
func main() {
	HelloWorld()
	fmt.Println("2 + 3 =", Add(2, 3))
	fmt.Println("Random number:", RandomNumber())
}
```
> *source: main.go:main*


## 🔢 Doing some math

Here's where things get spicy (or at least mildly seasoned):

```go
func Add(a int, b int) int {
	return a + b
}
```
> *source: main.go:add*


## 🎲 Randomness appears

Who doesn't like a bit of chaos?
```go
func RandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
```
> *source: main.go:randomNumber*


## 📟 Output for CMD commands

Go specific commands:

```shell
$ go run .

Hello, world!
2 + 3 = 5
Random number: 34

--------------------------------------------------------------------------------
Go Version: go1.24.1

```

Repo structure:

```shell
$ tree .

.
├── go.mod
├── hype.md
├── main_test.go
├── main.go
├── Makefile
└── README.md

1 directory, 6 files
```

Or even test results and coverage:

```shell
$ make test

go test -v -cover -race ./...
=== RUN   TestHelloWorld
Hello, world!
--- PASS: TestHelloWorld (0.00s)
=== RUN   TestAdd
--- PASS: TestAdd (0.00s)
=== RUN   TestRandomNumber
--- PASS: TestRandomNumber (0.00s)
PASS
coverage: 57.1% of statements
ok  	hype/demo	(cached)	coverage: 57.1% of statements
```

## 🔨 Needs a Fix


* Fails silently on error (e.g. typos, wrong use of cli, duplicate tag names, referencing non-existent tags)

* Non-clickable path to the source


