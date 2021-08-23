package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/nazarov-pro/nazarovsh/pkg/tgbot"
)

func main() {
	port := os.Getenv("PORT")
	fmt.Printf("Application starting on PORT: %s\n", port)
	http.ListenAndServe(":" + port, http.HandlerFunc(tgbot.Handler))
}
