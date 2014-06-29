/*
Copyright 2014 Ralf Rottmann. All rights reserved.
Licensed to grandcentrix GmbH.
 
About this package
 
Package rrsample is a playground for Go on App Engine
*/
package fibonacci
 
import (
	"fmt"
	"net/http"
	"time"
 
	"appengine"
	"appengine/runtime"
)
 
func init() {
	http.HandleFunc("/_ah/start", handleStart)
	http.HandleFunc("/_ah/stop", handleStop)
}
 
func handleStart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
	c := appengine.NewContext(r)
	c.Infof("pre run")
	error := runtime.RunInBackground(c, runLoop)
	c.Infof("post run")
	c.Infof("error %s", error)
}
 
func runLoop(c appengine.Context) {
	for {
		c.Infof("i would like to work on smth here...")
		time.Sleep(5 * time.Second)
	}
	return
}
 
func handleStop(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "OK")
	c := appengine.NewContext(r)
	c.Infof("post run")
}