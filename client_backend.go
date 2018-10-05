package main
 
import (
    "fmt"
    "log"
    "net/http"
    "io/ioutil"
    "os"
    //"net/http/httputil"
)

const INNER_PAGE = "form.html"
const MAIN_PAGE = "home.html"

// logs the http request on the server currently it prints to the command line
func logServer(r *http.Request) {
     log.Printf("\tIP:%s\t Method:%s\t URL:%s\n", r.RemoteAddr, r.Method, r.URL)
}

// the process for when the inner page is asked for
func form(w http.ResponseWriter, r *http.Request) {
    logServer(r)

    // makes sure the url is correct
    if r.URL.Path != ("/" + INNER_PAGE) {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    // switch statement based on the client message request
    switch r.Method {
    case "GET":
    	 // sends the inner html page to the user
         http.ServeFile(w, r, INNER_PAGE)
    default:
        fmt.Fprintf(w, "Sorry, only GET method is supported.")
    }
}

func hello(w http.ResponseWriter, r *http.Request) {
    logServer(r)

    // makes sure the url is correct
    if r.URL.Path != ("/" + MAIN_PAGE) {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    // switch statement based on the client message request
    switch r.Method {
    case "GET":     
         http.ServeFile(w, r, MAIN_PAGE)
    case "POST":
    	 // converts the http body into bytes
         body, err := ioutil.ReadAll(r.Body)
         if err != nil {
            fmt.Println(err)
         }

	 // converts the body(bytes) into a readable string
         bodyString := string(body)
	 // creates a blank inner page
	 file, err := os.Create(INNER_PAGE)
    	 if err != nil {
            log.Fatal("Cannot create file", err)
    	 }
    	 defer file.Close()

	 // writes the html body to the file
    	 fmt.Fprintf(file, bodyString)

    default:
         fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
    }
}
 
func main() {
    // waits for the user to go to "website"/form.html and calls form when they go there
    http.HandleFunc("/" + INNER_PAGE, form)

    // waits for the user to go to "website"/ and calls hello when they go there
    http.HandleFunc("/" + MAIN_PAGE, hello)
    http.Handle("/", http.FileServer(http.Dir("css/")))

    fmt.Printf("Starting server...\n")

    // listens on port 8080 for a http request
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}