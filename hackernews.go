package main

import (
	"example/hackernews/api"
	"example/hackernews/templateparse"
	"fmt"
	"net/http"
)

func main() {
    stories := api.FetchStories()
    template := templateparse.Render()


    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		// Execute the template with the data and write to ResponseWriter
        err := template.Execute(w, stories)
		if err != nil {
			fmt.Println("Error executing template:", err)
			return
		}



		// Step 4: Set the Headers (Optional)
		w.Header().Set("Content-Type", "text/plain")

		// Step 5: Send the Plate (Flush the response)
		w.(http.Flusher).Flush()
	})

	http.ListenAndServe(":8080", nil)


}
