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
	validTokenEndpoint = "/api/v1/users/tokens"
)

type ValidTokenResponse struct {
	UserID string `json:"user_id"`
}

func (cl *Client) ValidToken(token string) (*ValidTokenResponse, error) {
	ctx := context.Background()
	return cl.ValidTokenWithCtx(ctx, token)
}

func (cl *Client) ValidTokenWithCtx(ctx context.Context, token string) (*ValidTokenResponse, error) {
	body := &struct {
		Token string `json:"token"`
	}{
		Token: token,
	}

	b, err := json.Marshal(body)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidTokenWithCtx json.Marshal error")
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		fmt.Sprintf("%s%s", cl.baseURL, validTokenEndpoint),
		bytes.NewBuffer(b),
	)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidTokenWithCtx http.NewRequestWithContext error")
	}

	res, err := cl.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidTokenWithCtx cl.client.Do error")
	}
	if res.StatusCode != http.StatusOK {
		return nil, errors.New(
			fmt.Sprintf("client: Client.ValidTokenWithCtx request rejected with status %d", res.StatusCode),
		)
	}

	var response ValidTokenResponse

	bodyRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidTokenWithCtx ioutil.ReadAll error")
	}

	err = json.Unmarshal(bodyRes, &response)
	if err != nil {
		return nil, errors.Wrap(err, "client: Client.ValidTokenWithCtx json.Unmarshal error")
	}
	defer res.Body.Close()

	return &response, nil
}
