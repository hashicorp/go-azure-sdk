package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformation struct {
	DelegatedPrivilegeStatus              *ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus          `json:"delegatedPrivilegeStatus,omitempty"`
	LastDelegatedPrivilegeRefreshDateTime *string                                                                 `json:"lastDelegatedPrivilegeRefreshDateTime,omitempty"`
	ODataType                             *string                                                                 `json:"@odata.type,omitempty"`
	OffboardedByUserId                    *string                                                                 `json:"offboardedByUserId,omitempty"`
	OffboardedDateTime                    *string                                                                 `json:"offboardedDateTime,omitempty"`
	OnboardedByUserId                     *string                                                                 `json:"onboardedByUserId,omitempty"`
	OnboardedDateTime                     *string                                                                 `json:"onboardedDateTime,omitempty"`
	OnboardingStatus                      *ManagedTenantsTenantStatusInformationOnboardingStatus                  `json:"onboardingStatus,omitempty"`
	TenantOnboardingEligibilityReason     *ManagedTenantsTenantStatusInformationTenantOnboardingEligibilityReason `json:"tenantOnboardingEligibilityReason,omitempty"`
	WorkloadStatuses                      *[]ManagedTenantsWorkloadStatus                                         `json:"workloadStatuses,omitempty"`
}
