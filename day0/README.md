## Intro
Today, we will start with creating a basic web app

## Step 1: Create a module
1. type `go mod init github.com/SimpleNotesApp`
2. Three commands that are used often for `go mod command`
- `go mod tidy`: add missing or remove unused modules
- `go mod vendor`: add vendored copy of dependencies
- `go mod download`: download modules to local cache

## Step 2: create a basic web app
There are 3 components important for a basic golang web app
1. a web server listening for incoming requests using net/http package.
2. router or servemux that registers a handler based on a pattern
3. a handler that handles the request and writes a response back to the user

## Step 3: a more implicit way of a basic web app
```go
package main
func main(){
	// HandleFunc registers the handler function for the given pattern in defaultServeMux
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
	// ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections. Accepted connections are configured to enable TCP keep-alives.
	// The handler is typically nil, in which case DefaultServeMux is used.
	// ListenAndServe always returns a non-nil error.
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```
## Notes:
Earlier we explored that the Handler interface is what we need to implement in order to make a server. 
Typically we do that by creating a struct and make it implement the interface by implementing its own ServeHTTP method.
However the use-case for structs is for holding data but currently we have no state, so it doesn't feel right to be creating one.
**The HandlerFunc type is an adapter to allow the use of ordinary functions as HTTP handlers.** 
**If f is a function with the appropriate signature, HandlerFunc(f) is a Handler that calls f.**


