# httpcodes

A lightweight CLI tool for quick HTTP status code lookups. No more context switching to web browsers or documentation sites.

## What it does

This little tool provides HTTP status code lookups in a purple output. Simply run `httpcodes <code>` to get detailed information about any HTTP status code, including standard codes and extended ones from services like Cloudflare and nginx.

## Installation

### Quick install

```bash
go install github.com/joluc/httpcodes@latest
```

### Build from source

```bash
git clone https://github.com/joluc/httpcodes.git
cd httpcodes
make install
```

## Usage

```bash
httpcodes <status_code>
```

### Examples

```bash
httpcodes 418
```

**Output:**

```
╭─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╮
│                                                                                                                                                 │
│  HTTP 418: I'm a teapot                                                                                                                         │
│                                                                                                                                                 │
│  Any attempt to brew coffee with a teapot should result in the error code "418 I'm a teapot". The resulting entity body MAY be short and stout. │
│                                                                                                                                                 │
╰─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────╯
```

### Help

Run `httpcodes` without arguments to see usage information and examples.

## License

This project is licensed under the MIT License.

## Credits

- Initial HTTP status codes data from [MattIPv4/status-codes](https://github.com/MattIPv4/status-codes) (GPL-3.0 License)
- Added some standard codes (305, 308, 425, 430, 599) for completeness
- Built with [lipgloss](https://github.com/charmbracelet/lipgloss) for beautiful terminal styling
