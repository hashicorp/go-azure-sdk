package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationManagerClientEnabledFeatures struct {
	CompliancePolicy         *bool   `json:"compliancePolicy,omitempty"`
	DeviceConfiguration      *bool   `json:"deviceConfiguration,omitempty"`
	Inventory                *bool   `json:"inventory,omitempty"`
	ModernApps               *bool   `json:"modernApps,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	ResourceAccess           *bool   `json:"resourceAccess,omitempty"`
	WindowsUpdateForBusiness *bool   `json:"windowsUpdateForBusiness,omitempty"`
}
