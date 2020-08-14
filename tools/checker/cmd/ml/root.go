package ml

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/spf13/cobra"
)

var (
	host = ""
	port = 0
)

type PredictRequest struct {
	Ads    []string `json:"ads"`
	UserId string   `json:"user_id"`
}

var PredictResponse map[string]int

func NewRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ml",
		Short: "ML (Marchine Lerning) Tools",
		Run: func(cmd *cobra.Command, args []string) {
			listenHost := fmt.Sprintf("%s:%d", host, port)
			fmt.Printf("Listen on: %s\n", listenHost)
			http.HandleFunc("/predict", checkHandler)
			http.ListenAndServe(listenHost, nil)
		},
	}

	rootCmd.SetArgs(args)

	rootCmd.PersistentFlags().StringVar(&host, "host", "localhost", "Host Name")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8079, "Port Number")

	return rootCmd
}

func checkHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Printf("Request Methodが、 '%s' でした。POSTである必要があります。\n", r.Method)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "METHODが違います。")
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		fmt.Printf("Content Typeが、 '%s' でした。 'application/json'である必要があります。\n", r.Header.Get("Content-Type"))
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Content-Typeが違います。")
		return
	}

	if r.Header.Get("Content-Length") == "" {
		fmt.Printf("Content Lengthがありません。 'Content-Length'がある必要があります。\n")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Content-Lengthがありません。")
		return
	}
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		fmt.Printf("Content Lengthのパースに失敗しました。 (%s)\n", err)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Content-Typeが違います。")
		return
	}

	body := make([]byte, length)
	_, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		fmt.Printf("JSONの読み込みに失敗しています。POSTが正常にされているか、確認してください。\n")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var preq PredictRequest

	err = json.Unmarshal(body, &preq)
	if err != nil {
		fmt.Printf("JSONの読み込みに失敗しています。JSONのフォーマットを確認してください。(%s)\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(preq.Ads) == 0 {
		fmt.Printf("リクエスト広告がありません。1つ以上の広告を含めてください。\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if preq.UserId == "" {
		fmt.Printf("ユーザーIDがありません。ユーザーIDは必須です。\n")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	PredictResponse = make(map[string]int)

	for _, ad := range preq.Ads {
		n := rand.Intn(100)
		PredictResponse[ad] = n
	}

	jsonRaw, _ := json.Marshal(PredictResponse)
	w.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(w, bytes.NewBuffer(jsonRaw).String())
}
