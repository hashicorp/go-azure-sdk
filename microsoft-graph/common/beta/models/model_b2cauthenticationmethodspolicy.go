package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2cAuthenticationMethodsPolicy struct {
	Id                                          *string `json:"id,omitempty"`
	IsEmailPasswordAuthenticationEnabled        *bool   `json:"isEmailPasswordAuthenticationEnabled,omitempty"`
	IsPhoneOneTimePasswordAuthenticationEnabled *bool   `json:"isPhoneOneTimePasswordAuthenticationEnabled,omitempty"`
	IsUserNameAuthenticationEnabled             *bool   `json:"isUserNameAuthenticationEnabled,omitempty"`
	ODataType                                   *string `json:"@odata.type,omitempty"`
}
