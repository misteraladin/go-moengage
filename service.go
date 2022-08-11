package moengage

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
)

type Gateway struct {
	Client Client
}

type MoEngageGateway interface {
	SendEvent(data interface{}) (Response, error)
}

func NewMoEngageGateway(req Client) MoEngageGateway {
	client := NewClient(req.BaseUrl, req.APIKey, req.APPID)

	return &Gateway{
		Client: client,
	}
}

func (gateway *Gateway) SendEvent(data interface{}) (Response, error) {
	path := "/v1/event/" + gateway.Client.APPID
	jsonReq, _ := json.Marshal(data)
	var response Response

	res, err := gateway.Client.Call("POST", gateway.Client.BaseUrl+path, bytes.NewBuffer(jsonReq))
	if err != nil {
		return response, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Print("Cannot read response body: ", err)
		return response, nil
	}

	err = json.Unmarshal(resBody, &response)
	if err != nil {
		return response, err
	}
	return response, nil
}
