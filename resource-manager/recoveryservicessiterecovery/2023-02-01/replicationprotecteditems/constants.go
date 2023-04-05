package replicationprotecteditems

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AgentUpgradeBlockedReason string

const (
	AgentUpgradeBlockedReasonAgentNoHeartbeat              AgentUpgradeBlockedReason = "AgentNoHeartbeat"
	AgentUpgradeBlockedReasonAlreadyOnLatestVersion        AgentUpgradeBlockedReason = "AlreadyOnLatestVersion"
	AgentUpgradeBlockedReasonDistroIsNotReported           AgentUpgradeBlockedReason = "DistroIsNotReported"
	AgentUpgradeBlockedReasonDistroNotSupportedForUpgrade  AgentUpgradeBlockedReason = "DistroNotSupportedForUpgrade"
	AgentUpgradeBlockedReasonIncompatibleApplianceVersion  AgentUpgradeBlockedReason = "IncompatibleApplianceVersion"
	AgentUpgradeBlockedReasonInvalidAgentVersion           AgentUpgradeBlockedReason = "InvalidAgentVersion"
	AgentUpgradeBlockedReasonInvalidDriverVersion          AgentUpgradeBlockedReason = "InvalidDriverVersion"
	AgentUpgradeBlockedReasonMissingUpgradePath            AgentUpgradeBlockedReason = "MissingUpgradePath"
	AgentUpgradeBlockedReasonNotProtected                  AgentUpgradeBlockedReason = "NotProtected"
	AgentUpgradeBlockedReasonProcessServerNoHeartbeat      AgentUpgradeBlockedReason = "ProcessServerNoHeartbeat"
	AgentUpgradeBlockedReasonRcmProxyNoHeartbeat           AgentUpgradeBlockedReason = "RcmProxyNoHeartbeat"
	AgentUpgradeBlockedReasonRebootRequired                AgentUpgradeBlockedReason = "RebootRequired"
	AgentUpgradeBlockedReasonUnknown                       AgentUpgradeBlockedReason = "Unknown"
	AgentUpgradeBlockedReasonUnsupportedProtectionScenario AgentUpgradeBlockedReason = "UnsupportedProtectionScenario"
)

func PossibleValuesForAgentUpgradeBlockedReason() []string {
	return []string{
		string(AgentUpgradeBlockedReasonAgentNoHeartbeat),
		string(AgentUpgradeBlockedReasonAlreadyOnLatestVersion),
		string(AgentUpgradeBlockedReasonDistroIsNotReported),
		string(AgentUpgradeBlockedReasonDistroNotSupportedForUpgrade),
		string(AgentUpgradeBlockedReasonIncompatibleApplianceVersion),
		string(AgentUpgradeBlockedReasonInvalidAgentVersion),
		string(AgentUpgradeBlockedReasonInvalidDriverVersion),
		string(AgentUpgradeBlockedReasonMissingUpgradePath),
		string(AgentUpgradeBlockedReasonNotProtected),
		string(AgentUpgradeBlockedReasonProcessServerNoHeartbeat),
		string(AgentUpgradeBlockedReasonRcmProxyNoHeartbeat),
		string(AgentUpgradeBlockedReasonRebootRequired),
		string(AgentUpgradeBlockedReasonUnknown),
		string(AgentUpgradeBlockedReasonUnsupportedProtectionScenario),
	}
}

type AutoProtectionOfDataDisk string

const (
	AutoProtectionOfDataDiskDisabled AutoProtectionOfDataDisk = "Disabled"
	AutoProtectionOfDataDiskEnabled  AutoProtectionOfDataDisk = "Enabled"
)

func PossibleValuesForAutoProtectionOfDataDisk() []string {
	return []string{
		string(AutoProtectionOfDataDiskDisabled),
		string(AutoProtectionOfDataDiskEnabled),
	}
}

type DisableProtectionReason string

const (
	DisableProtectionReasonMigrationComplete DisableProtectionReason = "MigrationComplete"
	DisableProtectionReasonNotSpecified      DisableProtectionReason = "NotSpecified"
)

func PossibleValuesForDisableProtectionReason() []string {
	return []string{
		string(DisableProtectionReasonMigrationComplete),
		string(DisableProtectionReasonNotSpecified),
	}
}

type DiskAccountType string

const (
	DiskAccountTypePremiumLRS     DiskAccountType = "Premium_LRS"
	DiskAccountTypeStandardLRS    DiskAccountType = "Standard_LRS"
	DiskAccountTypeStandardSSDLRS DiskAccountType = "StandardSSD_LRS"
)

func PossibleValuesForDiskAccountType() []string {
	return []string{
		string(DiskAccountTypePremiumLRS),
		string(DiskAccountTypeStandardLRS),
		string(DiskAccountTypeStandardSSDLRS),
	}
}

type DiskReplicationProgressHealth string

const (
	DiskReplicationProgressHealthInProgress   DiskReplicationProgressHealth = "InProgress"
	DiskReplicationProgressHealthNoProgress   DiskReplicationProgressHealth = "NoProgress"
	DiskReplicationProgressHealthNone         DiskReplicationProgressHealth = "None"
	DiskReplicationProgressHealthQueued       DiskReplicationProgressHealth = "Queued"
	DiskReplicationProgressHealthSlowProgress DiskReplicationProgressHealth = "SlowProgress"
)

func PossibleValuesForDiskReplicationProgressHealth() []string {
	return []string{
		string(DiskReplicationProgressHealthInProgress),
		string(DiskReplicationProgressHealthNoProgress),
		string(DiskReplicationProgressHealthNone),
		string(DiskReplicationProgressHealthQueued),
		string(DiskReplicationProgressHealthSlowProgress),
	}
}

type EthernetAddressType string

const (
	EthernetAddressTypeDynamic EthernetAddressType = "Dynamic"
	EthernetAddressTypeStatic  EthernetAddressType = "Static"
)

func PossibleValuesForEthernetAddressType() []string {
	return []string{
		string(EthernetAddressTypeDynamic),
		string(EthernetAddressTypeStatic),
	}
}

type HealthErrorCustomerResolvability string

const (
	HealthErrorCustomerResolvabilityAllowed    HealthErrorCustomerResolvability = "Allowed"
	HealthErrorCustomerResolvabilityNotAllowed HealthErrorCustomerResolvability = "NotAllowed"
)

func PossibleValuesForHealthErrorCustomerResolvability() []string {
	return []string{
		string(HealthErrorCustomerResolvabilityAllowed),
		string(HealthErrorCustomerResolvabilityNotAllowed),
	}
}

type InMageRcmFailbackRecoveryPointType string

const (
	InMageRcmFailbackRecoveryPointTypeApplicationConsistent InMageRcmFailbackRecoveryPointType = "ApplicationConsistent"
	InMageRcmFailbackRecoveryPointTypeCrashConsistent       InMageRcmFailbackRecoveryPointType = "CrashConsistent"
)

func PossibleValuesForInMageRcmFailbackRecoveryPointType() []string {
	return []string{
		string(InMageRcmFailbackRecoveryPointTypeApplicationConsistent),
		string(InMageRcmFailbackRecoveryPointTypeCrashConsistent),
	}
}

type LicenseType string

const (
	LicenseTypeNoLicenseType LicenseType = "NoLicenseType"
	LicenseTypeNotSpecified  LicenseType = "NotSpecified"
	LicenseTypeWindowsServer LicenseType = "WindowsServer"
)

func PossibleValuesForLicenseType() []string {
	return []string{
		string(LicenseTypeNoLicenseType),
		string(LicenseTypeNotSpecified),
		string(LicenseTypeWindowsServer),
	}
}

type MobilityAgentUpgradeState string

const (
	MobilityAgentUpgradeStateCommit    MobilityAgentUpgradeState = "Commit"
	MobilityAgentUpgradeStateCompleted MobilityAgentUpgradeState = "Completed"
	MobilityAgentUpgradeStateNone      MobilityAgentUpgradeState = "None"
	MobilityAgentUpgradeStateStarted   MobilityAgentUpgradeState = "Started"
)

func PossibleValuesForMobilityAgentUpgradeState() []string {
	return []string{
		string(MobilityAgentUpgradeStateCommit),
		string(MobilityAgentUpgradeStateCompleted),
		string(MobilityAgentUpgradeStateNone),
		string(MobilityAgentUpgradeStateStarted),
	}
}

type MultiVMGroupCreateOption string

const (
	MultiVMGroupCreateOptionAutoCreated   MultiVMGroupCreateOption = "AutoCreated"
	MultiVMGroupCreateOptionUserSpecified MultiVMGroupCreateOption = "UserSpecified"
)

func PossibleValuesForMultiVMGroupCreateOption() []string {
	return []string{
		string(MultiVMGroupCreateOptionAutoCreated),
		string(MultiVMGroupCreateOptionUserSpecified),
	}
}

type PlannedFailoverStatus string

const (
	PlannedFailoverStatusCancelled PlannedFailoverStatus = "Cancelled"
	PlannedFailoverStatusFailed    PlannedFailoverStatus = "Failed"
	PlannedFailoverStatusSucceeded PlannedFailoverStatus = "Succeeded"
	PlannedFailoverStatusUnknown   PlannedFailoverStatus = "Unknown"
)

func PossibleValuesForPlannedFailoverStatus() []string {
	return []string{
		string(PlannedFailoverStatusCancelled),
		string(PlannedFailoverStatusFailed),
		string(PlannedFailoverStatusSucceeded),
		string(PlannedFailoverStatusUnknown),
	}
}

type RecoveryPointType string

const (
	RecoveryPointTypeCustom     RecoveryPointType = "Custom"
	RecoveryPointTypeLatestTag  RecoveryPointType = "LatestTag"
	RecoveryPointTypeLatestTime RecoveryPointType = "LatestTime"
)

func PossibleValuesForRecoveryPointType() []string {
	return []string{
		string(RecoveryPointTypeCustom),
		string(RecoveryPointTypeLatestTag),
		string(RecoveryPointTypeLatestTime),
	}
}

type ResyncState string

const (
	ResyncStateNone                         ResyncState = "None"
	ResyncStatePreparedForResynchronization ResyncState = "PreparedForResynchronization"
	ResyncStateStartedResynchronization     ResyncState = "StartedResynchronization"
)

func PossibleValuesForResyncState() []string {
	return []string{
		string(ResyncStateNone),
		string(ResyncStatePreparedForResynchronization),
		string(ResyncStateStartedResynchronization),
	}
}

type SqlServerLicenseType string

const (
	SqlServerLicenseTypeAHUB          SqlServerLicenseType = "AHUB"
	SqlServerLicenseTypeNoLicenseType SqlServerLicenseType = "NoLicenseType"
	SqlServerLicenseTypeNotSpecified  SqlServerLicenseType = "NotSpecified"
	SqlServerLicenseTypePAYG          SqlServerLicenseType = "PAYG"
)

func PossibleValuesForSqlServerLicenseType() []string {
	return []string{
		string(SqlServerLicenseTypeAHUB),
		string(SqlServerLicenseTypeNoLicenseType),
		string(SqlServerLicenseTypeNotSpecified),
		string(SqlServerLicenseTypePAYG),
	}
}

type VMEncryptionType string

const (
	VMEncryptionTypeNotEncrypted     VMEncryptionType = "NotEncrypted"
	VMEncryptionTypeOnePassEncrypted VMEncryptionType = "OnePassEncrypted"
	VMEncryptionTypeTwoPassEncrypted VMEncryptionType = "TwoPassEncrypted"
)

func PossibleValuesForVMEncryptionType() []string {
	return []string{
		string(VMEncryptionTypeNotEncrypted),
		string(VMEncryptionTypeOnePassEncrypted),
		string(VMEncryptionTypeTwoPassEncrypted),
	}
}

type VMReplicationProgressHealth string

const (
	VMReplicationProgressHealthInProgress   VMReplicationProgressHealth = "InProgress"
	VMReplicationProgressHealthNoProgress   VMReplicationProgressHealth = "NoProgress"
	VMReplicationProgressHealthNone         VMReplicationProgressHealth = "None"
	VMReplicationProgressHealthSlowProgress VMReplicationProgressHealth = "SlowProgress"
)

func PossibleValuesForVMReplicationProgressHealth() []string {
	return []string{
		string(VMReplicationProgressHealthInProgress),
		string(VMReplicationProgressHealthNoProgress),
		string(VMReplicationProgressHealthNone),
		string(VMReplicationProgressHealthSlowProgress),
	}
}
