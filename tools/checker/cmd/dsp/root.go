package dsp

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	host = "localhost"
	port = 8080
)

type BidRequest struct {
	SspName     string `json:"ssp_name"`
	RequestTime string `json:"request_time"`
	RequestId   string `json:"request_id"`
	AppId       int    `json:"app_id"`
	BidFloor    int    `json:"bidfloor"`
	UserId      string `json:"user_id"`
}

type BidResponse struct {
	RequestId string `json:"request_id"`
	Url       string `json:"url"`
	Price     int    `json:"price"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		fmt.Printf("Content Type is ignore: %s\n", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		fmt.Printf("Content Length is missing or Invalid\n")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	leng, err := r.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Printf("Reading Body Error")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if leng != length {
		fmt.Printf("length mismatch %d != %d", length, leng)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var brq BidRequest

	err = json.Unmarshal(body, &brq)

	b := BidResponse{
		RequestId: brq.RequestId,
		Price:     100,
		Url:       "http://example.jp/img/211",
	}

	w.Header().Set("Content-Type", "application/json")

	jsonRaw, _ := json.Marshal(b)
	fmt.Fprintf(w, bytes.NewBuffer(jsonRaw).String())
}

func NewRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "dsp",
		Short: "DSP Tools",
		RunE: func(cmd *cobra.Command, args []string) error {
			return errors.New("No Command")
		},
	}

	rootCmd.SetArgs(args)
	rootCmd.AddCommand(newServerCmd())

	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "Host Name")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "Port Number")

	return rootCmd
}

func newServerCmd() *cobra.Command {
	listenHost := fmt.Sprintf("%s:%d", host, port)
	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "Normal Server Mode",
		Run: func(cmd *cobra.Command, args []string) {
			http.HandleFunc("/dsp/req", handler)
			http.ListenAndServe(listenHost, nil)
		},
	}

	return serverCmd
}
