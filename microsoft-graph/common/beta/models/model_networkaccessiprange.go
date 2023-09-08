package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessIpRange struct {
	BeginAddress *string `json:"beginAddress,omitempty"`
	EndAddress   *string `json:"endAddress,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
}
