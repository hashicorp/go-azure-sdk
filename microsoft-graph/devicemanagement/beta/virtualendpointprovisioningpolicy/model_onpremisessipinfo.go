package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesSipInfo struct {
	IsSipEnabled          *bool   `json:"isSipEnabled,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
	SipDeploymentLocation *string `json:"sipDeploymentLocation,omitempty"`
	SipPrimaryAddress     *string `json:"sipPrimaryAddress,omitempty"`
}
