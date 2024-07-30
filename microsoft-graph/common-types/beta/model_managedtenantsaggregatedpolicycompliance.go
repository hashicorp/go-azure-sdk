package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsAggregatedPolicyCompliance struct {
	CompliancePolicyId          *string `json:"compliancePolicyId,omitempty"`
	CompliancePolicyName        *string `json:"compliancePolicyName,omitempty"`
	CompliancePolicyPlatform    *string `json:"compliancePolicyPlatform,omitempty"`
	CompliancePolicyType        *string `json:"compliancePolicyType,omitempty"`
	Id                          *string `json:"id,omitempty"`
	LastRefreshedDateTime       *string `json:"lastRefreshedDateTime,omitempty"`
	NumberOfCompliantDevices    *int64  `json:"numberOfCompliantDevices,omitempty"`
	NumberOfErrorDevices        *int64  `json:"numberOfErrorDevices,omitempty"`
	NumberOfNonCompliantDevices *int64  `json:"numberOfNonCompliantDevices,omitempty"`
	ODataType                   *string `json:"@odata.type,omitempty"`
	PolicyModifiedDateTime      *string `json:"policyModifiedDateTime,omitempty"`
	TenantDisplayName           *string `json:"tenantDisplayName,omitempty"`
	TenantId                    *string `json:"tenantId,omitempty"`
}
