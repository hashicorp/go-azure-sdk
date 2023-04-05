package lab

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionType string

const (
	ConnectionTypeNone    ConnectionType = "None"
	ConnectionTypePrivate ConnectionType = "Private"
	ConnectionTypePublic  ConnectionType = "Public"
)

func PossibleValuesForConnectionType() []string {
	return []string{
		string(ConnectionTypeNone),
		string(ConnectionTypePrivate),
		string(ConnectionTypePublic),
	}
}

type CreateOption string

const (
	CreateOptionImage      CreateOption = "Image"
	CreateOptionTemplateVM CreateOption = "TemplateVM"
)

func PossibleValuesForCreateOption() []string {
	return []string{
		string(CreateOptionImage),
		string(CreateOptionTemplateVM),
	}
}

type EnableState string

const (
	EnableStateDisabled EnableState = "Disabled"
	EnableStateEnabled  EnableState = "Enabled"
)

func PossibleValuesForEnableState() []string {
	return []string{
		string(EnableStateDisabled),
		string(EnableStateEnabled),
	}
}

type LabState string

const (
	LabStateDraft      LabState = "Draft"
	LabStatePublished  LabState = "Published"
	LabStatePublishing LabState = "Publishing"
	LabStateScaling    LabState = "Scaling"
	LabStateSyncing    LabState = "Syncing"
)

func PossibleValuesForLabState() []string {
	return []string{
		string(LabStateDraft),
		string(LabStatePublished),
		string(LabStatePublishing),
		string(LabStateScaling),
		string(LabStateSyncing),
	}
}

type OsType string

const (
	OsTypeLinux   OsType = "Linux"
	OsTypeWindows OsType = "Windows"
)

func PossibleValuesForOsType() []string {
	return []string{
		string(OsTypeLinux),
		string(OsTypeWindows),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateLocked    ProvisioningState = "Locked"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateLocked),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type ShutdownOnIdleMode string

const (
	ShutdownOnIdleModeLowUsage    ShutdownOnIdleMode = "LowUsage"
	ShutdownOnIdleModeNone        ShutdownOnIdleMode = "None"
	ShutdownOnIdleModeUserAbsence ShutdownOnIdleMode = "UserAbsence"
)

func PossibleValuesForShutdownOnIdleMode() []string {
	return []string{
		string(ShutdownOnIdleModeLowUsage),
		string(ShutdownOnIdleModeNone),
		string(ShutdownOnIdleModeUserAbsence),
	}
}

type SkuTier string

const (
	SkuTierBasic    SkuTier = "Basic"
	SkuTierFree     SkuTier = "Free"
	SkuTierPremium  SkuTier = "Premium"
	SkuTierStandard SkuTier = "Standard"
)

func PossibleValuesForSkuTier() []string {
	return []string{
		string(SkuTierBasic),
		string(SkuTierFree),
		string(SkuTierPremium),
		string(SkuTierStandard),
	}
}
