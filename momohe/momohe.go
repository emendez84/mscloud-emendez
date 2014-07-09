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
		fmt.Println("Specific Stock Query Web Service")
	}

	// Get the key and load it into "data"
	var data Stock

	if item, err := memcache.Get(c, keyField); err == memcache.ErrCacheMiss {
		c.Infof("Stock not in the cache. Retrieving info from DataStore")
		// Decode the key
		key, err := datastore.DecodeKey(keyField)
		if err != nil { // Couldn't decode the key
			// Do some error handling
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = datastore.Get(c, key, &data)
		if err != nil { // Couldn't find the entity
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
        fmt.Println(err)
        return
    }
    
    w.Header().Set("Content-Type", "application/json")
    
    fmt.Fprintln(w, string(b))

}
