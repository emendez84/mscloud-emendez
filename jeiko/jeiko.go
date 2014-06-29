package jeiko
 
import (
	"fmt"
	"net/http"
)
 
func init() {
	http.HandleFunc("/", handleStart)
}
 
func handleStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "jeiko.handleStart")
}