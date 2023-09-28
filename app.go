package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

const (
	BASEURL = "http://localhost:8080"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleDefault).Methods("GET")
	r.HandleFunc("/yay", handleOther).Methods("GET")

	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/"))))

	http.ListenAndServe(":8080", r)
}

func handleDefault(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		fmt.Fprintf(resp, `<!DOCTYPE html><html lang="en"><head><script src="https://unpkg.com/htmx.org@1.9.2"></script></head><body>
		<script>
		document.body.addEventListener('htmx', (event)={
			if(event.detail.xhr.status == 404){
				alert("page not found")
			}
		});
		</script>
		<button hx-get="%v/yay" hx-swap="beforeend" hx-target="next .list"> click me bitch
			<img class="htmx-indicator" width="20" height="20" src="/assets/spin.gif"/>
		</button>
		<button hx-get="%v/yay2" hx-swap="beforeend" hx-target="next .list"> click me bitch 2 </button>
		<button hx-on="click: alert('Clicked!!!')"> click me bitch3 </button>
		<ul class="list">
			<li>Item1</li>
		</ul>
		</body>
		<style>
		.htmx-indicator{
			opacity:0;
			transition: opacity 500ms ease-in;
		}
		.htmx-request .htmx-indicator{
			opacity:1
		}
		.htmx-request.htmx-indicator{
			opacity:1
		}
		</style>
		</html>`, BASEURL, BASEURL)
	}

}
func handleOther(resp http.ResponseWriter, req *http.Request) {
	if req.Method == "GET" {
		time.Sleep(4 * time.Second) // DO NOT REMOVE MY SLEEP SO YOU CAN WATCH THE SPINNER
		fmt.Fprintf(resp, `<li>hmm</li>`)
	}
	if req.Method == "POST" {
		fmt.Fprintf(resp, "<p>POST</p>")
	}
}
