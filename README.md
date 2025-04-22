# CtrlCSS

CtrlCSS is a minimal utility class CSS generator.

It scans template files (e.g. .heex, .html) from stdin, extracts class names, matches them against a user-defined theme, and emits only the necessary CSS rules.

## Usage

```bash
cat index.html | ctrlcss > styles.css
```

## Installation

### From source

```bash
git clone https://github.com/andreamancuso/ctrlcss.git
cd ctrlcss
go build -o ctrlcss ./cmd/ctrlcss
```

### Requirements
- Go 1.18 or newer

## Theme Configuration

Create a `ctrlcss.theme.toml` file in the working directory:

```toml
[spacing]
0 = "0"
1 = "0.25rem"
2 = "0.5rem"
4 = "1rem"
8 = "2rem"
12 = "3rem"

[colors]
white = "#ffffff"
black = "#000000"
blue-500 = "#3b82f6"
gray-800 = "#1f2937"
```

## Output

CtrlCSS emits one CSS rule per line, based on matched classes.

## License

MIT

