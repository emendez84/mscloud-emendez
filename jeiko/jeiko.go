package jeiko
 
import (
    "fmt"
    "encoding/json"
	"net/http"
    "appengine"
    "appengine/datastore"
)

type Stock struct {
        Empresa  string
        Puntos int
}
 
func init() {
	http.HandleFunc("/", handleStart)
}
 
func handleStart(w http.ResponseWriter, r *http.Request) {
    c := appengine.NewContext(r)

	c.Infof("Obteniendo lista de stocks")

	c.Debugf("Consultando datastore.")

    // This query it's not optimized, it should bing only the keys, not all the data
    q := datastore.NewQuery("Stock")
    
    var stocks []Stock
    if _, err := q.GetAll(c, &stocks); err != nil {
		c.Errorf("Error al obtener los stocks desde el datastore. %v", err)
    }
    
	initStock(c, w, stocks)
    
    b, err := json.Marshal(stocks)
    if err != nil {
        fmt.Println(err)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    
    fmt.Fprintln(w, string(b))
    
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
