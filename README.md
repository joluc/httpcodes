# httpcodes

I was bored to google random HTTP status codes, so I wrote this little Go helper to show me the info directly in my CLI.

## What it does

Tired of context switching to look up HTTP status codes? This tiny tool gives you instant status code lookups right in your terminal with purple output. Just run `httpcodes <code>`.

## Get it running

### The lazy way (recommended)

```bash
go install github.com/joluc/httpcodes@latest
```

### Or build it yourself

```bash
git clone https://github.com/joluc/httpcodes.git
cd httpcodes
make install
```

## How to use it

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

### Need help?

Just run `httpcodes` without any arguments and it'll show you what to do.

## Licence

This project is licensed under the MIT License (do whatever you want with it).

## Props

- Initial HTTP status codes data comes from the awesome [MattIPv4/status-codes](https://github.com/MattIPv4/status-codes) repo (GPL-3.0 License)
- Added some missing codes (305, 308, 425, 430, 599)
- Built with [lipgloss](https://github.com/charmbracelet/lipgloss) for those sweet purple tables
