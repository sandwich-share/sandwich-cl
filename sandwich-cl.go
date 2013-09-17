package main

import(
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type FileOrDir struct {
	Type int
	Name string
}

func fetchPeerFiles(peerIP string, path string, start string, step string) {
	requestURL := fmt.Sprintf("http://%s/peer?peer=%s&path=%s&start=%s&step=%s",
	"localhost:9001", peerIP, path, start, step)
	res, _ := http.Get(requestURL)
	body, _ := ioutil.ReadAll(res.Body)
	var data []FileOrDir

	json.Unmarshal(body, &data)
	
	for _, file := range data {
		var fileType string
		if file.Type == 1 {
			fileType = "File"
		} else {
			fileType = "Dir"
		}
		fmt.Printf("%s: %s\n", fileType, file.Name)
	}
}

func main() {
	fetchPeerFiles("107.21.226.221", "", "0", "10")
}
