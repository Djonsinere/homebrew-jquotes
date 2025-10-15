# Japanese Quotes

A small Go command-line utility that prints a random inspirational quote in Japanese and English inside a simple ASCII "card" sized to your terminal.

This project contains a collection of English quotes and their Japanese equivalents. The program selects a random quote and prints it centered inside a bordered box, using ANSI color escapes for nicer terminal presentation.

## HomeBrew

```
brew tap Djonsinere/jquotes
brew install japanese-quotes
```


Files of interest

- `main.go` – program entry point. Calls the library functions to select and print a quote.
- `src/config.go` – contains the `Quote` type and the `Quote_bank` slice with many bilingual entries.
- `src/random.go` – selects a random quote from `Quote_bank`.
- `src/paint.go` – formats and prints the quote inside a bordered box, and reads terminal size.

Requirements

- Go 1.24 or newer (module requires Go 1.24.2 in `go.mod`)
- A POSIX-like terminal (code uses ioctl to read terminal size; works on macOS and Linux terminals)

##Build

From the repository root:

```bash
go build -o japanese_quotes
```

This will produce a `japanese_quotes` executable in the current directory.

Run

```bash
./japanese_quotes
```

Example output

The program prints a box like:

![example.png](example.png)

Notes and small caveats

- The terminal size read uses a direct ioctl call that is platform-dependent; it is written to work on macOS and Linux. If it fails, the program will still attempt to print but the layout may be off.
- Character width for Japanese text is estimated by counting characters >127 as double-width; this is a heuristic and may not be perfect for all terminals or fonts.

Contributing

Contributions are welcome. Suggestions:

- Add or improve quotes in `src/config.go`.
- Improve terminal width calculation and multi-line wrapping.
- Add tests and CI for cross-platform behavior.

License

This repository includes a `LICENSE` file. Follow the terms in that file when contributing or reusing code.

Author

Djonsinere (repository owner)
