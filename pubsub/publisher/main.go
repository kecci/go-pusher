package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.AllowAll().Handler)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})

	// Message
	r.Post("/api/message", PushMessage)

	http.ListenAndServe(":9090", r)
}

// PushMessage push
func PushMessage(w http.ResponseWriter, r *http.Request) {
	pusherClient := pusher.Client{
		AppID:   "1492736",
		Key:     "91fcd5f8f73e694c6a38",
		Secret:  "63ff1c71c1daa8949714",
		Cluster: "ap1",
		Secure:  true,
	}

	data := map[string]string{"message": "hello world"}
	// Channel: Service (Payment Service, Order Service, Product Service)
	// Event: Action (Notif Payment, Cancel Order, Request Product)
	pusherClient.Trigger("my-channel", "my-event", data)
	w.Write([]byte("success"))
}
