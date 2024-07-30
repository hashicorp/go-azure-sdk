package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecurityProfile struct {
	AzureSubscriptionId    *string                                     `json:"azureSubscriptionId,omitempty"`
	AzureTenantId          *string                                     `json:"azureTenantId,omitempty"`
	CreatedDateTime        *string                                     `json:"createdDateTime,omitempty"`
	DeploymentPackageUrl   *string                                     `json:"deploymentPackageUrl,omitempty"`
	DestinationServiceName *string                                     `json:"destinationServiceName,omitempty"`
	Id                     *string                                     `json:"id,omitempty"`
	IsSigned               *bool                                       `json:"isSigned,omitempty"`
	LastModifiedDateTime   *string                                     `json:"lastModifiedDateTime,omitempty"`
	Manifest               *string                                     `json:"manifest,omitempty"`
	Name                   *string                                     `json:"name,omitempty"`
	ODataType              *string                                     `json:"@odata.type,omitempty"`
	PermissionsRequired    *CloudAppSecurityProfilePermissionsRequired `json:"permissionsRequired,omitempty"`
	Platform               *string                                     `json:"platform,omitempty"`
	PolicyName             *string                                     `json:"policyName,omitempty"`
	Publisher              *string                                     `json:"publisher,omitempty"`
	RiskScore              *string                                     `json:"riskScore,omitempty"`
	Tags                   *[]string                                   `json:"tags,omitempty"`
	Type                   *string                                     `json:"type,omitempty"`
	VendorInformation      *SecurityVendorInformation                  `json:"vendorInformation,omitempty"`
}
