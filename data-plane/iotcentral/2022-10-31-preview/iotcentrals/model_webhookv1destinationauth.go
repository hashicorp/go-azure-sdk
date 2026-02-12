package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebhookV1DestinationAuth interface {
	WebhookV1DestinationAuth() BaseWebhookV1DestinationAuthImpl
}

var _ WebhookV1DestinationAuth = BaseWebhookV1DestinationAuthImpl{}

type BaseWebhookV1DestinationAuthImpl struct {
	Type string `json:"type"`
}

func (s BaseWebhookV1DestinationAuthImpl) WebhookV1DestinationAuth() BaseWebhookV1DestinationAuthImpl {
	return s
}

var _ WebhookV1DestinationAuth = RawWebhookV1DestinationAuthImpl{}

// RawWebhookV1DestinationAuthImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWebhookV1DestinationAuthImpl struct {
	webhookV1DestinationAuth BaseWebhookV1DestinationAuthImpl
	Type                     string
	Values                   map[string]interface{}
}

func (s RawWebhookV1DestinationAuthImpl) WebhookV1DestinationAuth() BaseWebhookV1DestinationAuthImpl {
	return s.webhookV1DestinationAuth
}

func UnmarshalWebhookV1DestinationAuthImplementation(input []byte) (WebhookV1DestinationAuth, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WebhookV1DestinationAuth into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "header") {
		var out WebhookV1DestinationHeaderAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebhookV1DestinationHeaderAuth: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "oauth") {
		var out WebhookV1DestinationOAuthAuth
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WebhookV1DestinationOAuthAuth: %+v", err)
		}
		return out, nil
	}

	var parent BaseWebhookV1DestinationAuthImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWebhookV1DestinationAuthImpl: %+v", err)
	}

	return RawWebhookV1DestinationAuthImpl{
		webhookV1DestinationAuth: parent,
		Type:                     value,
		Values:                   temp,
	}, nil

}
