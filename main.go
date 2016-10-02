package main

//go:generate genqrc resources

import (
	"fmt"
	"log"
	"net/url"
	"os"

	"github.com/prologic/go-gopher"
	"github.com/prologic/gopherproxy"

	"github.com/mitchellh/go-homedir"

	"gopkg.in/qml.v1"
	"gopkg.in/qml.v1/webengine"
)

var (
	gopherHome string
	uriField   qml.Object
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

func main() {
	var err error

	gopherHome, err = homedir.Expand("~/.gopher")
	if err != nil {
		log.Fatal(err)
	}

	err = ensureGopherHome(gopherHome)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(qml.Run(func() error {
		go localGopherServer("127.0.0.1:7070", gopherHome)
		go gopherproxy.ListenAndServe("127.0.0.1:8070", "127.0.0.1:7070")

		webengine.Initialize()

		engine := qml.NewEngine()
		engine.On("quit", func() { os.Exit(0) })

		component, err := engine.LoadFile("qrc:///resources/main.qml")
		if err != nil {
			return err
		}

		win := component.CreateWindow(nil)

		root := win.Root()

		backButton := root.ObjectByName("backButton")
		forwardButton := root.ObjectByName("forwardButton")
		refreshButton := root.ObjectByName("refreshButton")
		homeButton := root.ObjectByName("homeButton")

		goButton := root.ObjectByName("goButton")
		mainView := root.ObjectByName("mainView")
		uriField = root.ObjectByName("uriField")

		navigateTo := func() {
			uri := uriField.String("text")

			u, e := url.Parse(uri)
			if e != nil {
				log.Printf("ERROR: %s", e)
				return
			}

			mainView.Set(
				"url",
				fmt.Sprintf(
					"http://127.0.0.1:8070/%s/%s",
					u.Host,
					u.Path,
				),
			)
		}

		backButton.On("clicked", func() {
			mainView.Call("goBack")
		})

		forwardButton.On("clicked", func() {
			mainView.Call("goForward")
		})

		refreshButton.On("clicked", func() {
			mainView.Call("reload")
		})

		homeButton.On("clicked", func() {
			mainView.Set("url", "http://127.0.0.1:8070/127.0.0.1:7070")
		})

		goButton.On("clicked", navigateTo)
		uriField.On("accepted", navigateTo)

		win.Show()
		win.Wait()

		return nil
	}))
}
