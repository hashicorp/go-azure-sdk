package batchdeployment

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BatchLoggingLevel string

const (
	BatchLoggingLevelDebug   BatchLoggingLevel = "Debug"
	BatchLoggingLevelInfo    BatchLoggingLevel = "Info"
	BatchLoggingLevelWarning BatchLoggingLevel = "Warning"
)

func PossibleValuesForBatchLoggingLevel() []string {
	return []string{
		string(BatchLoggingLevelDebug),
		string(BatchLoggingLevelInfo),
		string(BatchLoggingLevelWarning),
	}
}

type BatchOutputAction string

const (
	BatchOutputActionAppendRow   BatchOutputAction = "AppendRow"
	BatchOutputActionSummaryOnly BatchOutputAction = "SummaryOnly"
)

func PossibleValuesForBatchOutputAction() []string {
	return []string{
		string(BatchOutputActionAppendRow),
		string(BatchOutputActionSummaryOnly),
	}
}

type DeploymentProvisioningState string

const (
	DeploymentProvisioningStateCanceled  DeploymentProvisioningState = "Canceled"
	DeploymentProvisioningStateCreating  DeploymentProvisioningState = "Creating"
	DeploymentProvisioningStateDeleting  DeploymentProvisioningState = "Deleting"
	DeploymentProvisioningStateFailed    DeploymentProvisioningState = "Failed"
	DeploymentProvisioningStateScaling   DeploymentProvisioningState = "Scaling"
	DeploymentProvisioningStateSucceeded DeploymentProvisioningState = "Succeeded"
	DeploymentProvisioningStateUpdating  DeploymentProvisioningState = "Updating"
)

func PossibleValuesForDeploymentProvisioningState() []string {
	return []string{
		string(DeploymentProvisioningStateCanceled),
		string(DeploymentProvisioningStateCreating),
		string(DeploymentProvisioningStateDeleting),
		string(DeploymentProvisioningStateFailed),
		string(DeploymentProvisioningStateScaling),
		string(DeploymentProvisioningStateSucceeded),
		string(DeploymentProvisioningStateUpdating),
	}
}

type ReferenceType string

const (
	ReferenceTypeDataPath   ReferenceType = "DataPath"
	ReferenceTypeId         ReferenceType = "Id"
	ReferenceTypeOutputPath ReferenceType = "OutputPath"
)

func PossibleValuesForReferenceType() []string {
	return []string{
		string(ReferenceTypeDataPath),
		string(ReferenceTypeId),
		string(ReferenceTypeOutputPath),
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
