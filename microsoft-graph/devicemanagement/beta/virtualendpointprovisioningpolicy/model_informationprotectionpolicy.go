package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionPolicy struct {
	Id        *string                       `json:"id,omitempty"`
	Labels    *[]InformationProtectionLabel `json:"labels,omitempty"`
	ODataType *string                       `json:"@odata.type,omitempty"`
}
