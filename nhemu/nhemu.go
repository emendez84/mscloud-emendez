package nhemu
 
import (
	"fmt"
	"net/http"
)
 
func init() {
	http.HandleFunc("/nhemu/", handleStart)
    http.HandleFunc("/nhemu/check", checkSells)
}
 
func handleStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "nhemu.handleStart")
}

func checkSells(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "checkSells")
}