package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetPolicyEndpoints struct {
	ODataType     *string   `json:"@odata.type,omitempty"`
	PlatformTypes *[]string `json:"platformTypes,omitempty"`
}
