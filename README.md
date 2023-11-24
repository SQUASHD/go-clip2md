# go-clip2md

`go-clip2md` is a CLI tool that seamlessly converts HTML content from your
clipboard into Markdown format. Especially useful for anyone who used Obsidian.

## Features

- **Clipboard Conversion**: Instantly convert clipboard content from HTML to Markdown.
- **Watch Mode**: Automatically monitor clipboard changes and write Markdown files to a specified directory.
- **Flexible Output**: Choose between appending to a single file or creating new files for each clipboard content change.
- **Customizable**: Set domain for proper image path resolution, and specify language for code block formatting.

## Copyright and Responsible Use

`go-clip2md` is a tool for converting content from HTML to Markdown. Users are
responsible for ensuring that they have the necessary rights to copy, convert,
and use the content they process with this tool. Content copied from the web or
other sources is typically subject to copyright law. 

## Intended Use Case

`go-clip2md` is ideal for anyone looking to aggregate documentation or notes in Markdown format. It's particularly useful for:

- Obsidian users wanting a rich text preview of Markdown documents.
- Developers and researchers who compile information from various online sources into organized Markdown files.

## Installation

Install `go-clip2md` using Go:

```bash
go install github.com/SQUASHD/go-clip2md@latest
```

This command handles all dependencies and creates a binary for your operating system.

## Usage

Basic command to convert clipboard content to Markdown:

```bash
go-clip2md
```

To watch the clipboard and automatically write files:

```bash
go-clip2md watch [flags]
```

### Flags

Global Flags (applicable to all commands):

- `--domain` (or `-d`): Specify the domain for copying content, for correct image path resolution.
- `--lang` (or `-l`): Set the language for code block conversion. Note: this does not handle the case where there is no langauge specified in the codeblock

Watch Command Flags:

- `--out` (or `-o`): Directory to write files (default: current working directory).
- `--interval` (or `-i`): Clipboard check interval in seconds (default: 1).
- `--pattern` (or `-p`): Prefix for file names, followed by an auto-incremented count, e.g., `doc1-header.md`, `doc2-header.md`.
- `--mode` (or `-m`): File writing mode ( `append` or `new`). Default mode is `new`.

## Contributing

Contributions, fixes, and feature requests are welcome. 

## License

`go-clip2md` is MIT licensed, as found in the [license file](./LICENSE/go-clip2md).

## Third-Party Licenses

This project uses the following open-source libraries:

- [goquery](https://github.com/PuerkitoBio/goquery) - Licensed under BSD 3-Clause License. [View License](./LICENSE/goquery).
- [html-to-markdown](https://github.com/JohannesKaufmann/html-to-markdown) - Licensed under MIT License. [View License](./LICENSE/html-to-markdown).
- [clipboard](https://github.com/atotto/clipboard) - Licensed under BSD 3-Clause License. [View License](./LICENSE/clipboard).

