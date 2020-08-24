package ssp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"time"

	"github.com/google/uuid"
)

var client *http.Client

type BidRequest struct {
	SspName     string `json:"ssp_name"`
	RequestTime string `json:"request_time"`
	RequestId   string `json:"request_id"`
	AppId       int    `json:"app_id"`
	BidFloor    int    `json:"bidfloor"`
	UserId      string `json:"user_id"`
}

type BidResponse struct {
	RequestId *string `json:"request_id"`
	URL       *string `json:"url"`
	Price     *int    `json:"price"`
}

func runBidRequestOnce(host string, port int, floorPrice int) {
	client = &http.Client{
		Timeout: 10 * time.Second,
	}

	tm := time.Now()

	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("UUID Generated Error")
		return
	}
	reqId := u.String()
	u, err = uuid.NewRandom()
	if err != nil {
		fmt.Printf("UUID Generated Error")
		return
	}
	userId := u.String()

	appId := rand.Intn(10)

	bidRequest := BidRequest{
		SspName:     "Test SSP",
		RequestId:   reqId,
		RequestTime: tm.Format("20060102-150405.0000"),
		UserId:      userId,
		BidFloor:    floorPrice,
		AppId:       appId,
	}

	url := fmt.Sprintf("http://%s:%d/dsp/req", host, port)

	sampleJson, _ := json.Marshal(bidRequest)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(sampleJson))
	req.Header.Set("User-Agent", "Internship Mini DSP Course Tools")
	req.Header.Set("Content-Type", "application/json")

	if verbose {
		dumpReq, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("------------- Request -----------\n")
			fmt.Printf("%s\n", dumpReq)
			fmt.Printf("---------------------------------\n")
		}
	}

	res, err := client.Do(req)
	if err != nil {
		fmt.Printf("Response Error: %s\n", err)
		return
	}

	if res.StatusCode != 200 {
		fmt.Printf("Error: Status Code NOT 200 Got %d\n", res.StatusCode)
	} else {
		fmt.Printf("OK\n")
	}

	if verbose {
		dumpRes, err := httputil.DumpResponse(res, true)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("------------- Response -----------\n")
			fmt.Printf("%s\n", dumpRes)
			fmt.Printf("----------------------------------\n")
		}
	}
}

func runBidRequestFinal(host string, port int, floorPrice int) {
	client = &http.Client{
		Timeout: 10 * time.Second,
	}

	tm := time.Now()

	u, err := uuid.NewRandom()
	if err != nil {
		fmt.Printf("UUID Generated Error")
		return
	}
	reqId := u.String()
	u, err = uuid.NewRandom()
	if err != nil {
		fmt.Printf("UUID Generated Error")
		return
	}
	userId := u.String()

	appId := rand.Intn(10)

	bidRequest := BidRequest{
		SspName:     "Test SSP",
		RequestId:   reqId,
		RequestTime: tm.Format("20060102-150405.0000"),
		UserId:      userId,
		BidFloor:    floorPrice,
		AppId:       appId,
	}

	url := fmt.Sprintf("http://%s:%d/dsp/req", host, port)

	sampleJson, _ := json.Marshal(bidRequest)

	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(sampleJson))
	req.Header.Set("User-Agent", "Internship Mini DSP Course Tools")
	req.Header.Set("Content-Type", "application/json")

	if verbose {
		dumpReq, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
		} else {
			fmt.Printf("------------- Request -----------\n")
			fmt.Printf("%s\n", dumpReq)
			fmt.Printf("---------------------------------\n")
		}
	}

	tmPost := time.Now()
	res, err := client.Do(req)
	tmReturn := time.Now()
	if err != nil {
		fmt.Printf("Response Error: %s\n", err)
		return
	}

	if !(res.StatusCode == 200 || res.StatusCode == 204) {
		fmt.Printf("Error: Status Code NOT 200 Got %d", res.StatusCode)
		return
	} else {
		t := tmReturn.Sub(tmPost)
		fmt.Printf("OK Responded %d ms\n", t.Milliseconds())
	}

	length := res.ContentLength

	if verbose {
		dumpRes, err := httputil.DumpResponse(res, true)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Printf("------------- Response -----------\n")
			fmt.Printf("%s\n", dumpRes)
			fmt.Printf("----------------------------------\n")
		}
	}

	body := make([]byte, length)
	_, err = res.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Printf("内部エラー\n")
		return
	}

	if res.StatusCode == 200 {
		var pres BidResponse

		err = json.Unmarshal(body, &pres)
		if err != nil {
			fmt.Printf("JSONフォーマットエラー\n")
			return
		}

		if pres.RequestId == nil || pres.URL == nil || pres.Price == nil {
			fmt.Printf("JSONフォーマットエラー\n")
			return
		}

		if *pres.RequestId != reqId {
			fmt.Printf("Request ID is ignore expected id is %s\n", reqId)
		}

		fmt.Printf("RequestID: %s\n", *pres.RequestId)
		fmt.Printf("URL      : %s\n", *pres.URL)
		fmt.Printf("Price    : %d\n", *pres.Price)
	} else {
		fmt.Printf("No Contents\n")
	}

}
