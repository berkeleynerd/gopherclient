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
$ go install github.com/prologic/gopherclient
```

### OS X Homebrew

```#!bash
$ brew tap prologic/gopherclient
$ brew install --HEAD gopherclient
```

Make sure you have QT 5.6+ installed:

```#!bash
$ brew install qt
```

gopherclient is still early days so contributions, ideas and UI expertise are
much appreciated and highly welome!

## Usage

```#!bash
GODEBUG=cgocheck=0 gopherclient
```

**NB:** The `GODEBUG=cgocheck=0` is necessary to successfully run gopherclient
        at this time due to Go 1.6+ compatibility issues with go-qml.
        See: go-qml/qml#170 and go-qml/qml#179 (Sorry!)

## Licnese

MIT
