// Package to output recent github repo releases, or pull requests, for supplied repository
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

// Github REST API Pulls reference: https://docs.github.com/en/rest/pulls/pulls
type GithubPulls struct {
	URL    string `json:"url"`
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
	Number int    `json:"number"`
	Title  string `json:"title"`
}

// Github REST API Releases reference: https://docs.github.com/en/rest/releases/releases
type GithubReleases struct {
	ID     int    `json:"id"`
	NodeID string `json:"node_id"`
	Name   string `json:"name"`
}

func ValidateGithubApiEndpoint(githubApiEndpoint string) bool {
	if githubApiEndpoint != "releases" && githubApiEndpoint != "pulls" {
		fmt.Println("githubApiEndpoint must be either pulls or releases. Try -e pulls")
		return false
	}
	return true
}

func getJsonData(url string) string {
	// matches non-alphanumberic characters
	request, error := http.NewRequest("GET", url, nil)
	request.Header.Set("Accept", "application/vnd.github+json")
	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		fmt.Printf("%v, %v", response.StatusCode, "FATAL: http response status code was not 200\n")
		fmt.Printf("USAGE: go run main.go -i [url] -n [int]\n")
		fmt.Printf("EXAMPLE: go run main.go -i https://api.github.com/repos/mailchimp/mc-magento2 -n 3\n")
		os.Exit(1)
	}

	body, err := ioutil.ReadAll(response.Body)
	JsonData := fmt.Sprintf("%s", body)
	if err != nil {
		panic(err)
	}

	return JsonData
}

func listGithubReleases(url string) {
	var githubBlob []GithubReleases
	err2 := json.Unmarshal([]byte(getJsonData(url)), &githubBlob)
	if err2 != nil {
		panic(err2)
	}
	regex, err := regexp.Compile("[^a-zA-Z0-9.-]+")
	if err != nil {
		fmt.Printf("regex error: %v\n", err)
	}

	for _, v := range githubBlob {
		processedString := regex.ReplaceAllString(v.Name, " ")
		fmt.Printf("%s\n", processedString)
	}
}

func listGithubPulls(url string) {
	var githubBlob []GithubPulls
	err2 := json.Unmarshal([]byte(getJsonData(url)), &githubBlob)
	if err2 != nil {
		panic(err2)
	}
	// todo: stop redefining this regest in multiple places!?
	regex, err := regexp.Compile("[^a-zA-Z0-9.-]+")
	if err != nil {
		fmt.Printf("regex error: %v\n", err)
	}

	for _, v := range githubBlob {
		processedString := regex.ReplaceAllString(v.Title, " ")
		fmt.Printf("%v, %s\n", v.Number, processedString)
	}
}

func main() {

	// COMMAND LINE switches
	var resultCount int
	var githubApiEndpoint, githubRepoUrl string
	flag.IntVar(&resultCount, "n", 3, "print NUM lines")
	flag.StringVar(&githubApiEndpoint, "e", "pulls", "Github Api endpoint. Can be 'pulls' or 'releases'")
	flag.StringVar(&githubRepoUrl, "i", "https://api.github.com/repos/mailchimp/mc-magento2", "Specify github repo URL")
	flag.Parse()

	ValidateGithubApiEndpoint(githubApiEndpoint)

	// append query string to API URL. see https://docs.github.com/en/rest/pulls/pulls
	queryString := fmt.Sprintf("%v%v", "per_page=", resultCount)

	// assemble the url
	url := fmt.Sprintf("%v/%s?%v", githubRepoUrl, githubApiEndpoint, queryString)

	switch {
	case githubApiEndpoint == "pulls":
		listGithubPulls(url)
	case githubApiEndpoint == "releases":
		listGithubReleases(url)
	}

}
