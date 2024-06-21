package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProperties struct {
	AadApplicationObjectId             *string                             `json:"aadApplicationObjectId,omitempty"`
	AadClientId                        *string                             `json:"aadClientId,omitempty"`
	AadServicePrincipalObjectId        *string                             `json:"aadServicePrincipalObjectId,omitempty"`
	AadTenantId                        *string                             `json:"aadTenantId,omitempty"`
	BillingModel                       *string                             `json:"billingModel,omitempty"`
	CloudId                            *string                             `json:"cloudId,omitempty"`
	CloudManagementEndpoint            *string                             `json:"cloudManagementEndpoint,omitempty"`
	ConnectivityStatus                 *ConnectivityStatus                 `json:"connectivityStatus,omitempty"`
	DesiredProperties                  *ClusterDesiredProperties           `json:"desiredProperties,omitempty"`
	IsolatedVMAttestationConfiguration *IsolatedVMAttestationConfiguration `json:"isolatedVmAttestationConfiguration,omitempty"`
	LastBillingTimestamp               *string                             `json:"lastBillingTimestamp,omitempty"`
	LastSyncTimestamp                  *string                             `json:"lastSyncTimestamp,omitempty"`
	ProvisioningState                  *ProvisioningState                  `json:"provisioningState,omitempty"`
	RegistrationTimestamp              *string                             `json:"registrationTimestamp,omitempty"`
	ReportedProperties                 *ClusterReportedProperties          `json:"reportedProperties,omitempty"`
	ResourceProviderObjectId           *string                             `json:"resourceProviderObjectId,omitempty"`
	ServiceEndpoint                    *string                             `json:"serviceEndpoint,omitempty"`
	SoftwareAssuranceProperties        *SoftwareAssuranceProperties        `json:"softwareAssuranceProperties,omitempty"`
	Status                             *Status                             `json:"status,omitempty"`
	TrialDaysRemaining                 *float64                            `json:"trialDaysRemaining,omitempty"`
}
