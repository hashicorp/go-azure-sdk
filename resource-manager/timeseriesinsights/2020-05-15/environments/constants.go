package environments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnvironmentKind string

const (
	EnvironmentKindGenOne EnvironmentKind = "Gen1"
	EnvironmentKindGenTwo EnvironmentKind = "Gen2"
)

func PossibleValuesForEnvironmentKind() []string {
	return []string{
		string(EnvironmentKindGenOne),
		string(EnvironmentKindGenTwo),
	}
}

type IngressState string

const (
	IngressStateDisabled IngressState = "Disabled"
	IngressStatePaused   IngressState = "Paused"
	IngressStateReady    IngressState = "Ready"
	IngressStateRunning  IngressState = "Running"
	IngressStateUnknown  IngressState = "Unknown"
)

func PossibleValuesForIngressState() []string {
	return []string{
		string(IngressStateDisabled),
		string(IngressStatePaused),
		string(IngressStateReady),
		string(IngressStateRunning),
		string(IngressStateUnknown),
	}
}

type Kind string

const (
	KindGenOne Kind = "Gen1"
	KindGenTwo Kind = "Gen2"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindGenOne),
		string(KindGenTwo),
	}
}

type PropertyType string

const (
	PropertyTypeString PropertyType = "String"
)

func PossibleValuesForPropertyType() []string {
	return []string{
		string(PropertyTypeString),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type SkuName string

const (
	SkuNameLOne SkuName = "L1"
	SkuNamePOne SkuName = "P1"
	SkuNameSOne SkuName = "S1"
	SkuNameSTwo SkuName = "S2"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameLOne),
		string(SkuNamePOne),
		string(SkuNameSOne),
		string(SkuNameSTwo),
	}
}

type StorageLimitExceededBehavior string

const (
	StorageLimitExceededBehaviorPauseIngress StorageLimitExceededBehavior = "PauseIngress"
	StorageLimitExceededBehaviorPurgeOldData StorageLimitExceededBehavior = "PurgeOldData"
)

func PossibleValuesForStorageLimitExceededBehavior() []string {
	return []string{
		string(StorageLimitExceededBehaviorPauseIngress),
		string(StorageLimitExceededBehaviorPurgeOldData),
	}
}

type WarmStoragePropertiesState string

const (
	WarmStoragePropertiesStateError   WarmStoragePropertiesState = "Error"
	WarmStoragePropertiesStateOk      WarmStoragePropertiesState = "Ok"
	WarmStoragePropertiesStateUnknown WarmStoragePropertiesState = "Unknown"
)

func PossibleValuesForWarmStoragePropertiesState() []string {
	return []string{
		string(WarmStoragePropertiesStateError),
		string(WarmStoragePropertiesStateOk),
		string(WarmStoragePropertiesStateUnknown),
	}
}
