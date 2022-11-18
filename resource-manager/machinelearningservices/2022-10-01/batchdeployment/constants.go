package batchdeployment

import "strings"

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

func parseBatchLoggingLevel(input string) (*BatchLoggingLevel, error) {
	vals := map[string]BatchLoggingLevel{
		"debug":   BatchLoggingLevelDebug,
		"info":    BatchLoggingLevelInfo,
		"warning": BatchLoggingLevelWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BatchLoggingLevel(input)
	return &out, nil
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

func parseBatchOutputAction(input string) (*BatchOutputAction, error) {
	vals := map[string]BatchOutputAction{
		"appendrow":   BatchOutputActionAppendRow,
		"summaryonly": BatchOutputActionSummaryOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BatchOutputAction(input)
	return &out, nil
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

func parseDeploymentProvisioningState(input string) (*DeploymentProvisioningState, error) {
	vals := map[string]DeploymentProvisioningState{
		"canceled":  DeploymentProvisioningStateCanceled,
		"creating":  DeploymentProvisioningStateCreating,
		"deleting":  DeploymentProvisioningStateDeleting,
		"failed":    DeploymentProvisioningStateFailed,
		"scaling":   DeploymentProvisioningStateScaling,
		"succeeded": DeploymentProvisioningStateSucceeded,
		"updating":  DeploymentProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentProvisioningState(input)
	return &out, nil
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

func parseReferenceType(input string) (*ReferenceType, error) {
	vals := map[string]ReferenceType{
		"datapath":   ReferenceTypeDataPath,
		"id":         ReferenceTypeId,
		"outputpath": ReferenceTypeOutputPath,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ReferenceType(input)
	return &out, nil
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

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":    SkuTierBasic,
		"free":     SkuTierFree,
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}
