package jeiko
 
import (
    "html/template"
	"net/http"
    "appengine"
    "appengine/datastore"
)

type Stock struct {
        Empresa  string
        Puntos int
}

var homeTemplate = template.Must(template.New("home").Parse(`
<html>
  <head>
    <title>Stock Market Values</title>
  </head>
  <body>
    {{range .}}
      <p>{{.Empresa}}</p>
      <p>{{.Puntos}}</p>
    {{end}}
  </body>
</html>
`))
 
func init() {
	http.HandleFunc("/", handleStart)
}
 
func handleStart(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

	c.Infof("Obteniendo lista de stocks")

	c.Debugf("Consultando datastore.")

    // This query it's not optimized, it should bing only the keys, not all the data
    q := datastore.NewQuery("Stock").Limit(10)
    
    var stocks []Stock
    if _, err := q.GetAll(c, &stocks); err != nil {
		c.Errorf("Error al obtener los stocks desde el datastore. %v", err)
    }
    
	initStock(c, w, stocks)
    
    if err := homeTemplate.Execute(w, stocks); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func initStock(c appengine.Context, w http.ResponseWriter, stocks []Stock) {
    if len(stocks) == 0 {
        //fmt.Fprintln(w, "No existen empresas. Inicializando...")
        googleStock := &Stock{Empresa: "Google", Puntos:  1000}
        amazonStock := &Stock{Empresa: "Amazon",Puntos:  900}
        keyGoogle := datastore.NewIncompleteKey(c, "Stock", nil)
        keyAmazon := datastore.NewIncompleteKey(c, "Stock", nil)
        if _, err := datastore.Put(c, keyGoogle, googleStock); err != nil {
			c.Errorf("Error al inicializar los valores de google. %v", err)
        }
        if _, err := datastore.Put(c, keyAmazon, amazonStock); err != nil {
			c.Errorf("Error al inicializar los valores de amazon. %v", err)
        }
        c.Debugf("datastore inicializado con los valores de prueba inciales")
    }
}
