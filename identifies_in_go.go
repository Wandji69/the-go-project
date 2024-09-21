package main

import "fmt"

//Constants identified MaxRetries
const MaxRetries = 3

//struct Data type GithubUser
type GithubUser struct {
	Login	string `json:"login"`
	Name	string `json:"name"`
	PublicRepos int `josn:"public_repost"`

}

// Package-Private Variable identifier _internalVariable
var _internalVariable int


// Export Function identifier CalculateSum
func CalculateSum(a, b int) int {
	return a + b
}

func main()  {
	// Local variable indentifier count
	countt := 10

	//printing Using the identifiers
	fmt.Println("Count:", count)
	fmt.Println("Max Retries:", MaxRetries)

	//Using the GithubUser Struct type as for identifer person
	person:= GithubUser{Login: "John", Name: "Doe", PublicRepos: 7}

	fmt.Println("Perspon:" person)
}