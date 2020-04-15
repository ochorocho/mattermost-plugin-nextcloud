package nextcloud

import (
	"bytes"
	"github.com/basgys/goxml2json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	Endpoint string
	Method   string
	Body     string
}

func (nc Client) Request() string {

	buffer := bytes.NewBufferString(nc.Body)

	// https://cloud.knallimall.org/ocs/v2.php/
	// 								cloud/capabilities
	//								apps/spreed/api/v1/room

	req, err := http.NewRequest(nc.Method, "https://cloud.knallimall.org/ocs/v2.php/"+nc.Endpoint, buffer)
	if err != nil {
		log.Fatal("Error reading request. ", err)
	}

	req.Header.Add("OCS-APIRequest", "true")
	req.Header.Add("Content-Type", "application/json")
	req.SetBasicAuth("admin", "XXXXXXX")

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
