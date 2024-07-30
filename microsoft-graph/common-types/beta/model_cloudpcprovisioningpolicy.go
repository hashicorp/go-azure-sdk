package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCProvisioningPolicy struct {
	AlternateResourceUrl     *string                                    `json:"alternateResourceUrl,omitempty"`
	Assignments              *[]CloudPCProvisioningPolicyAssignment     `json:"assignments,omitempty"`
	CloudPCGroupDisplayName  *string                                    `json:"cloudPcGroupDisplayName,omitempty"`
	CloudPCNamingTemplate    *string                                    `json:"cloudPcNamingTemplate,omitempty"`
	Description              *string                                    `json:"description,omitempty"`
	DisplayName              *string                                    `json:"displayName,omitempty"`
	DomainJoinConfiguration  *CloudPCDomainJoinConfiguration            `json:"domainJoinConfiguration,omitempty"`
	DomainJoinConfigurations *[]CloudPCDomainJoinConfiguration          `json:"domainJoinConfigurations,omitempty"`
	EnableSingleSignOn       *bool                                      `json:"enableSingleSignOn,omitempty"`
	GracePeriodInHours       *int64                                     `json:"gracePeriodInHours,omitempty"`
	Id                       *string                                    `json:"id,omitempty"`
	ImageDisplayName         *string                                    `json:"imageDisplayName,omitempty"`
	ImageId                  *string                                    `json:"imageId,omitempty"`
	ImageType                *CloudPCProvisioningPolicyImageType        `json:"imageType,omitempty"`
	LocalAdminEnabled        *bool                                      `json:"localAdminEnabled,omitempty"`
	ManagedBy                *CloudPCProvisioningPolicyManagedBy        `json:"managedBy,omitempty"`
	MicrosoftManagedDesktop  *MicrosoftManagedDesktop                   `json:"microsoftManagedDesktop,omitempty"`
	ODataType                *string                                    `json:"@odata.type,omitempty"`
	OnPremisesConnectionId   *string                                    `json:"onPremisesConnectionId,omitempty"`
	ProvisioningType         *CloudPCProvisioningPolicyProvisioningType `json:"provisioningType,omitempty"`
	ScopeIds                 *[]string                                  `json:"scopeIds,omitempty"`
	WindowsSettings          *CloudPCWindowsSettings                    `json:"windowsSettings,omitempty"`
}
