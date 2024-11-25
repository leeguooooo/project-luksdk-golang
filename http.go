package luksdk

import (
	"bytes"
	"encoding/json"
	"net/http"
)

func request(cli *http.Client, method, url string, body any, response any) (err error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := cli.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	// 读取响应
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return err
	}

	return nil
}
