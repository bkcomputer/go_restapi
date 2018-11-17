package main

import (
	
)

// Book struct (Model)

// Main function
func main() {
	// Init router
	
	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Request sample
// {
// 	"isbn":"4545454",
// 	"title":"Book Three",
// 	"author":{"firstname":"Harry","lastname":"White"}
// }
