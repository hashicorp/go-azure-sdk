package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessHeaders struct {
	ODataType     *string `json:"@odata.type,omitempty"`
	Origin        *string `json:"origin,omitempty"`
	Referrer      *string `json:"referrer,omitempty"`
	XForwardedFor *string `json:"xForwardedFor,omitempty"`
}
