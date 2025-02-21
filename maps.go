package main

import "fmt"

func main() {
	websites := map[string]string{
		"google":              "https://google.com",
		"amazon web services": "https://aws.com",
	}

	fmt.Println(websites["google"])
	websites["linkedIn"] = "https://linkedin.com"
	fmt.Println(websites)

	delete(websites, "google")
	fmt.Println(websites)
}
