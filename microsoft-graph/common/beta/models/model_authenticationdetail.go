package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationDetail struct {
	AuthenticationMethod           *string `json:"authenticationMethod,omitempty"`
	AuthenticationMethodDetail     *string `json:"authenticationMethodDetail,omitempty"`
	AuthenticationStepDateTime     *string `json:"authenticationStepDateTime,omitempty"`
	AuthenticationStepRequirement  *string `json:"authenticationStepRequirement,omitempty"`
	AuthenticationStepResultDetail *string `json:"authenticationStepResultDetail,omitempty"`
	ODataType                      *string `json:"@odata.type,omitempty"`
	Succeeded                      *bool   `json:"succeeded,omitempty"`
}
