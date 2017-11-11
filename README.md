# gopherclient

gopherclient is a cross-platform Gopher (RFC 1436) GUI client written in Go
using the QT toolkit with QML. Standard features include:

- Back
- Forward
- Refresh
- Home

Coming soon:

- Bookmarks
- Search
- Tabs
- Download Manager

![Gopher Client](/screenshot.png?raw=true "Gopher Client")

## Installation

### Source

```#!bash
$ make && make install
```

### OS X Homebrew

```#!bash
$ brew tap prologic/gopherclient
$ brew install --HEAD gopherclient
```

gopherclient is still early days so contributions, ideas and UI expertise are
much appreciated and highly welome!

### Linux

Fedora:

You'll need `qt5-qtbase-devel` and `webkitgtk3-devel`:

```#!bash
$ sudo dnf install qt5-qtbase-devel webkitgtk3-devel
```

## Usage

```#!bash
$ gopherclient
```

## Licnese

MIT
