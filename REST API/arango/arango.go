package arango

import (
	"context"
	"crypto/tls"
	"fmt"
	"time"

	driver "github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

//

var botoncito []Boton

type Boton struct {
	ID     string `json:"ID"`
	LUMINO string `json:"LUMINO"`

	TEMP    string `json:"TEMP"`
	HUMEDAD string `json:"HUMEDAD"`
	PH      string `json:"PH"`

	BOMBA string `json:"BOMBA"`
	UV    string `json:"UV"`
}

var conexion, error = http.NewConnection(http.ConnectionConfig{
	Endpoints: []string{"http://localhost:900"},
	TLSConfig: &tls.Config{ /*...*/ },
})
var co, errorr = driver.NewClient(driver.ClientConfig{
	Connection:     conexion,
	Authentication: driver.BasicAuthentication("root", "rozenmaiden"),
})
var ctxe = context.Background()
var dbe, errrr = co.Database(ctxe, "_system")
var colecc, errrrr = dbe.Collection(ctxe, "boton_botones")

func Conexion(doc Boton) {
	colecc.CreateDocument(ctxe, doc)
	time.Sleep(1 * time.Second)
}

func Obtener(key string) Boton {

	var cole Boton
	colecc.ReadDocument(ctxe, key, &cole)
	return cole

}

func ObtenerTodo() []Boton {

	mula, erro123 := colecc.Count(ctxe)
	if erro123 != nil {
		// handle errorlo
	}
	//fmt.Println(mula)
	//var vici [100]string
	vici := make([]string, 100050)
	mila := vici[:mula]
	notf := make([]Boton, 100050)
	notd := notf[:mula]
	var i int
	query := "FOR d IN boton_botones RETURN d"
	cursor, err := dbe.Query(ctxe, query, nil)
	if err != nil {
		// handle error
		//log.Fatal(err)
	}
	defer cursor.Close()
	for {
		var doc Boton
		meta, err := cursor.ReadDocument(ctxe, &doc)
		if driver.IsNoMoreDocuments(err) {
			break
		} else if err != nil {
			// handle other errors
		}
		fmt.Printf("Got doc with key '%s' from query\n", meta.Key)
		mila[i] = meta.Key
		notd[i] = doc
		i = i + 1
		fmt.Println(doc)

	}
	fmt.Println(mila)
	fmt.Println(notd)
	return notd
}
