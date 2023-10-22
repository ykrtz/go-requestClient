package requestClient
import (
	"io/ioutil"
	"net/http"

)

type Category struct {
	filter string
}

func SendRequest(url string) ([]byte, error) {
	// Create a POST request with custom headers
	client := &http.Client{}
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}

	// Add custom headers
	req.Header.Set("accept", "text/html,*/*")
	req.Header.Set("accept-language", "de-DE,de;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("content-type", "application/x-www-form-urlencoded")
	req.Header.Set("device-memory", "8")
	req.Header.Set("downlink", "10")
	req.Header.Set("dpr", "1.125")
	req.Header.Set("ect", "4g")
	req.Header.Set("rtt", "50")
	req.Header.Set("sec-ch-device-memory", "8")
	req.Header.Set("sec-ch-dpr", "1.125")
	req.Header.Set("sec-ch-ua", "\"Google Chrome\";v=\"113\", \"Chromium\";v=\"113\", \"Not-A.Brand\";v=\"24\"")
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Set("sec-ch-ua-platform-version", "\"10.0.0\"")
	req.Header.Set("sec-ch-viewport-width", "802")
	req.Header.Set("sec-fetch-dest", "empty")
	req.Header.Set("sec-fetch-mode", "cors")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("viewport-width", "802")
	req.Header.Set("x-requested-with", "XMLHttpRequest")

	// Send the request
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
