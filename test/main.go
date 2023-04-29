package main

import (
	"log"

	jurl "github.com/jbrekelmans/go-url"
)

func main() {
	u, err := jurl.ValidateURL("https://bla@github.com:443//?foo=bar#xyz", jurl.ValidateURLOptions{
		Abs:                      jurl.NewBool(true),  // Require the URL to be absolute
		AllowedSchemes:           []string{"https"},   // Only allow scheme https
		NormalizePort:            jurl.NewBool(false), // Remove port if it is the default port for the URL.
		StripFragment:            true,
		StripQuery:               true, // Remove query string.
		StripPathTrailingSlashes: true,
		StripUser:                true,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%v", u)
}
