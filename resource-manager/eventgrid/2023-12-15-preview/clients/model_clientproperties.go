package clients

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClientProperties struct {
	Attributes                      *map[string]interface{}          `json:"attributes,omitempty"`
	AuthenticationName              *string                          `json:"authenticationName,omitempty"`
	ClientCertificateAuthentication *ClientCertificateAuthentication `json:"clientCertificateAuthentication,omitempty"`
	Description                     *string                          `json:"description,omitempty"`
	ProvisioningState               *ClientProvisioningState         `json:"provisioningState,omitempty"`
	State                           *ClientState                     `json:"state,omitempty"`
}
