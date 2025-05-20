package netops

import (
	"fmt"
	"imageGO/internal/util"
	"io"
	"net/http"
	"os"
	"time"
)

func Fetch(request *http.Request) string {
	client := &http.Client{Timeout: 100 * time.Second}

	response, err := client.Do(request)
	util.CheckForFailure(err)

	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("FAILED: ", response.Status)
		os.Exit(1)
	}

	responseText, err := io.ReadAll(response.Body)
	util.CheckForFailure(err)

	return string(responseText)
}
