package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

var (
	validApiKeyEndpoint = "/api/v1/users/api-key"
)

type ValidApiKeyResponse struct {
	UserID string `json:"user_id"`
}

type ValidApiKeyRequest struct {
	ApiKey string `json:"api_key"`
}

func (cl *Client) ValidApiKey(token string) (*ValidApiKeyResponse, error) {
	ctx := context.Background()
	return cl.ValidApiKeyWithCtx(ctx, token)
}

func (cl *Client) ValidApiKeyWithCtx(ctx context.Context, apiKey string) (*ValidApiKeyResponse, error) {
	body := &ValidApiKeyRequest{
		ApiKey: apiKey,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidApiKeyWithCtx json.Marshal error")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", cl.baseURL, validTokenEndpoint),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidApiKeyWithCtx http.NewRequestWithContext error")
	}

	res, err := cl.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidApiKeyWithCtx cl.client.Do error")
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(
			fmt.Sprintf("client: Client.ValidApiKeyWithCtx request rejected with status %d", res.StatusCode),
		)
	}

	var response ValidApiKeyResponse

	bodyRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidApiKeyWithCtx ioutil.ReadAll error")
	}

	err = json.Unmarshal(bodyRes, &response)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidApiKeyWithCtx json.Unmarshal error")
	}
	defer res.Body.Close()

	return &response, nil
}
