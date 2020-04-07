package nextcloud

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	app string
	endpoint string
	method string
	body string
}

func (nc Client) Request() string {

	buffer := bytes.NewBufferString(nc.body)

	req, err := http.NewRequest(nc.method, "https://cloud.nextcloud.org/ocs/v2.php/apps/"+ nc.app +"/api/v1/" + nc.endpoint, buffer)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Add("OCS-APIRequest", "true")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth("XXXXX", "XXXXXX")

	client := &http.Client{Timeout: time.Second * 5}

	resp, err := client.Do(req)
	if err != nil {
		return "{'message': 'Error reading response'}"
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Error reading body. ", err)
		return "{'message': 'Error reading body'}"
	}

	xml := strings.NewReader(string(body))
	json, err := xml2json.Convert(xml)
	if err != nil {
		return "{'message': 'Could not convert XML to JSON'}"
	}

	return json.String()
}
