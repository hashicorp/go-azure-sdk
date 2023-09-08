package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsConditionalAccessPolicyCoverage struct {
	ConditionalAccessPolicyState *string `json:"conditionalAccessPolicyState,omitempty"`
	Id                           *string `json:"id,omitempty"`
	LatestPolicyModifiedDateTime *string `json:"latestPolicyModifiedDateTime,omitempty"`
	ODataType                    *string `json:"@odata.type,omitempty"`
	RequiresDeviceCompliance     *bool   `json:"requiresDeviceCompliance,omitempty"`
	TenantDisplayName            *string `json:"tenantDisplayName,omitempty"`
}
