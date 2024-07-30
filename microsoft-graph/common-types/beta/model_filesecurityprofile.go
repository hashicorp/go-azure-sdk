package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileSecurityProfile struct {
	ActivityGroupNames    *[]string                  `json:"activityGroupNames,omitempty"`
	AzureSubscriptionId   *string                    `json:"azureSubscriptionId,omitempty"`
	AzureTenantId         *string                    `json:"azureTenantId,omitempty"`
	CertificateThumbprint *string                    `json:"certificateThumbprint,omitempty"`
	Extensions            *[]string                  `json:"extensions,omitempty"`
	FileType              *string                    `json:"fileType,omitempty"`
	FirstSeenDateTime     *string                    `json:"firstSeenDateTime,omitempty"`
	Hashes                *[]FileHash                `json:"hashes,omitempty"`
	Id                    *string                    `json:"id,omitempty"`
	LastSeenDateTime      *string                    `json:"lastSeenDateTime,omitempty"`
	MalwareStates         *[]MalwareState            `json:"malwareStates,omitempty"`
	Names                 *[]string                  `json:"names,omitempty"`
	ODataType             *string                    `json:"@odata.type,omitempty"`
	RiskScore             *string                    `json:"riskScore,omitempty"`
	Size                  *int64                     `json:"size,omitempty"`
	Tags                  *[]string                  `json:"tags,omitempty"`
	VendorInformation     *SecurityVendorInformation `json:"vendorInformation,omitempty"`
	VulnerabilityStates   *[]VulnerabilityState      `json:"vulnerabilityStates,omitempty"`
}
