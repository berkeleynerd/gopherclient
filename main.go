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
	w          webview.WebView
	server     *Server
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

type Server struct {
	url  string
	home string

	tpl *template.Template
}

func NewServer(home string) *Server {
	return &Server{
		home: home,
	}
}

func (s *Server) Start() string {
	var content string

	bs, err := ioutil.ReadFile(".template")
	if err == nil {
		content = string(bs)
	} else {
		content = string(defaultTemplate)
	}

	tpl, err := template.New("index").Parse(content)
	if err != nil {
		log.Fatal(err)
	}
	s.tpl = tpl

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer ln.Close()
		http.HandleFunc("/", s.Handler)
		http.Handle("/assets", http.FileServer(rice.MustFindBox("assets").HTTPBox()))
		log.Fatal(http.Serve(ln, nil))
	}()
	return "http://" + ln.Addr().String()
}

func (s *Server) SetHome(url string) {
	s.home = url
}

func (s *Server) Back() {
	w.Eval("window.history.back();")
}

func (s *Server) Forward() {
	w.Eval("window.history.forward();")
}

func (s *Server) Home() {
	s.Open(s.home)
}

func (s *Server) Reload() {
	w.Eval(fmt.Sprintf("window.location.reload();"))
}

func (s *Server) Open(url string) {
	s.url = url
	w.Eval(fmt.Sprintf("window.location.pathname = \"%s\";", url))
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	gopherproxy.GopherHandler(s.tpl, nil, s.url).ServeHTTP(w, r)
}

func (s *Server) HandleRPC(w webview.WebView, data string) {
	args := strings.Split(data, ":")

	switch args[0] {
	case "back":
		s.Back()
	case "forwrd":
		s.Forward()
	case "reload":
		s.Reload()
	case "home":
		s.Home()
	case "open":
		s.Open(args[1])
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

	server = NewServer("floodgap.com")
	url := server.Start()

	w = webview.New(webview.Settings{
		Width:                  windowWidth,
		Height:                 windowHeight,
		Title:                  "Gopher Client",
		URL:                    url,
		ExternalInvokeCallback: server.HandleRPC,
		Debug:                  true,
	})
	defer w.Exit()

	server.Home()

	w.Run()
}
