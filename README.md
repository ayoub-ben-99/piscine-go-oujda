# Piscine Go — Oujda

![Go](https://img.shields.io/badge/Go-1.25.9-00ADD8?style=flat-square&logo=go)
![Zone01](https://img.shields.io/badge/Zone01-Oujda-blue?style=flat-square)
![Exercises](https://img.shields.io/badge/Exercises-77-green?style=flat-square)

> My solutions to all Piscine Go exercises at Zone01 Oujda — from Go basics to Binary Trees and Linked Lists.

---

## 📖 About

This repository contains my solutions to the **Piscine Go** exercises. It covers fundamentals, string manipulation, math, slices, linked lists, binary trees, and small projects.

**Module:** `piscine`
**Dependency:** `github.com/01-edu/z01 v0.1.0`

---

## 📂 Content

### 🔤 Basics

| File | Description |
|------|-------------|
| `printstr.go` | Print a string |
| `strlen.go` | String length |
| `strrev.go` | Reverse a string |
| `swap.go` | Swap two values |
| `pointone.go` / `ultimatepointone.go` | Pointers |
| `divmod.go` / `ultimatedivmod.go` | Division and modulo |
| `hello.sh` / `myfamily.sh` / `who-are-you.sh` | Shell basics |

### 🔢 Math & Logic

| File | Description |
|------|-------------|
| `iterativefactorial.go` / `recursivefactorial.go` | Factorial |
| `iterativepower.go` / `recursivepower.go` | Power |
| `fibonacci.go` | Fibonacci |
| `sqrt.go` | Square root |
| `isprime.go` / `findnextprime.go` | Prime numbers |
| `activebits.go` | Count active bits |
| `max.go` | Max value |
| `atoi.go` / `basicatoi.go` / `basicatoi2.go` | String to int |
| `collatzcountdown.go` | Collatz countdown |

### 🔡 Strings

| File | Description |
|------|-------------|
| `rot14.go` | ROT14 cipher |
| `join.go` | Join strings |
| `split.go` / `splitwhitespaces.go` | Split strings |
| `concatparams.go` | Concat params |
| `printwordstables.go` | Print words table |
| `stringtointslice.go` | String to int slice |
| `enigma.go` | Enigma |
| `jumpover.go` | Jump Over |

### 📦 Slices & Arrays

| File | Description |
|------|-------------|
| `appendrange.go` / `makerange.go` | Make range |
| `descendappendrange.go` / `descendcomb.go` | Descending combinations |
| `compact.go` | Compact slice |
| `sortintegertable.go` | Sort integer table |
| `reversemenuindex.go` | Reverse menu index |
| `unmatch.go` | Unmatch |
| `countif.go` / `any.go` / `foreach.go` / `map.go` | Functional helpers |
| `issorted.go` | Check if sorted |

### 🔗 Linked Lists

| File | Description |
|------|-------------|
| `listpushfront.go` / `listpushback.go` | Push front / back |
| `listsize.go` / `listlast.go` / `listat.go` | Size / Last / At |
| `listclear.go` / `listreverse.go` / `listmerge.go` | Clear / Reverse / Merge |
| `listfind.go` / `listforeach.go` / `listforeachif.go` | Find / ForEach |
| `listremoveif.go` | RemoveIf |

### 🌳 Binary Trees

| File | Description |
|------|-------------|
| `btreeinsertdata.go` | Insert data |
| `btreeapplyinorder.go` / `btreapplypreorder.go` / `btreapplypostorder.go` / `btreeapplybylevel.go` | Tree traversal |
| `btreesearchitem.go` | Search item |
| `btreelevelcount.go` | Level count |
| `btreemin.go` / `btreemax.go` | Min / Max |
| `btreeisbinary.go` / `btreetransplant.go` | IsBinary / Transplant |

### 🍔 Projects

| File | Description |
|------|-------------|
| `loafofbread.go` | Loaf of Bread |
| `fooddeliverytime.go` | Food Delivery Time |
| `dealapackofcards.go` | Deal a Pack of Cards |
| `rockandroll.go` | Rock and Roll |
| `podiumposition.go` | Podium Position |
| `shoppinglistsort.go` / `shoppingsummarycounter.go` | Shopping project |
| `abort.go` | Abort |

---

## 🚀 Usage

### Requirements

- Go `1.25.9` or later

```bash
go version
```

### Clone

```bash
git clone https://github.com/<username>/piscine-go-oujda.git
cd piscine-go-oujda
```

### Install dependencies

```bash
go mod tidy
```

### Run an exercise

```bash
go run <file>.go
# Example:
go run fibonacci.go
```

Or create a temporary `main.go` to call a function:

```go
package main

import (
    "fmt"
    "piscine"
)

func main() {
    fmt.Println(piscine.Fibonacci(6)) // 8
}
```

```bash
go run main.go
rm main.go
```

---

## 📁 Project Structure

```
piscine-go-oujda/
├── *.go            # 77 exercise solutions
├── go.mod          # module piscine + z01
├── go.sum
├── comcheck/
├── fixthemain/
├── pilot/
├── printparams/
├── printprogramname/
├── revparams/
├── sortparams/
├── test/
└── README.md
```

---

## ✅ Checkpoint

My checkpoint exam solutions are maintained in a separate repository:

- **[checkpoint-01](https://github.com/ayoub-ben-99/checkpoint-01)** — Pool checkpoint exercises at Zone01 Oujda (Levels 2–10: `checknumber`, `firstword`, `gcd`, `itoa`, `addprimesum`, `brackets`, `rpncalc`, `brainfuck`, and more).

---

## ⚠️ Disclaimer

> These solutions are for learning purposes only. If you are currently in the Piscine, try solving the exercises yourself first before checking them.

---

## 👤 Author

**abenrkia** — Zone01 Oujda

- GitHub: [@abenrkia](https://github.com/abenrkia)
- School: [Zone01 Oujda](https://learn.zone01oujda.ma)

---

## ⭐ Support

If you like this project, don't forget to give it a ⭐ on GitHub!
