package main

import (
	"fmt"
	"net/http"
	"os"
	"io/ioutil"
	"encoding/json"
	"flag"
)
type Giphka struct {
	Url string `json:"url"`
	Name string `json:"title"`
}

type GiphsApiResponse struct {
	GiphsList [] Giphka `json:"data"`
}

const giphyURL string = "http://api.giphy.com/v1/gifs/trending?api_key=WsJyksJKHPWzjL7dWwhRGx8PXp6VJkZV&limit="
const fileName string = "result.txt"

func getGiphs(body []byte) (*GiphsApiResponse, error){
	var giphList = new(GiphsApiResponse)
	err := json.Unmarshal(body, giphList)
	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	}
	return giphList,err
}

func main() {


	limitPtr := flag.Int("limit", 5, "Limit of giphs to be fetched")
	flag.Parse()
	fmt.Printf("Flag limit option value: %v \n", *limitPtr)
	formattedUrl := fmt.Sprintf("%s%v", giphyURL, *limitPtr)
	fmt.Println(formattedUrl)
	response, err := http.Get(formattedUrl)
	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		contentsBytes, err := ioutil.ReadAll(response.Body)
		if err != nil{
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		//contentString := string(contentsBytes)
		//fmt.Println(contentString)
		giphsApi, err := getGiphs(contentsBytes)
		giphs := giphsApi.GiphsList
		file, err := os.Create(fileName)
		for _, giph := range giphs{
			fmt.Printf("Url: %s, Name: %s \n", giph.Url, giph.Name)
			if err != nil {
				fmt.Printf("%s", err)
				os.Exit(1)
			}
			fmt.Fprintf(file, "Url: %s, Name: %s \n", giph.Url, giph.Name)
		}
		defer file.Close()

	}

}
