package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WebhookV1DestinationAuth = WebhookV1DestinationHeaderAuth{}

type WebhookV1DestinationHeaderAuth struct {
	Value string `json:"value"`

	// Fields inherited from WebhookV1DestinationAuth

	Type string `json:"type"`
}

func (s WebhookV1DestinationHeaderAuth) WebhookV1DestinationAuth() BaseWebhookV1DestinationAuthImpl {
	return BaseWebhookV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = WebhookV1DestinationHeaderAuth{}

func (s WebhookV1DestinationHeaderAuth) MarshalJSON() ([]byte, error) {
	type wrapper WebhookV1DestinationHeaderAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebhookV1DestinationHeaderAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebhookV1DestinationHeaderAuth: %+v", err)
	}

	decoded["type"] = "header"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebhookV1DestinationHeaderAuth: %+v", err)
	}

	return encoded, nil
}
