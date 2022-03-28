package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptrace"
	"time"
)

type App struct {
	client *http.Client
}

func appWithHttpClient() *App {
	client := &http.Client{
		Timeout: time.Second * 10,
		Transport: &http.Transport{
			MaxIdleConns:        10,
			IdleConnTimeout:     30 * time.Second,
			MaxIdleConnsPerHost: 10,
		},
	}
	return &App{client: client}
}

func main() {
	app := appWithHttpClient()
	// app.clientWithTrace()
	result, _ := app.postNewItem()
	fmt.Println(string(result))
}

func (a *App) getFromAPI() ([]byte, error) {
	resp, err := a.client.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *App) getFromAPIWithClient() ([]byte, error) {
	req, err := http.NewRequest("GET", "https://jsonplaceholder.typicode.com/posts", nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *App) getWithQueryString() ([]byte, error) {

	// id := "1"
	// params := "postId=" + url.QueryEscape(id)
	// path := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/comments?%s", params)

	// üstteki kod parçası yerine url'e ekstra bir şey eklemek istersek (mesela yukarıdaki örnekte comments?%s yerine aşağıda path'e direk comments olarak vereceğiz)

	req, err := http.NewRequest(http.MethodGet, "https://jsonplaceholder.typicode.com/comments", nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("postId", "1")
	req.URL.RawQuery = q.Encode()

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *App) signInWithBearerToken() ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, "http://httpbin.org/bearer", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InVzZXIxIiwiZXhwIjoxNTQ3OTc0MDgyfQ.2Ye5_w1z3zpD4dSGdRp3s98ZipCNQqmsHRB9vioOx54"))

	resp, err := a.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func (a *App) clientWithTrace() {
	clientTrace := &httptrace.ClientTrace{
		GotConn: func(info httptrace.GotConnInfo) {
			log.Printf("conn was reused: %t", info.Reused)
		},
	}
	traceCtx := httptrace.WithClientTrace(context.Background(), clientTrace)

	req, err := http.NewRequestWithContext(traceCtx, http.MethodGet, "http://example.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := io.Copy(ioutil.Discard, res.Body); err != nil {
		log.Fatal(err)
	}

	res.Body.Close()

	req, err = http.NewRequestWithContext(traceCtx, http.MethodGet, "http://example.com", nil)

	if err != nil {
		log.Fatal(err)
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) postNewItem() ([]byte, error) {
	values := map[string]string{
		"name": "Ahmet",
		"age":  "21",
	}

	bodyData, err := json.Marshal(values)
	if err != nil {
		return nil, err
	}

	resp, err := a.client.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(bodyData))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil

}
