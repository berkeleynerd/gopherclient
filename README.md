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

Make sure you have QT 5.6+ installed:

```#!bash
$ brew install qt5
```

### OS X Homebrew

```#!bash
$ brew tap prologic/gopherclient
$ brew install --HEAD gopherclient
```

gopherclient is still early days so contributions, ideas and UI expertise are
much appreciated and highly welome!

## Other Platforms

Please note that at this time (2dn October 2016) gopherclient is only supported
and tested on Mac OS X with Homebrew installed qt5 and go. Ubuntu 14.04
and 16.10 have been attempted and faield. In theory it should be possible to
build gopherclient for other platforms as long as you meet the following
requirements:

- QT 5.4+
- QTWebEngine 1.1+
- Go 1.7+

The biggest problem in trying to support Ubuntu was getting QTWebgine installed
(there are no packages yet).

## Usage

```#!bash
GODEBUG=cgocheck=0 gopherclient
```

**NB:** The `GODEBUG=cgocheck=0` is necessary to successfully run gopherclient
        at this time due to Go 1.6+ compatibility issues with go-qml.
        See: [go-qml/qml#170](https://github.com/go-qml/qml/issues/170)
             and [go-qml/qml#179](https://github.com/go-qml/qml/issues/179)
             (Sorry!)

## Licnese

MIT
