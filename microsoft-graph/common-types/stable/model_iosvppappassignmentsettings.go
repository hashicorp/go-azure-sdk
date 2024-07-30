package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVppAppAssignmentSettings struct {
	ODataType          *string `json:"@odata.type,omitempty"`
	UseDeviceLicensing *bool   `json:"useDeviceLicensing,omitempty"`
	VpnConfigurationId *string `json:"vpnConfigurationId,omitempty"`
}
