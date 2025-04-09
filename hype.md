# 🛠️ CodeSnip-Go: Markdown Snippet Magic

Welcome to the **CodeSnip-Go** demo repo!
This is a quick, fun playground to show a CLI tool that **automagically injects code snippets** from Go files right into your Markdown files. No copy-pasta needed ✨

## Tags

- `<code>`
- `<go>`
- `<cmd>`

---

## 🐣 The Beginning: Hello World

Let's start with the obligatory classic:

<code src="main.go" snippet="main"></code>


## 🔢 Doing some math

Here's where things get spicy (or at least mildly seasoned):

<code src="main.go" snippet="add"></code>

## 🎲 Randomness appears

Who doesn't like a bit of chaos?
<code src="main.go" snippet="randomNumber"></code>

## 📟 Output for CMD commands

Go specific commands:

<go run="."></go>

Repo structure:

<cmd exec="tree ."></cmd>

Or even test results and coverage:

<!-- <go test="-v -cover -race ./..."></go> -->
<!-- <cmd exec="go test -v -cover -race ./..."></exec> -->
<cmd exec="make test"></cmd>

## 🔨 Needs a Fix

- Fails silently on error (e.g. typos, wrong use of cli, duplicate tag names, referencing non-existent tags)
<!-- <go test="-v -cover -racce ./..."></go> -->

- Non-clickable path to the source
