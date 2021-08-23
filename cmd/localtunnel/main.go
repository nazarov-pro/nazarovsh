package main

import (
	"fmt"
	"github.com/NoahShen/gotunnelme/src/gotunnelme"
	"os"
	"strconv"
)

func main() {
	port := os.Getenv("PORT")
	subdomain := os.Getenv("SUBDOMAIN")

	fmt.Printf("PORT: %s, SUBDOMAIN: %s\n", port, subdomain)

	i, err := strconv.Atoi(port)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	t := gotunnelme.NewTunnel()
	url, err := t.GetUrl(subdomain)
	if err != nil {
		panic(err)
	}
	print(url)
	err = t.CreateTunnel(i)
	if err != nil {
		panic(err)
	}
	t.StopTunnel()
}
