package main

//go:generate rice embed-go

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"runtime"
	"strings"

	"github.com/GeertJohan/go.rice"
	"github.com/mitchellh/go-homedir"
	"github.com/namsral/flag"
	"github.com/prologic/go-gopher"
	"github.com/prologic/gopherproxy"
	"github.com/zserge/webview"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

var (
	gopherHome string
)

func ensureGopherHome(root string) error {
	err := os.MkdirAll(root, 0755)
	if err != nil {
		return err
	}

	return nil
}

func localGopherServer(bind, root string) {
	gopher.Handle("/", gopher.FileServer(gopher.Dir(root)))
	log.Fatal(gopher.ListenAndServe(bind, nil))
}

func startServer(uri string) string {
	var (
		content string
		tpl     *template.Template
	)

	bs, err := ioutil.ReadFile(".template")
	if err == nil {
		content = string(bs)
	} else {
		content = string(defaultTemplate)
	}

	tpl, err = template.New("index").Parse(content)
	if err != nil {
		log.Fatal(err)
	}

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", gopherproxy.GopherHandler(tpl, nil, uri))
		http.Handle("/assets", http.FileServer(rice.MustFindBox("assets").HTTPBox()))
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

func handleRPC(w webview.WebView, data string) {
	switch {
	case data == "back":
		log.Println("back")
	case data == "forwrd":
		log.Println("forward")
	case data == "reload":
		log.Println("reload")
	case data == "home":
		log.Println("home")
	case strings.HasPrefix(data, "open:"):
		log.Println("open")
	}
}

func init() {
	runtime.LockOSThread()
}

func main() {
	var (
		err error

		version bool
	)

	flag.BoolVar(&version, "v", false, "display version information")
	flag.Parse()

	if version {
		fmt.Printf("gopherclient v%s", FullVersion())
		os.Exit(0)
	}

	gopherHome, err = homedir.Expand("~/.gopher")
	if err != nil {
		log.Fatal(err)
	}

	err = ensureGopherHome(gopherHome)
	if err != nil {
		log.Fatal(err)
	}

	url := startServer("floodgap.com")
	w := webview.New(webview.Settings{
		Width:  windowWidth,
		Height: windowHeight,
		Title:  "Gopher Client",
		URL:    url,
		ExternalInvokeCallback: handleRPC,
	})
	defer w.Exit()
	w.Run()
}
