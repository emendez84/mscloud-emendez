package nhemu
 
import (
	"fmt"
	"encoding/json"
	"net/http"
	"appengine"
	"appengine/datastore"
	"strconv"
)

type Stock struct {
	Empresa  string
	Puntos int
}
 
func init() {
	http.HandleFunc("/nhemu/", handleStart)
}
 
func handleStart(w http.ResponseWriter, r *http.Request) {

	c := appengine.NewContext(r)

	c.Infof("Get the key from the URL")
	// Get the key from the URL
	keyField := r.FormValue("key")

	// Decode the key
	key, err := datastore.DecodeKey(keyField)
	if err != nil { // Couldn't decode the key
		// Do some error handling
		c.Errorf("Couldn't decode the key. Do some error handling. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the key and load it into "data"
	var data Stock
	err = datastore.Get(c, key, &data)
	if err != nil { // Couldn't find the entity
		c.Errorf("Couldn't find the entity. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	sellField, err := strconv.Atoi(r.FormValue("sell"))
	if err != nil {
		c.Errorf("Couldn't string converter. %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data.Puntos = data.Puntos - sellField

	datastore.Put(c, key, &data)
	
    b, err := json.Marshal(data)
    if err != nil {
    	c.Errorf("Json error. %v", err)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    
    c.Infof(string(b))
    fmt.Fprintln(w, string(b))

}