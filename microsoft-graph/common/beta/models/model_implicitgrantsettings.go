package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImplicitGrantSettings struct {
	EnableAccessTokenIssuance *bool   `json:"enableAccessTokenIssuance,omitempty"`
	EnableIdTokenIssuance     *bool   `json:"enableIdTokenIssuance,omitempty"`
	ODataType                 *string `json:"@odata.type,omitempty"`
}
