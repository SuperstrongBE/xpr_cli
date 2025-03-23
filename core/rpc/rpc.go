package rpc

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const TN_EP = "https://testnet.rockerone.io"
const MN_EP = "https://api.rockerone.io"

type Data struct {
	AccountName string `json:"account_name"`
}

func GetAbi(contract string, testnet bool) string {

	service := "/v1/chain/get_abi"
	baseAPI := MN_EP
	if testnet {
		baseAPI = TN_EP
	}
	url := baseAPI + service
	postBody, _ := json.Marshal(map[string]string{
		"account_name": contract,
	})
	log.Print(string(postBody))
	postBodyBytes := bytes.NewBuffer(postBody)

	resp, err := http.Post(url, "application/json", postBodyBytes)
	if err != nil {
		log.Fatalf("RPC error occured %v", err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body)

}
