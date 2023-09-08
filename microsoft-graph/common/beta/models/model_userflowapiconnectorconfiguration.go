package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserFlowApiConnectorConfiguration struct {
	ODataType               *string               `json:"@odata.type,omitempty"`
	PostAttributeCollection *IdentityApiConnector `json:"postAttributeCollection,omitempty"`
	PostFederationSignup    *IdentityApiConnector `json:"postFederationSignup,omitempty"`
	PreTokenIssuance        *IdentityApiConnector `json:"preTokenIssuance,omitempty"`
}
