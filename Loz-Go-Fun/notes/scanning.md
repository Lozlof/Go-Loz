The Go programming language provides a family of functions to scan and parse formatted input. These functions are split based on where they read the input from:

1. **Scan, Scanf, and Scanln** – These functions read from the standard input (`os.Stdin`).
2. **Fscan, Fscanf, and Fscanln** – These read from any `io.Reader` specified by the user.
3. **Sscan, Sscanf, and Sscanln** – These read from a given string.

### Function Behaviors and Differences

- **Scan, Fscan, and Sscan**: Treat all whitespace (except newline) as equivalent, so newlines and spaces are handled the same.
- **Scanln, Fscanln, and Sscanln**: Stop reading once a newline is encountered, meaning the input items must be on a single line and followed by a newline or end of file (EOF).
- **Scanf, Fscanf, and Sscanf**: Use a **format string** to parse the input. This format string uses special symbols (verbs) to describe how to interpret the input data (similar to Printf's verbs). Each verb, beginning with `%`, tells Go to interpret input in a specific way, like scanning an integer in hexadecimal with `%x`.

### Special Formatting Details

In the format string, spacing rules are specific:
- **Space and newline rules**: Spaces in the format string are flexible, matching any amount of whitespace in the input, but newlines must match exactly.
- **Other characters**: Characters in the format string that aren’t `%`, space, or newline (like letters or symbols) must appear in the input exactly as they are in the format.
- **Width**: A width specified (e.g., `%5s`) limits the maximum characters read. For example, `%5s` will read up to five characters for a string, but precision formatting (like `%5.2f`) is not supported.

### Verbs and Parsing

Verbs in the format string determine how the input is parsed:
- Common verbs like `%x` and `%v` handle specific types, where `%x` reads hexadecimal, `%v` handles a default format, and `%s` scans until the first whitespace or newline.
- Numeric formats recognize common prefixes, like `0x` for hexadecimal, `0b` for binary, and `0o` for octal.

### Key Differences from C's `scanf`

Unlike C:
- Newlines are handled more strictly, meaning Go will interpret newlines differently and may require specific spaces.
- Go automatically skips leading spaces for most verbs, making the input less strict than in C.

### Other Features

- **Error Handling**: If fewer items are scanned than provided, an error is returned.
- **Rune Reading**: Functions like `Fscan` may read one character too many (a "rune") beyond what’s returned, which might skip input in a loop if not handled carefully.
- **Scanner Interface**: If an object implements the `Scanner` interface with a `Scan` method, it will control how the input is parsed for that object.

When it’s said that *Scan*, *Fscan*, and *Sscan* treat all whitespace as equivalent (except newlines), it means these functions ignore distinctions between different types of whitespace (like spaces, tabs, and multiple spaces) and treat them as if they are the same. So, any whitespace between values—whether a single space, multiple spaces, or tabs—will be interpreted as a single delimiter.

However, newlines are treated as distinct and will not be ignored. Instead, they act as a boundary that could stop input processing in certain functions (like *Scanln*, *Fscanln*, and *Sscanln*).
