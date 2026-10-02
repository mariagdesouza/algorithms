package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type LogResponseData struct {
	Status      int    `json:"status"`
	StatusText  string `json:"statusText"`
	HTTPVersion string `json:"httpVersion"`
	Headers     []struct {
		Name  string `json:"name"`
		Value string `json:"value"`
	} `json:"headers"`
	Cookies []interface{} `json:"cookies"`
	Content struct {
		Size     int    `json:"size"`
		MimeType string `json:"mimeType"`
	} `json:"content"`
	RedirectURL  string `json:"redirectURL"`
	HeadersSize  int    `json:"headersSize"`
	BodySize     int    `json:"bodySize"`
	TransferSize int    `json:"_transferSize"`
}

type ReplayText struct {
	Log struct {
		Version string `json:"version"`
		Creator struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"creator"`
		Pages []struct {
			StartedDateTime time.Time `json:"startedDateTime"`
			ID              string    `json:"id"`
			Title           string    `json:"title"`
			PageTimings     struct {
				OnContentLoad float64 `json:"onContentLoad"`
				OnLoad        float64 `json:"onLoad"`
			} `json:"pageTimings"`
		} `json:"pages"`
		Entries []struct {
			StartedDateTime time.Time `json:"startedDateTime"`
			Time            float64   `json:"time"`
			Request         struct {
				Method      string `json:"method"`
				URL         string `json:"url"`
				HTTPVersion string `json:"httpVersion"`
				Headers     []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"headers"`
				QueryString []struct {
					Name  string `json:"name"`
					Value string `json:"value"`
				} `json:"queryString"`
				Cookies []struct {
					Name     string      `json:"name"`
					Value    string      `json:"value"`
					Expires  interface{} `json:"expires"`
					HTTPOnly bool        `json:"httpOnly"`
					Secure   bool        `json:"secure"`
				} `json:"cookies"`
				HeadersSize int `json:"headersSize"`
				BodySize    int `json:"bodySize"`
			} `json:"request"`
			LogResponse LogResponseData `json:"response"`
			Cache       struct {
			} `json:"cache"`
			Timings struct {
				Blocked         float64 `json:"blocked"`
				DNS             float64 `json:"dns"`
				Ssl             float64 `json:"ssl"`
				Connect         float64 `json:"connect"`
				Send            float64 `json:"send"`
				Wait            float64 `json:"wait"`
				Receive         float64 `json:"receive"`
				BlockedQueueing float64 `json:"_blocked_queueing"`
			} `json:"timings"`
			ServerIPAddress string `json:"serverIPAddress"`
			Connection      string `json:"connection,omitempty"`
			Pageref         string `json:"pageref,omitempty"`
			FromCache       string `json:"_fromCache,omitempty"`
		} `json:"entries"`
	} `json:"log"`
}

func main() {
	dat, err := ioutil.ReadFile("replaytext.json")
	if err != nil {
		log.Fatal("Error reading input file", err)
	}
	var b ReplayText
	err = json.Unmarshal(dat, &b)
	if err != nil {
		log.Fatal("json unmarshal fail", err)
	}
	for _, entry := range b.Log.Entries {
		client := &http.Client{}

		if entry.Request.Method == "GET" || entry.Request.Method == "POST" {
			req, err := http.NewRequest(entry.Request.Method, entry.Request.URL, nil)
			if err != nil {
				log.Println(entry.Request.Method, "http request error", entry.Request.URL, err)
				continue
			}
			for _, header := range entry.Request.Headers {
				req.Header.Add(header.Name, header.Value)
			}
			resp, err := client.Do(req)
			if err != nil {
				log.Println(entry.Request.Method, "http request error", entry.Request.URL, err)
				continue
			}
			/*respBody, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(entry.Request.Method, "http reading response body error", entry.Request.URL, err)
				continue
			}*/
			log.Println(entry.Request.Method, "http request done. Status:", resp.Status)
			compareResponses(resp, &entry.LogResponse)
		} else {
			fmt.Println(entry.Request.Method, entry.Request.URL)
		}

	}
}

func compareResponses(r *http.Response, lr *LogResponseData) {
	fmt.Println(r.Status, lr.Status)
	fmt.Println(r.ContentLength, lr.Content.Size)

}
