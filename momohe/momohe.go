package momohe
 
import (
	"fmt"
	"encoding/json"	
	"net/http"
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
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

	// Get the key from the URL
	keyField := r.FormValue("key")

	if keyField == "" {
		c.Infof("Consulta especifica del stock")
	}

	// Get the key and load it into "data"
	var data Stock

	if item, err := memcache.Get(c, keyField); err == memcache.ErrCacheMiss {
		c.Infof("El stock no esta en la cache. Consultando información del datastore")
		// Decode the key
		key, err := datastore.DecodeKey(keyField)
		if err != nil { // Couldn't decode the key
			// Do some error handling
			c.Errorf("Error al decodificar los stocks desde el datastore. %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = datastore.Get(c, key, &data)
		if err != nil { // Couldn't find the entity
			c.Errorf("Error al decodificar los stocks desde el datastore. %v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		newCacheItem := &memcache.Item{
			Key:   keyField,
			Object: data,
		}

		memcache.Gob.Set(c, newCacheItem)
	} else if err != nil {
		c.Errorf("error getting stock from cache: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		c.Infof("Stock found in cache. %s", item.Key)
		memcache.Gob.Get(c, keyField, &data)
	}

    b, err := json.Marshal(data)
    if err != nil {
    	c.Errorf("Json error. %v", err)
        fmt.Println(err)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    
    c.Infof(string(b))
    fmt.Fprintln(w, string(b))

}
