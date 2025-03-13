package dataconnectorsconnect

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataConnectorConnectBody struct {
	ApiKey                        *string          `json:"apiKey,omitempty"`
	AuthorizationCode             *string          `json:"authorizationCode,omitempty"`
	ClientId                      *string          `json:"clientId,omitempty"`
	ClientSecret                  *string          `json:"clientSecret,omitempty"`
	DataCollectionEndpoint        *string          `json:"dataCollectionEndpoint,omitempty"`
	DataCollectionRuleImmutableId *string          `json:"dataCollectionRuleImmutableId,omitempty"`
	Kind                          *ConnectAuthKind `json:"kind,omitempty"`
	OutputStream                  *string          `json:"outputStream,omitempty"`
	Password                      *string          `json:"password,omitempty"`
	RequestConfigUserInputValues  *[]interface{}   `json:"requestConfigUserInputValues,omitempty"`
	UserName                      *string          `json:"userName,omitempty"`
}
