package utils

import "fmt"

func Info() {
	fmt.Println("----- gosh CLI -----")
	fmt.Println("gly <url>        // Shorten a URL")
	fmt.Println("gly look <id>    // Lookup a ID / Return the URL")
	fmt.Println("gly get <url>    // Get the original URL from a shortened one")
	fmt.Println("gly config       // Configure CLI (Removes existing config)")
	fmt.Println("gly help         // Show this help")
	fmt.Println("--------------------")
}
