package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/DanielFillol/API_Caller/models"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var APIkey string

//APIRequest uses Call to make a request to the API. Here is our main worker
func APIRequest(name string, email string, password string) (models.WriteStruct, error) {
	res, err := Call("http://localhost:8080/metaphone/"+name, "GET", email, password)
	if err != nil {
		return models.WriteStruct{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return models.WriteStruct{}, err
	}

	var response models.ResponseAPI
	err = json.Unmarshal(body, &response)
	if err != nil {
		return models.WriteStruct{}, err
	}

	return models.WriteStruct{
		SearchName:     name,
		ID:             strconv.Itoa(response.ID),
		CreatedAt:      response.CreatedAt.String(),
		UpdatedAt:      response.UpdatedAt.String(),
		DeletedAt:      response.DeletedAt.String(),
		Name:           response.Name,
		Classification: response.Classification,
		Metaphone:      response.Metaphone,
		NameVariations: response.NameVariations,
		Err:            nil,
	}, nil
}

//Call creates a http.Client passing KEY on the header
func Call(url, method string, email string, password string) (*http.Response, error) {
	client := &http.Client{Timeout: time.Second * 10}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Token", APIkey)
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == http.StatusUnauthorized {
		c, err := login(email, password)
		if err != nil {
			return nil, err
		}
		APIkey = c.Value

		rq, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, err
		}
		rq.Header.Add("Token", APIkey)

		r, err := client.Do(rq)
		if err != nil {
			return nil, err
		}

		return r, nil
	}

	return response, nil
}

//login on the API returning the cookie
func login(email string, password string) (*http.Cookie, error) {
	url := "http://localhost:8080/login"

	// Define the request body as a JSON object
	data := map[string]string{"email": email, "password": password}
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Make the POST request with the JSON payload
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Check if the response status code is 200
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("login failed with status code %d", resp.StatusCode)
	}

	// Return the cookie from the response
	cookies := resp.Cookies()
	if len(cookies) == 0 {
		return nil, fmt.Errorf("no cookie in response")
	}
	return cookies[0], nil
}
