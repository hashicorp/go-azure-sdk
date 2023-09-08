package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MdmWindowsInformationProtectionPolicy struct {
	Assignments                            *[]TargetedManagedAppPolicyAssignment                  `json:"assignments,omitempty"`
	AzureRightsManagementServicesAllowed   *bool                                                  `json:"azureRightsManagementServicesAllowed,omitempty"`
	CreatedDateTime                        *string                                                `json:"createdDateTime,omitempty"`
	DataRecoveryCertificate                *WindowsInformationProtectionDataRecoveryCertificate   `json:"dataRecoveryCertificate,omitempty"`
	Description                            *string                                                `json:"description,omitempty"`
	DisplayName                            *string                                                `json:"displayName,omitempty"`
	EnforcementLevel                       *MdmWindowsInformationProtectionPolicyEnforcementLevel `json:"enforcementLevel,omitempty"`
	EnterpriseDomain                       *string                                                `json:"enterpriseDomain,omitempty"`
	EnterpriseIPRanges                     *[]WindowsInformationProtectionIPRangeCollection       `json:"enterpriseIPRanges,omitempty"`
	EnterpriseIPRangesAreAuthoritative     *bool                                                  `json:"enterpriseIPRangesAreAuthoritative,omitempty"`
	EnterpriseInternalProxyServers         *[]WindowsInformationProtectionResourceCollection      `json:"enterpriseInternalProxyServers,omitempty"`
	EnterpriseNetworkDomainNames           *[]WindowsInformationProtectionResourceCollection      `json:"enterpriseNetworkDomainNames,omitempty"`
	EnterpriseProtectedDomainNames         *[]WindowsInformationProtectionResourceCollection      `json:"enterpriseProtectedDomainNames,omitempty"`
	EnterpriseProxiedDomains               *[]WindowsInformationProtectionProxiedDomainCollection `json:"enterpriseProxiedDomains,omitempty"`
	EnterpriseProxyServers                 *[]WindowsInformationProtectionResourceCollection      `json:"enterpriseProxyServers,omitempty"`
	EnterpriseProxyServersAreAuthoritative *bool                                                  `json:"enterpriseProxyServersAreAuthoritative,omitempty"`
	ExemptAppLockerFiles                   *[]WindowsInformationProtectionAppLockerFile           `json:"exemptAppLockerFiles,omitempty"`
	ExemptApps                             *[]WindowsInformationProtectionApp                     `json:"exemptApps,omitempty"`
	IconsVisible                           *bool                                                  `json:"iconsVisible,omitempty"`
	Id                                     *string                                                `json:"id,omitempty"`
	IndexingEncryptedStoresOrItemsBlocked  *bool                                                  `json:"indexingEncryptedStoresOrItemsBlocked,omitempty"`
	IsAssigned                             *bool                                                  `json:"isAssigned,omitempty"`
	LastModifiedDateTime                   *string                                                `json:"lastModifiedDateTime,omitempty"`
	NeutralDomainResources                 *[]WindowsInformationProtectionResourceCollection      `json:"neutralDomainResources,omitempty"`
	ODataType                              *string                                                `json:"@odata.type,omitempty"`
	ProtectedAppLockerFiles                *[]WindowsInformationProtectionAppLockerFile           `json:"protectedAppLockerFiles,omitempty"`
	ProtectedApps                          *[]WindowsInformationProtectionApp                     `json:"protectedApps,omitempty"`
	ProtectionUnderLockConfigRequired      *bool                                                  `json:"protectionUnderLockConfigRequired,omitempty"`
	RevokeOnUnenrollDisabled               *bool                                                  `json:"revokeOnUnenrollDisabled,omitempty"`
	RightsManagementServicesTemplateId     *string                                                `json:"rightsManagementServicesTemplateId,omitempty"`
	RoleScopeTagIds                        *[]string                                              `json:"roleScopeTagIds,omitempty"`
	SmbAutoEncryptedFileExtensions         *[]WindowsInformationProtectionResourceCollection      `json:"smbAutoEncryptedFileExtensions,omitempty"`
	Version                                *string                                                `json:"version,omitempty"`
}
