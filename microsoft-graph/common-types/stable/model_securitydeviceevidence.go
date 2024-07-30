package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityDeviceEvidence struct {
	AzureAdDeviceId          *string                                  `json:"azureAdDeviceId,omitempty"`
	CreatedDateTime          *string                                  `json:"createdDateTime,omitempty"`
	DefenderAvStatus         *SecurityDeviceEvidenceDefenderAvStatus  `json:"defenderAvStatus,omitempty"`
	DetailedRoles            *[]string                                `json:"detailedRoles,omitempty"`
	DeviceDnsName            *string                                  `json:"deviceDnsName,omitempty"`
	FirstSeenDateTime        *string                                  `json:"firstSeenDateTime,omitempty"`
	HealthStatus             *SecurityDeviceEvidenceHealthStatus      `json:"healthStatus,omitempty"`
	IpInterfaces             *[]string                                `json:"ipInterfaces,omitempty"`
	LoggedOnUsers            *[]SecurityLoggedOnUser                  `json:"loggedOnUsers,omitempty"`
	MdeDeviceId              *string                                  `json:"mdeDeviceId,omitempty"`
	ODataType                *string                                  `json:"@odata.type,omitempty"`
	OnboardingStatus         *SecurityDeviceEvidenceOnboardingStatus  `json:"onboardingStatus,omitempty"`
	OsBuild                  *int64                                   `json:"osBuild,omitempty"`
	OsPlatform               *string                                  `json:"osPlatform,omitempty"`
	RbacGroupId              *int64                                   `json:"rbacGroupId,omitempty"`
	RbacGroupName            *string                                  `json:"rbacGroupName,omitempty"`
	RemediationStatus        *SecurityDeviceEvidenceRemediationStatus `json:"remediationStatus,omitempty"`
	RemediationStatusDetails *string                                  `json:"remediationStatusDetails,omitempty"`
	RiskScore                *SecurityDeviceEvidenceRiskScore         `json:"riskScore,omitempty"`
	Roles                    *SecurityDeviceEvidenceRoles             `json:"roles,omitempty"`
	Tags                     *[]string                                `json:"tags,omitempty"`
	Verdict                  *SecurityDeviceEvidenceVerdict           `json:"verdict,omitempty"`
	Version                  *string                                  `json:"version,omitempty"`
	VmMetadata               *SecurityVmMetadata                      `json:"vmMetadata,omitempty"`
}
