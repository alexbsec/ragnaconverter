package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	http2 "github.com/alexbsec/ragconverter/core/http"
	"github.com/alexbsec/ragconverter/types"
)

const ItemEndpoint = "/api/database/Item"

type DivineRequester struct {
	host   string
	client http.Client
	apiKey string
}

func NewDivineRequester(host, apiKey string) DivineRequester {
	return DivineRequester{
		host:   host,
		client: http.Client{},
		apiKey: apiKey,
	}
}

func (dr DivineRequester) GetItem(itemRequest types.ItemRequest) (types.ItemResponse, error) {
	var response types.ItemResponse

	queryParams := map[string]string{
		"id":     strconv.Itoa(itemRequest.Id),
		"apiKey": itemRequest.ApiKey,
	}

	uri := fmt.Sprintf("%s/%s", dr.host, ItemEndpoint)

	requester := http2.NewRequest(dr.client, http2.GET, uri, nil, queryParams, nil)
	bytes, statusCode, err := requester.Send()
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		return response, err
	}

	response.StatusCode = statusCode
	if err := json.Unmarshal(bytes, &response); err != nil {
		response.StatusCode = http.StatusInternalServerError
		return response, err
	}

	return response, nil
}
