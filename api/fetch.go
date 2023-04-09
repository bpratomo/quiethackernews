package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

func FetchStories() ([]Item, error) {
    client := &http.Client{}
    var items []Item

    storiesURL := "https://hacker-news.firebaseio.com/v0/topstories.json"
    response, err := client.Get(storiesURL)
    if err != nil {
        return nil, fmt.Errorf("failed to make GET request: %v", err)
    }
    defer response.Body.Close()

    var data []int
    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        return nil, fmt.Errorf("failed to decode JSON: %v", err)
    }

    relevantStories := data[:30]
    items = make([]Item, len(relevantStories))

    var wg sync.WaitGroup
    for index, value := range relevantStories {
        wg.Add(1)
        go func(index int, value int) {
            FetchStory(value, &items, index)
            wg.Done()
        }(index, value)
    }
    wg.Wait()

    return items, nil
}

func FetchStory(storyID int, items *[]Item, index int) {
    client := &http.Client{}
    storyURL := fmt.Sprintf("https://hacker-news.firebaseio.com/v0/item/%d.json", storyID)
    response, err := client.Get(storyURL)
    if err != nil {
        fmt.Println("Failed to make GET request:", err)
        return
    }
    defer response.Body.Close()

    var data Item
    err = json.NewDecoder(response.Body).Decode(&data)
    if err != nil {
        fmt.Println("Failed to decode JSON:", err)
        return
    }

    (*items)[index] = data

}
