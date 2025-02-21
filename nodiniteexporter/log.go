package nodiniteexporter

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.opentelemetry.io/collector/pdata/plog"
)

// No default function for this. It must be implemented
// Note: You can change the function name if you like
func pushLogs(
	ctx context.Context,
	td plog.Logs,
) (err error) {

	testurl := "https://63f32cb6fe3b595e2edc301f.mockapi.io/LogEvent"

	data := map[string]int{
		"id": 1,
	}

	// Convert the data map to a JSON string
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", testurl, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Set the headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_ACCESS_TOKEN") // If authentication is required

	// Create an HTTP client and send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode == http.StatusOK {
		// Successfully received a response
		var responseData map[string]interface{}
		if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
			fmt.Println("Error decoding response:", err)
		}
		fmt.Println("Response:", responseData)
	} else {
		// Handle errors
		fmt.Println("Error:", resp.StatusCode)
	}

	return nil
}
