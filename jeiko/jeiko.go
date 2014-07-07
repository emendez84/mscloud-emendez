package jeiko
 
import (
    "fmt"
    "bufio"
    "io"
    "strings"
    "strconv"
    "os"
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
    
	initStockFile(c, w, stocks)
    
    b, err := json.Marshal(stocks)
    if err != nil {
        c.Errorf("Error en conversión del Json. %v", err)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    
    fmt.Fprintln(w, string(b))
    
}

func initStock(c appengine.Context, w http.ResponseWriter, stocks []Stock) {
    if len(stocks) == 0 {
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

func initStockFile(c appengine.Context, w http.ResponseWriter, stocks []Stock) {
    c.Debugf("Inicializando por el archivo")
    if len(stocks) == 0 {
        f, err := os.Open("data.txt")
        if err != nil {
            c.Errorf("Error al leer el archivo. %v", err)
        }
        bf := bufio.NewReader(f)
        for {
            switch line, err := bf.ReadString('\n'); err {
            case nil:
                // valid line, echo it.  note that line contains trailing \n.
                i,err := strconv.Atoi(strings.Fields(line)[1])
                if err != nil{
                    c.Errorf("Error de conversion de enteros. %v", err)
                }
                companyStock := &Stock{Empresa:strings.Fields(line)[0] ,Puntos:i}
                keyCompany := datastore.NewIncompleteKey(c,"Stock",nil)
                if _, err := datastore.Put(c, keyCompany, companyStock); err != nil {
                    c.Errorf("Error al inicializar los valores. %v", err)
                 }
            case io.EOF:
                if line > "" {
                    // last line of file missing \n, but still valid
                     i,err := strconv.Atoi(strings.Fields(line)[1])
                    if err != nil{
                        c.Errorf("Error de conversion de enteros. %v", err)
                    }
                    companyStock := &Stock{Empresa:strings.Fields(line)[0] ,Puntos:i}
                    keyCompany := datastore.NewIncompleteKey(c,"Stock",nil)
                    if _, err := datastore.Put(c, keyCompany, companyStock); err != nil {
                        c.Errorf("Error al inicializar los valores. %v", err)
                    }
                }
                return
            default:
                c.Errorf("Error en la lectura de la linea. %v", err)
            }
        }
        c.Debugf("Archivo Leido")
    }
}