package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = OAuthModel{}

type OAuthModel struct {
	AccessTokenPrepend                   *string            `json:"accessTokenPrepend,omitempty"`
	AuthorizationCode                    *string            `json:"authorizationCode,omitempty"`
	AuthorizationEndpoint                *string            `json:"authorizationEndpoint,omitempty"`
	AuthorizationEndpointHeaders         *map[string]string `json:"authorizationEndpointHeaders,omitempty"`
	AuthorizationEndpointQueryParameters *map[string]string `json:"authorizationEndpointQueryParameters,omitempty"`
	ClientId                             string             `json:"clientId"`
	ClientSecret                         string             `json:"clientSecret"`
	GrantType                            string             `json:"grantType"`
	IsCredentialsInHeaders               *bool              `json:"isCredentialsInHeaders,omitempty"`
	IsJwtBearerFlow                      *bool              `json:"isJwtBearerFlow,omitempty"`
	RedirectUri                          *string            `json:"redirectUri,omitempty"`
	Scope                                *string            `json:"scope,omitempty"`
	TokenEndpoint                        string             `json:"tokenEndpoint"`
	TokenEndpointHeaders                 *map[string]string `json:"tokenEndpointHeaders,omitempty"`
	TokenEndpointQueryParameters         *map[string]string `json:"tokenEndpointQueryParameters,omitempty"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s OAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = OAuthModel{}

func (s OAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper OAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OAuthModel: %+v", err)
	}

	decoded["type"] = "OAuth2"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OAuthModel: %+v", err)
	}

	return encoded, nil
}
