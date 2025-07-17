package luksdk

import (
	"encoding/json"
	"time"

	"github.com/CFGameTech/project-luksdk-golang/luksdk/luksdkmodels/apimodels"
)

func newApis(client *LukSDK) *Apis {
	return &Apis{
		client: client,
	}
}

type Apis struct {
	client *LukSDK
}

func (a *Apis) GetGameServiceList(req apimodels.GetGameServiceListRequest) (apimodels.GetGameServiceListResponse, error) {
	var response apimodels.GetGameServiceListResponse

	if req.AppId == 0 {
		req.AppId = a.client.config.AppId
	}
	if req.Sign == "" {
		req.Sign = NewSignature(a.client.config.AppSecret, req)
	}
	if req.Timestamp == nil {
		now := time.Now().Unix()
		req.Timestamp = &now
	}
	result, err := a.client.config.HttpClient.NewRequest().
		SetBody(req).
		Post("/sdk/get_game_service_list")
	if err != nil {
		return response, err
	}

	return response, json.Unmarshal(result.Bytes(), &response)
}

func (a *Apis) QueryNotifyEvent(req apimodels.QueryNotifyEventRequest) (apimodels.QueryNotifyEventResponse, error) {
	var response apimodels.QueryNotifyEventResponse

	if req.AppId == 0 {
		req.AppId = a.client.config.AppId
	}
	if req.Sign == "" {
		req.Sign = NewSignature(a.client.config.AppSecret, req)
	}
	if req.Timestamp == nil {
		now := time.Now().Unix()
		req.Timestamp = &now
	}
	result, err := a.client.config.HttpClient.NewRequest().
		SetBody(req).
		Post("/sdk/query_notify_event")
	if err != nil {
		return response, err
	}

	return response, json.Unmarshal(result.Bytes(), &response)
}

func (a *Apis) QueryOrder(req apimodels.QueryOrderRequest) (apimodels.QueryOrderResponse, error) {
	var response apimodels.QueryOrderResponse

	if req.AppId == 0 {
		req.AppId = a.client.config.AppId
	}
	if req.Sign == "" {
		req.Sign = NewSignature(a.client.config.AppSecret, req)
	}
	if req.Timestamp == nil {
		now := time.Now().Unix()
		req.Timestamp = &now
	}
	result, err := a.client.config.HttpClient.NewRequest().
		SetBody(req).
		Post("/sdk/query_order")
	if err != nil {
		return response, err
	}

	return response, json.Unmarshal(result.Bytes(), &response)
}

func (a *Apis) PublishControlEvent(req apimodels.PublishControlEventRequest) (apimodels.PublishControlEventResponse, error) {
	var response apimodels.PublishControlEventResponse

	if req.AppId == 0 {
		req.AppId = a.client.config.AppId
	}
	if req.Sign == "" {
		req.Sign = NewSignature(a.client.config.AppSecret, req)
	}
	if req.Timestamp == nil {
		now := time.Now().Unix()
		req.Timestamp = &now
	}
	result, err := a.client.config.HttpClient.NewRequest().
		SetBody(req).
		Post("/sdk/publish_control_event")
	if err != nil {
		return response, err
	}

	return response, json.Unmarshal(result.Bytes(), &response)
}
