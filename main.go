package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type (
	Giphka struct {
		Url  string `json:"url"`
		Name string `json:"title"`
	}

	GiphsApiResponse struct {
		GiphsList []Giphka `json:"data"`
	}
)

const (
	// приймай api ключ як аргумент з командного рядку
	giphyURL string = "http://api.giphy.com/v1/gifs/trending?api_key=WsJyksJKHPWzjL7dWwhRGx8PXp6VJkZV&limit="
	fileName string = "result.txt"
)

// Створимо змінну в heap
var limitPtr = flag.Int("limit", 5, "Limit of giphs to be fetched")

func getGiphs(body []byte) (*GiphsApiResponse, error) {
	var giphList = new(GiphsApiResponse)
	//another way: giphList := &GiphsApiResponse{}
	err := json.Unmarshal(body, giphList)
	if err != nil {
		// є рекомендації по роботі з обробленням помилок,
		// згідно з ними, ми опрацьовуємо помилки "зверху",
		// в нижній частині програми ми просто передаємо її на верх
		return nil, err
	}
	return giphList, nil
}

func main() {
	flag.Parse()
	fmt.Printf("Flag limit option value: %v \n", *limitPtr)
	formattedUrl := fmt.Sprintf("%s%v", giphyURL, *limitPtr)
	fmt.Println(formattedUrl)
	response, err := http.Get(formattedUrl)
	if err != nil {
		fmt.Println("%s", err)
		os.Exit(1)
	}
	// надлишковий else:
	//якщо err != nil - програма закінчить виконання

	defer response.Body.Close()
	contentsBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	//contentString := string(contentsBytes)
	//fmt.Println(contentString)
	giphsApi, err := getGiphs(contentsBytes)
	// !!! main функція в даному випадку є "найвищою" точкою програми,
	// і тут ми опрацюєм помилки, які прийшли "знизу"
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	giphs := giphsApi.GiphsList
	file, err := os.Create(fileName)
	// !!!!!завжди обрацьовуєм помилки
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	// краще викликати defer поруч з місцем контексту
	defer file.Close()

	for _, giph := range giphs {
		fmt.Printf("Url: %s, Name: %s \n", giph.Url, giph.Name)
		if err != nil {
			fmt.Printf("%s", err)
			os.Exit(1)
		}
		fmt.Fprintf(file, "Url: %s, Name: %s \n", giph.Url, giph.Name)
	}

}
