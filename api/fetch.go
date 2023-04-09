package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)




func FetchStories() []Item {

    client := &http.Client{}

    // Define the REST API endpoint

    storiesUrl := "https://hacker-news.firebaseio.com/v0/topstories.json"

    // Make an HTTP GET request
    response, err := client.Get(storiesUrl)
    if err != nil {
        fmt.Println("Failed to make GET request:", err)
        return nil
    }
    defer response.Body.Close()

    // Decode the response JSON into a Go struct
    var data []int
    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        fmt.Println("Failed to decode JSON:", err)
        return nil
    }

    relevantStories := data[:30]

    var items []Item
               
    // Example with an array
    for _, value := range relevantStories{
        FetchStory(value,&items)
    } 

    return items

}


func FetchStory(storyId int, items *[]Item) {
    client := &http.Client{}

    // Define the REST API endpoint

    storyUrl := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json",storyId)
    // Make an HTTP GET request
    response, err := client.Get(storyUrl)
    if err != nil {
        fmt.Println("Failed to make GET request:", err)
        return
    }
    defer response.Body.Close()

    // Decode the response JSON into a Go struct
    var data Item
    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        fmt.Println("Failed to decode JSON:", err)
        return
    }

    *items = append(*items, data)



}
