# lnkparse

A tool to parse .lnk files

## Installation

- Download a prebuilt binary from [release page](https://github.com/sleepytariq/lnkparse/releases/latest)

  _or_
- `git clone https://github.com/sleepytariq/lnkparse && cd lnkparse && go build -ldflags="-s -w" .`

## Usage

```console
Usage:
  lnkparse [flags] <PATH>

Examples:
  lnkparse C:\Users\user\Desktop\file.lnk
  lnkparse C:\Users\user\Desktop\*.lnk
  lnkparse C:\Users\user\**\*.lnk

Flags:
  -trim        trim leading/trailing spaces in command line
  -h, -help    show this message and exit
  -version     show version and exi
```

## References

- https://github.com/libyal/liblnk/blob/main/documentation/Windows%20Shortcut%20File%20(LNK)%20format.asciidoc
