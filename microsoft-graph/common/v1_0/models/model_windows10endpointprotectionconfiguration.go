package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfiguration struct {
	AppLockerApplicationControl                          *Windows10EndpointProtectionConfigurationAppLockerApplicationControl                  `json:"appLockerApplicationControl,omitempty"`
	ApplicationGuardAllowPersistence                     *bool                                                                                 `json:"applicationGuardAllowPersistence,omitempty"`
	ApplicationGuardAllowPrintToLocalPrinters            *bool                                                                                 `json:"applicationGuardAllowPrintToLocalPrinters,omitempty"`
	ApplicationGuardAllowPrintToNetworkPrinters          *bool                                                                                 `json:"applicationGuardAllowPrintToNetworkPrinters,omitempty"`
	ApplicationGuardAllowPrintToPDF                      *bool                                                                                 `json:"applicationGuardAllowPrintToPDF,omitempty"`
	ApplicationGuardAllowPrintToXPS                      *bool                                                                                 `json:"applicationGuardAllowPrintToXPS,omitempty"`
	ApplicationGuardBlockClipboardSharing                *Windows10EndpointProtectionConfigurationApplicationGuardBlockClipboardSharing        `json:"applicationGuardBlockClipboardSharing,omitempty"`
	ApplicationGuardBlockFileTransfer                    *Windows10EndpointProtectionConfigurationApplicationGuardBlockFileTransfer            `json:"applicationGuardBlockFileTransfer,omitempty"`
	ApplicationGuardBlockNonEnterpriseContent            *bool                                                                                 `json:"applicationGuardBlockNonEnterpriseContent,omitempty"`
	ApplicationGuardEnabled                              *bool                                                                                 `json:"applicationGuardEnabled,omitempty"`
	ApplicationGuardForceAuditing                        *bool                                                                                 `json:"applicationGuardForceAuditing,omitempty"`
	Assignments                                          *[]DeviceConfigurationAssignment                                                      `json:"assignments,omitempty"`
	BitLockerDisableWarningForOtherDiskEncryption        *bool                                                                                 `json:"bitLockerDisableWarningForOtherDiskEncryption,omitempty"`
	BitLockerEnableStorageCardEncryptionOnMobile         *bool                                                                                 `json:"bitLockerEnableStorageCardEncryptionOnMobile,omitempty"`
	BitLockerEncryptDevice                               *bool                                                                                 `json:"bitLockerEncryptDevice,omitempty"`
	BitLockerRemovableDrivePolicy                        *BitLockerRemovableDrivePolicy                                                        `json:"bitLockerRemovableDrivePolicy,omitempty"`
	CreatedDateTime                                      *string                                                                               `json:"createdDateTime,omitempty"`
	DefenderAdditionalGuardedFolders                     *[]string                                                                             `json:"defenderAdditionalGuardedFolders,omitempty"`
	DefenderAttackSurfaceReductionExcludedPaths          *[]string                                                                             `json:"defenderAttackSurfaceReductionExcludedPaths,omitempty"`
	DefenderExploitProtectionXml                         *string                                                                               `json:"defenderExploitProtectionXml,omitempty"`
	DefenderExploitProtectionXmlFileName                 *string                                                                               `json:"defenderExploitProtectionXmlFileName,omitempty"`
	DefenderGuardedFoldersAllowedAppPaths                *[]string                                                                             `json:"defenderGuardedFoldersAllowedAppPaths,omitempty"`
	DefenderSecurityCenterBlockExploitProtectionOverride *bool                                                                                 `json:"defenderSecurityCenterBlockExploitProtectionOverride,omitempty"`
	Description                                          *string                                                                               `json:"description,omitempty"`
	DeviceSettingStateSummaries                          *[]SettingStateDeviceSummary                                                          `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                                 *DeviceConfigurationDeviceOverview                                                    `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                       *[]DeviceConfigurationDeviceStatus                                                    `json:"deviceStatuses,omitempty"`
	DisplayName                                          *string                                                                               `json:"displayName,omitempty"`
	FirewallBlockStatefulFTP                             *bool                                                                                 `json:"firewallBlockStatefulFTP,omitempty"`
	FirewallCertificateRevocationListCheckMethod         *Windows10EndpointProtectionConfigurationFirewallCertificateRevocationListCheckMethod `json:"firewallCertificateRevocationListCheckMethod,omitempty"`
	FirewallIPSecExemptionsAllowDHCP                     *bool                                                                                 `json:"firewallIPSecExemptionsAllowDHCP,omitempty"`
	FirewallIPSecExemptionsAllowICMP                     *bool                                                                                 `json:"firewallIPSecExemptionsAllowICMP,omitempty"`
	FirewallIPSecExemptionsAllowNeighborDiscovery        *bool                                                                                 `json:"firewallIPSecExemptionsAllowNeighborDiscovery,omitempty"`
	FirewallIPSecExemptionsAllowRouterDiscovery          *bool                                                                                 `json:"firewallIPSecExemptionsAllowRouterDiscovery,omitempty"`
	FirewallIdleTimeoutForSecurityAssociationInSeconds   *int64                                                                                `json:"firewallIdleTimeoutForSecurityAssociationInSeconds,omitempty"`
	FirewallMergeKeyingModuleSettings                    *bool                                                                                 `json:"firewallMergeKeyingModuleSettings,omitempty"`
	FirewallPacketQueueingMethod                         *Windows10EndpointProtectionConfigurationFirewallPacketQueueingMethod                 `json:"firewallPacketQueueingMethod,omitempty"`
	FirewallPreSharedKeyEncodingMethod                   *Windows10EndpointProtectionConfigurationFirewallPreSharedKeyEncodingMethod           `json:"firewallPreSharedKeyEncodingMethod,omitempty"`
	FirewallProfileDomain                                *WindowsFirewallNetworkProfile                                                        `json:"firewallProfileDomain,omitempty"`
	FirewallProfilePrivate                               *WindowsFirewallNetworkProfile                                                        `json:"firewallProfilePrivate,omitempty"`
	FirewallProfilePublic                                *WindowsFirewallNetworkProfile                                                        `json:"firewallProfilePublic,omitempty"`
	Id                                                   *string                                                                               `json:"id,omitempty"`
	LastModifiedDateTime                                 *string                                                                               `json:"lastModifiedDateTime,omitempty"`
	ODataType                                            *string                                                                               `json:"@odata.type,omitempty"`
	SmartScreenBlockOverrideForFiles                     *bool                                                                                 `json:"smartScreenBlockOverrideForFiles,omitempty"`
	SmartScreenEnableInShell                             *bool                                                                                 `json:"smartScreenEnableInShell,omitempty"`
	UserStatusOverview                                   *DeviceConfigurationUserOverview                                                      `json:"userStatusOverview,omitempty"`
	UserStatuses                                         *[]DeviceConfigurationUserStatus                                                      `json:"userStatuses,omitempty"`
	Version                                              *int64                                                                                `json:"version,omitempty"`
}
