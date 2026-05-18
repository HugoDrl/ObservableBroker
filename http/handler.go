package http
import ("net/http")

func initHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", handleHello)

	return mux
}

func handleHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello !"))
}
