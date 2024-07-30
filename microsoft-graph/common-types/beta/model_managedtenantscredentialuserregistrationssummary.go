package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsCredentialUserRegistrationsSummary struct {
	Id                              *string `json:"id,omitempty"`
	LastRefreshedDateTime           *string `json:"lastRefreshedDateTime,omitempty"`
	MfaAndSsprCapableUserCount      *int64  `json:"mfaAndSsprCapableUserCount,omitempty"`
	MfaConditionalAccessPolicyState *string `json:"mfaConditionalAccessPolicyState,omitempty"`
	MfaExcludedUserCount            *int64  `json:"mfaExcludedUserCount,omitempty"`
	MfaRegisteredUserCount          *int64  `json:"mfaRegisteredUserCount,omitempty"`
	ODataType                       *string `json:"@odata.type,omitempty"`
	SecurityDefaultsEnabled         *bool   `json:"securityDefaultsEnabled,omitempty"`
	SsprEnabledUserCount            *int64  `json:"ssprEnabledUserCount,omitempty"`
	SsprRegisteredUserCount         *int64  `json:"ssprRegisteredUserCount,omitempty"`
	TenantDisplayName               *string `json:"tenantDisplayName,omitempty"`
	TenantId                        *string `json:"tenantId,omitempty"`
	TenantLicenseType               *string `json:"tenantLicenseType,omitempty"`
	TotalUserCount                  *int64  `json:"totalUserCount,omitempty"`
}
