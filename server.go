package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

// Message struct
type Message struct {
	Title   string `json:"title"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()
	response, err := http.Get(os.Getenv("PRIVATE_URL"))

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		temp, _ := ioutil.ReadAll(response.Body)

		var message Message
		if err := json.Unmarshal(temp, &message); err != nil {
			fmt.Println("There was an error:", err)

			e.GET("/", func(c echo.Context) error {
				return c.JSON(http.StatusOK, message)
			})
			e.Logger.Fatal(e.Start(":80"))
		}
	}
}
