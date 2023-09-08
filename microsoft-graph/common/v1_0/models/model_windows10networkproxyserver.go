package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10NetworkProxyServer struct {
	Address              *string   `json:"address,omitempty"`
	Exceptions           *[]string `json:"exceptions,omitempty"`
	ODataType            *string   `json:"@odata.type,omitempty"`
	UseForLocalAddresses *bool     `json:"useForLocalAddresses,omitempty"`
}
