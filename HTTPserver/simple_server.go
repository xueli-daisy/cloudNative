package main
//creates a simple HTTP server listening on port 8080.upon sending a request, the server responds with a "Hello, there" message
import (
        "fmt"
        "log"
        "net/http"
        "os"
        "github.com/prometheus/client_golang/prometheus/promhttp"
        "github.com/xueli-daisy/cloudNative/HTTPserver/metrics"
        "time"
        "math/rand"
)

func main() {
        //The HandleFunc registers the handler function for thegiven URL pattern
        http.HandleFunc("/", HelloHandler)
        http.HandleFunc("/healthz", healthz)
        http.HandleFunc("/images", images)
        http.Handle("/metrics", promhttp.Handler())
        fmt.Print("Server started at port 8080")
        //The ListenAndServe listens on the TCP network addressand then handles requests on incoming connections.
        log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthz(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "working")
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
        //The handler responds to an HTTP request. It takes two parameters: the response writer and the request object.
//      fmt.Fprintf(w, "Hello, there\n")
//      fmt.Fprintf(w, "%s %s %s \n", r.Method, r.URL, r.Proto)
        for k, v := range r.Header {
                fmt.Fprintf(w, "Header field %q, Value %a\n", k, v)
        }
//      fmt.Fprintf(w, "Host = %q\n", r.Host)
        fmt.Fprintf(w, "OS_VERSION = %q\n", os.Getenv("VERSION"))
        fmt.Fprintf(w, "<h1>Welcome to cloud native</h1>")
//      fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
//      fmt.Fprintf(w, "StatusCode = %q\n", w.StatusCode)
//      fmt.Fprintf(w, "ClientAddr = %q\n", r.ClientAddr)

}

func images(w http.ResponseWriter, r *http.Request) {
        timer := metrics.NewTimer()
        defer timer.ObserveTotal()
        randInt := rand.Intn(2000)
        time.Sleep(time.Millisecond * time.Duration(randInt))
        w.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
}


