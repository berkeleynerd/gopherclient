# gopherclient

gopherclient is a cross-platform Gopher (RFC 1436) GUI client written in Go.
Standard features include:

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

You'll need `webkitgtk3-devel`:

```#!bash
$ sudo dnf install webkitgtk3-devel
```

## Usage

```#!bash
$ gopherclient
```

## Related

Related projects:

- [go-gopher](https://git.mills.io/prologic/go-gopher)
  go-gopher is the Gopher client and server library used by gopherclient

- [gopherproxy](https://git.mills.io/prologic/gopherproxy)
  gopherproxy is Gopher to HTTP proxy that uses go-gopher
  for all of its core functionality.

## License

MIT
