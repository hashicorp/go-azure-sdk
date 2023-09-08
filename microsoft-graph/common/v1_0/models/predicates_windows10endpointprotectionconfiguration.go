package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationOperationPredicate struct {
	ApplicationGuardAllowPersistence                     *bool
	ApplicationGuardAllowPrintToLocalPrinters            *bool
	ApplicationGuardAllowPrintToNetworkPrinters          *bool
	ApplicationGuardAllowPrintToPDF                      *bool
	ApplicationGuardAllowPrintToXPS                      *bool
	ApplicationGuardBlockNonEnterpriseContent            *bool
	ApplicationGuardEnabled                              *bool
	ApplicationGuardForceAuditing                        *bool
	BitLockerDisableWarningForOtherDiskEncryption        *bool
	BitLockerEnableStorageCardEncryptionOnMobile         *bool
	BitLockerEncryptDevice                               *bool
	CreatedDateTime                                      *string
	DefenderExploitProtectionXml                         *string
	DefenderExploitProtectionXmlFileName                 *string
	DefenderSecurityCenterBlockExploitProtectionOverride *bool
	Description                                          *string
	DisplayName                                          *string
	FirewallBlockStatefulFTP                             *bool
	FirewallIPSecExemptionsAllowDHCP                     *bool
	FirewallIPSecExemptionsAllowICMP                     *bool
	FirewallIPSecExemptionsAllowNeighborDiscovery        *bool
	FirewallIPSecExemptionsAllowRouterDiscovery          *bool
	FirewallIdleTimeoutForSecurityAssociationInSeconds   *int64
	FirewallMergeKeyingModuleSettings                    *bool
	Id                                                   *string
	LastModifiedDateTime                                 *string
	ODataType                                            *string
	SmartScreenBlockOverrideForFiles                     *bool
	SmartScreenEnableInShell                             *bool
	Version                                              *int64
}

func (p Windows10EndpointProtectionConfigurationOperationPredicate) Matches(input Windows10EndpointProtectionConfiguration) bool {

	if p.ApplicationGuardAllowPersistence != nil && (input.ApplicationGuardAllowPersistence == nil || *p.ApplicationGuardAllowPersistence != *input.ApplicationGuardAllowPersistence) {
		return false
	}

	if p.ApplicationGuardAllowPrintToLocalPrinters != nil && (input.ApplicationGuardAllowPrintToLocalPrinters == nil || *p.ApplicationGuardAllowPrintToLocalPrinters != *input.ApplicationGuardAllowPrintToLocalPrinters) {
		return false
	}

	if p.ApplicationGuardAllowPrintToNetworkPrinters != nil && (input.ApplicationGuardAllowPrintToNetworkPrinters == nil || *p.ApplicationGuardAllowPrintToNetworkPrinters != *input.ApplicationGuardAllowPrintToNetworkPrinters) {
		return false
	}

	if p.ApplicationGuardAllowPrintToPDF != nil && (input.ApplicationGuardAllowPrintToPDF == nil || *p.ApplicationGuardAllowPrintToPDF != *input.ApplicationGuardAllowPrintToPDF) {
		return false
	}

	if p.ApplicationGuardAllowPrintToXPS != nil && (input.ApplicationGuardAllowPrintToXPS == nil || *p.ApplicationGuardAllowPrintToXPS != *input.ApplicationGuardAllowPrintToXPS) {
		return false
	}

	if p.ApplicationGuardBlockNonEnterpriseContent != nil && (input.ApplicationGuardBlockNonEnterpriseContent == nil || *p.ApplicationGuardBlockNonEnterpriseContent != *input.ApplicationGuardBlockNonEnterpriseContent) {
		return false
	}

	if p.ApplicationGuardEnabled != nil && (input.ApplicationGuardEnabled == nil || *p.ApplicationGuardEnabled != *input.ApplicationGuardEnabled) {
		return false
	}

	if p.ApplicationGuardForceAuditing != nil && (input.ApplicationGuardForceAuditing == nil || *p.ApplicationGuardForceAuditing != *input.ApplicationGuardForceAuditing) {
		return false
	}

	if p.BitLockerDisableWarningForOtherDiskEncryption != nil && (input.BitLockerDisableWarningForOtherDiskEncryption == nil || *p.BitLockerDisableWarningForOtherDiskEncryption != *input.BitLockerDisableWarningForOtherDiskEncryption) {
		return false
	}

	if p.BitLockerEnableStorageCardEncryptionOnMobile != nil && (input.BitLockerEnableStorageCardEncryptionOnMobile == nil || *p.BitLockerEnableStorageCardEncryptionOnMobile != *input.BitLockerEnableStorageCardEncryptionOnMobile) {
		return false
	}

	if p.BitLockerEncryptDevice != nil && (input.BitLockerEncryptDevice == nil || *p.BitLockerEncryptDevice != *input.BitLockerEncryptDevice) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DefenderExploitProtectionXml != nil && (input.DefenderExploitProtectionXml == nil || *p.DefenderExploitProtectionXml != *input.DefenderExploitProtectionXml) {
		return false
	}

	if p.DefenderExploitProtectionXmlFileName != nil && (input.DefenderExploitProtectionXmlFileName == nil || *p.DefenderExploitProtectionXmlFileName != *input.DefenderExploitProtectionXmlFileName) {
		return false
	}

	if p.DefenderSecurityCenterBlockExploitProtectionOverride != nil && (input.DefenderSecurityCenterBlockExploitProtectionOverride == nil || *p.DefenderSecurityCenterBlockExploitProtectionOverride != *input.DefenderSecurityCenterBlockExploitProtectionOverride) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.FirewallBlockStatefulFTP != nil && (input.FirewallBlockStatefulFTP == nil || *p.FirewallBlockStatefulFTP != *input.FirewallBlockStatefulFTP) {
		return false
	}

	if p.FirewallIPSecExemptionsAllowDHCP != nil && (input.FirewallIPSecExemptionsAllowDHCP == nil || *p.FirewallIPSecExemptionsAllowDHCP != *input.FirewallIPSecExemptionsAllowDHCP) {
		return false
	}

	if p.FirewallIPSecExemptionsAllowICMP != nil && (input.FirewallIPSecExemptionsAllowICMP == nil || *p.FirewallIPSecExemptionsAllowICMP != *input.FirewallIPSecExemptionsAllowICMP) {
		return false
	}

	if p.FirewallIPSecExemptionsAllowNeighborDiscovery != nil && (input.FirewallIPSecExemptionsAllowNeighborDiscovery == nil || *p.FirewallIPSecExemptionsAllowNeighborDiscovery != *input.FirewallIPSecExemptionsAllowNeighborDiscovery) {
		return false
	}

	if p.FirewallIPSecExemptionsAllowRouterDiscovery != nil && (input.FirewallIPSecExemptionsAllowRouterDiscovery == nil || *p.FirewallIPSecExemptionsAllowRouterDiscovery != *input.FirewallIPSecExemptionsAllowRouterDiscovery) {
		return false
	}

	if p.FirewallIdleTimeoutForSecurityAssociationInSeconds != nil && (input.FirewallIdleTimeoutForSecurityAssociationInSeconds == nil || *p.FirewallIdleTimeoutForSecurityAssociationInSeconds != *input.FirewallIdleTimeoutForSecurityAssociationInSeconds) {
		return false
	}

	if p.FirewallMergeKeyingModuleSettings != nil && (input.FirewallMergeKeyingModuleSettings == nil || *p.FirewallMergeKeyingModuleSettings != *input.FirewallMergeKeyingModuleSettings) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SmartScreenBlockOverrideForFiles != nil && (input.SmartScreenBlockOverrideForFiles == nil || *p.SmartScreenBlockOverrideForFiles != *input.SmartScreenBlockOverrideForFiles) {
		return false
	}

	if p.SmartScreenEnableInShell != nil && (input.SmartScreenEnableInShell == nil || *p.SmartScreenEnableInShell != *input.SmartScreenEnableInShell) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
