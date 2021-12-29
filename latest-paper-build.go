package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type ProjectVersionRepsonse struct {
	Builds []int `json:"builds"`
}

type BuildData struct {
	Downloads struct {
		Application struct {
			Name string `json:"name"`
		} `json:"application"`
	} `json:"downloads"`
}

func main() {
	version := os.Args[1]

	var versionResponse ProjectVersionRepsonse
	GetJsonFromGetResponse("https://papermc.io/api/v2/projects/paper/versions/" + version, &versionResponse)

	latestBuild := strconv.Itoa(versionResponse.Builds[len(versionResponse.Builds) - 1])

	var buildData BuildData
	GetJsonFromGetResponse("https://papermc.io/api/v2/projects/paper/versions/" + version + "/builds/" + latestBuild, &buildData)

	jarName := buildData.Downloads.Application.Name

	fmt.Println("https://papermc.io/api/v2/projects/paper/versions/" + version + "/builds/" + latestBuild + "/downloads/" + jarName);
}

func GetJsonFromGetResponse(url string, i interface{}) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	if err := json.NewDecoder(resp.Body).Decode(&i); err != nil {
		panic(err)
	}
}
