package sessionhostconfiguration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SessionHostConfigurationPatchProperties struct {
	AvailabilityZones            *[]int64                            `json:"availabilityZones,omitempty"`
	BootDiagnosticsInfo          *BootDiagnosticsInfoPatchProperties `json:"bootDiagnosticsInfo,omitempty"`
	CustomConfigurationScriptURL *string                             `json:"customConfigurationScriptUrl,omitempty"`
	DiskInfo                     *DiskInfoPatchProperties            `json:"diskInfo,omitempty"`
	DomainInfo                   *DomainInfoPatchProperties          `json:"domainInfo,omitempty"`
	FriendlyName                 *string                             `json:"friendlyName,omitempty"`
	ImageInfo                    *ImageInfoPatchProperties           `json:"imageInfo,omitempty"`
	NetworkInfo                  *NetworkInfoPatchProperties         `json:"networkInfo,omitempty"`
	SecurityInfo                 *SecurityInfoPatchProperties        `json:"securityInfo,omitempty"`
	VMAdminCredentials           *KeyVaultCredentialsPatchProperties `json:"vmAdminCredentials,omitempty"`
	VMLocation                   *string                             `json:"vmLocation,omitempty"`
	VMNamePrefix                 *string                             `json:"vmNamePrefix,omitempty"`
	VMResourceGroup              *string                             `json:"vmResourceGroup,omitempty"`
	VMSizeId                     *string                             `json:"vmSizeId,omitempty"`
	VMTags                       *map[string]string                  `json:"vmTags,omitempty"`
}
