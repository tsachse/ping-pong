package main

// gopherjs build app.go -o app.js --tags debug -v -w
// gopherjs serve -v --tags debug  --http ":3000"

import (
	"github.com/gopherjs/gopherjs/js"
	"github.com/johanbrandhorst/gopherjs-json"
	vue "github.com/oskca/gopherjs-vue"
	"honnef.co/go/js/xhr"
)

type Model struct {
	*js.Object         // this is needed for bidirectional data bindings
	Test       string  `js:"test"`
	MyResult   *Result `js:"my_result"`
	//NeuesFeld  string  `js:"neues_feld"`
}

func (m *Model) Ping() {
	req := xhr.NewRequest("GET", "/ping")
	req.SetRequestHeader("Content-Type", "application/json")
	req.ResponseType = xhr.JSON

	// Go-Routine aufrufen, damit der GET-Request nicht die Applikation blockiert
	go func() {
		// blockiert bis die Anwort erhalten ist
		err := req.Send(nil)
		if err != nil {
			panic(err)
		}
		if req.Status != 200 {
			panic(req.Response)
		}

		res := &Result{
			Object: req.Response,
		}
		m.MyResult = res
		m.Test = "Ping-Pong mit GopherJS"
	}()

}

func (m *Model) Save() {
	d, err := json.Marshal(m.MyResult.Object)
	if err != nil {
		panic(err)
	}
	req := xhr.NewRequest("POST", "/save")
	req.SetRequestHeader("Content-Type", "application/json")

	// Go-Routine aufrufen, damit der GET-Request nicht die Applikation blockiert
	go func() {
		// blockiert bis die Anwort erhalten ist
		err := req.Send([]byte(d))
		if err != nil {
			panic(err)
		}
		if req.Status != 200 {
			panic(req.Response)
		}

	}()

}

type Result struct {
	*js.Object
	// für den Zugrif in der View müssen die Felder aus der JSON-message
	// hier nicht unbedingt angeben werden.
	//Message   string `js:"message"`

}

func main() {

	m := &Model{
		Object: js.Global.Get("Object").New(),
	}
	// es müssen alle Felder des Models initialisiert werden, sonst
	// funktioniert der Zugriff später aus der View nicht
	m.Test = "beta"
	// m.NeuesFeld = ""

	// hier ein JS Object zugewiesen werden, damit später auf die
	// Felder der JSON-Nachrivt zugegriffen werden kann
	m.MyResult = &Result{
		Object: js.Global.Get("Object").New(),
	}

	vue.New("#app", m)

	println("OK")
	m.Ping()
}
