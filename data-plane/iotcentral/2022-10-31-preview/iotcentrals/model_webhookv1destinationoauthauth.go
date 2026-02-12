package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ WebhookV1DestinationAuth = WebhookV1DestinationOAuthAuth{}

type WebhookV1DestinationOAuthAuth struct {
	Audience     *string           `json:"audience,omitempty"`
	ClientId     string            `json:"clientId"`
	ClientSecret string            `json:"clientSecret"`
	RequestType  *OAuthRequestType `json:"requestType,omitempty"`
	Scope        *string           `json:"scope,omitempty"`
	TokenURL     string            `json:"tokenUrl"`

	// Fields inherited from WebhookV1DestinationAuth

	Type string `json:"type"`
}

func (s WebhookV1DestinationOAuthAuth) WebhookV1DestinationAuth() BaseWebhookV1DestinationAuthImpl {
	return BaseWebhookV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = WebhookV1DestinationOAuthAuth{}

func (s WebhookV1DestinationOAuthAuth) MarshalJSON() ([]byte, error) {
	type wrapper WebhookV1DestinationOAuthAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling WebhookV1DestinationOAuthAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling WebhookV1DestinationOAuthAuth: %+v", err)
	}

	decoded["type"] = "oauth"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling WebhookV1DestinationOAuthAuth: %+v", err)
	}

	return encoded, nil
}
