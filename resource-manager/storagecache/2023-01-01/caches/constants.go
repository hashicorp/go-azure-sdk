package caches

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainJoinedType string

const (
	DomainJoinedTypeError DomainJoinedType = "Error"
	DomainJoinedTypeNo    DomainJoinedType = "No"
	DomainJoinedTypeYes   DomainJoinedType = "Yes"
)

func PossibleValuesForDomainJoinedType() []string {
	return []string{
		string(DomainJoinedTypeError),
		string(DomainJoinedTypeNo),
		string(DomainJoinedTypeYes),
	}
}

type FirmwareStatusType string

const (
	FirmwareStatusTypeAvailable   FirmwareStatusType = "available"
	FirmwareStatusTypeUnavailable FirmwareStatusType = "unavailable"
)

func PossibleValuesForFirmwareStatusType() []string {
	return []string{
		string(FirmwareStatusTypeAvailable),
		string(FirmwareStatusTypeUnavailable),
	}
}

type HealthStateType string

const (
	HealthStateTypeDegraded      HealthStateType = "Degraded"
	HealthStateTypeDown          HealthStateType = "Down"
	HealthStateTypeFlushing      HealthStateType = "Flushing"
	HealthStateTypeHealthy       HealthStateType = "Healthy"
	HealthStateTypeStartFailed   HealthStateType = "StartFailed"
	HealthStateTypeStopped       HealthStateType = "Stopped"
	HealthStateTypeStopping      HealthStateType = "Stopping"
	HealthStateTypeTransitioning HealthStateType = "Transitioning"
	HealthStateTypeUnknown       HealthStateType = "Unknown"
	HealthStateTypeUpgradeFailed HealthStateType = "UpgradeFailed"
	HealthStateTypeUpgrading     HealthStateType = "Upgrading"
	HealthStateTypeWaitingForKey HealthStateType = "WaitingForKey"
)

func PossibleValuesForHealthStateType() []string {
	return []string{
		string(HealthStateTypeDegraded),
		string(HealthStateTypeDown),
		string(HealthStateTypeFlushing),
		string(HealthStateTypeHealthy),
		string(HealthStateTypeStartFailed),
		string(HealthStateTypeStopped),
		string(HealthStateTypeStopping),
		string(HealthStateTypeTransitioning),
		string(HealthStateTypeUnknown),
		string(HealthStateTypeUpgradeFailed),
		string(HealthStateTypeUpgrading),
		string(HealthStateTypeWaitingForKey),
	}
}

type NfsAccessRuleAccess string

const (
	NfsAccessRuleAccessNo NfsAccessRuleAccess = "no"
	NfsAccessRuleAccessRo NfsAccessRuleAccess = "ro"
	NfsAccessRuleAccessRw NfsAccessRuleAccess = "rw"
)

func PossibleValuesForNfsAccessRuleAccess() []string {
	return []string{
		string(NfsAccessRuleAccessNo),
		string(NfsAccessRuleAccessRo),
		string(NfsAccessRuleAccessRw),
	}
}

type NfsAccessRuleScope string

const (
	NfsAccessRuleScopeDefault NfsAccessRuleScope = "default"
	NfsAccessRuleScopeHost    NfsAccessRuleScope = "host"
	NfsAccessRuleScopeNetwork NfsAccessRuleScope = "network"
)

func PossibleValuesForNfsAccessRuleScope() []string {
	return []string{
		string(NfsAccessRuleScopeDefault),
		string(NfsAccessRuleScopeHost),
		string(NfsAccessRuleScopeNetwork),
	}
}

type PrimingJobState string

const (
	PrimingJobStateComplete PrimingJobState = "Complete"
	PrimingJobStatePaused   PrimingJobState = "Paused"
	PrimingJobStateQueued   PrimingJobState = "Queued"
	PrimingJobStateRunning  PrimingJobState = "Running"
)

func PossibleValuesForPrimingJobState() []string {
	return []string{
		string(PrimingJobStateComplete),
		string(PrimingJobStatePaused),
		string(PrimingJobStateQueued),
		string(PrimingJobStateRunning),
	}
}

type ProvisioningStateType string

const (
	ProvisioningStateTypeCancelled ProvisioningStateType = "Cancelled"
	ProvisioningStateTypeCreating  ProvisioningStateType = "Creating"
	ProvisioningStateTypeDeleting  ProvisioningStateType = "Deleting"
	ProvisioningStateTypeFailed    ProvisioningStateType = "Failed"
	ProvisioningStateTypeSucceeded ProvisioningStateType = "Succeeded"
	ProvisioningStateTypeUpdating  ProvisioningStateType = "Updating"
)

func PossibleValuesForProvisioningStateType() []string {
	return []string{
		string(ProvisioningStateTypeCancelled),
		string(ProvisioningStateTypeCreating),
		string(ProvisioningStateTypeDeleting),
		string(ProvisioningStateTypeFailed),
		string(ProvisioningStateTypeSucceeded),
		string(ProvisioningStateTypeUpdating),
	}
}

type UsernameDownloadedType string

const (
	UsernameDownloadedTypeError UsernameDownloadedType = "Error"
	UsernameDownloadedTypeNo    UsernameDownloadedType = "No"
	UsernameDownloadedTypeYes   UsernameDownloadedType = "Yes"
)

func PossibleValuesForUsernameDownloadedType() []string {
	return []string{
		string(UsernameDownloadedTypeError),
		string(UsernameDownloadedTypeNo),
		string(UsernameDownloadedTypeYes),
	}
}

type UsernameSource string

const (
	UsernameSourceAD   UsernameSource = "AD"
	UsernameSourceFile UsernameSource = "File"
	UsernameSourceLDAP UsernameSource = "LDAP"
	UsernameSourceNone UsernameSource = "None"
)

func PossibleValuesForUsernameSource() []string {
	return []string{
		string(UsernameSourceAD),
		string(UsernameSourceFile),
		string(UsernameSourceLDAP),
		string(UsernameSourceNone),
	}
}
