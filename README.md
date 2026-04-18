# BookBot

BookBot is a CLI program that reads a book (plain text file) and prints a statistical report of word and character usage.

This is a reimplementation in Go of the guided project from [Boot.dev](https://www.boot.dev).

See the original Python version: [bookbot-py](https://github.com/zoumas/bookbot-py).

## Usage

```sh
go run . <path_to_book>
```

### Example

```sh
go run . books/frankenstein.txt
```

```
============ BOOKBOT ============
Analyzing book found at books/frankenstein.txt...
----------- Word Count ----------
Found 75767 total words
--------- Character Count -------
e: 44538
t: 29493
a: 25894
o: 24494
i: 23927
n: 23643
s: 20360
r: 20079
h: 19176
d: 16318
l: 12306
...
============= END ===============
```

Books are stored as `.txt` files in the `books/` directory (git-ignored).

## Project Structure

```
├── main.go    # Entry point — reads the book, runs analysis, prints report
├── stats.go   # Pure functions for word and character counting
└── books/     # Plain text book files (not committed)
```

## What I Learned

- **Reading files** — Using `os.ReadFile` for simple one-shot file reads.
- **String manipulation** — Splitting text with `strings.Fields`, lowercasing with `strings.ToLower`, and iterating over runes.
- **Maps** — Building frequency maps (`map[rune]int`) to count character occurrences.
- **Sorting** — Sorting slices with `slices.SortFunc` and `cmp.Compare` from the standard library.
- **Structs** — Defining `CharacterCount` to pair characters with their frequencies for sorting.
- **Unicode** — Filtering non-alphabetic characters with `unicode.IsLetter`, supporting runes beyond ASCII.
- **CLI arguments** — Using `os.Args` to accept file paths from the command line.
- **Project structure** — Separating concerns across files and using `.gitignore` to keep data and binaries out of version control.
