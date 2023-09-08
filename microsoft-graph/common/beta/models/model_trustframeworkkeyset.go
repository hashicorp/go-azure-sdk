package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustFrameworkKeySet struct {
	Id        *string              `json:"id,omitempty"`
	Keys      *[]TrustFrameworkKey `json:"keys,omitempty"`
	ODataType *string              `json:"@odata.type,omitempty"`
}
