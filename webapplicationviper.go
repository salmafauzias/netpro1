package main

import (
	"fmt"
	"html"
	"net/http"

	"github.com/spf13/viper"
)

func main() {
	// Lokasi config
	viper.SetConfigFile("./config/config.json")

	// Jika file config tidak ada, show error
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Error config file!, &s", err)
	}

	// Set route ketika web diakses
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %q", html.EscapeString(r.URL.Path))
	})

	// Set port
	http.ListenAndServe(":"+viper.GetString("server.port"), nil)
}
