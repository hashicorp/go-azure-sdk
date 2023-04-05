package deploymentscripts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CleanupOptions string

const (
	CleanupOptionsAlways       CleanupOptions = "Always"
	CleanupOptionsOnExpiration CleanupOptions = "OnExpiration"
	CleanupOptionsOnSuccess    CleanupOptions = "OnSuccess"
)

func PossibleValuesForCleanupOptions() []string {
	return []string{
		string(CleanupOptionsAlways),
		string(CleanupOptionsOnExpiration),
		string(CleanupOptionsOnSuccess),
	}
}

type ManagedServiceIdentityType string

const (
	ManagedServiceIdentityTypeUserAssigned ManagedServiceIdentityType = "UserAssigned"
)

func PossibleValuesForManagedServiceIdentityType() []string {
	return []string{
		string(ManagedServiceIdentityTypeUserAssigned),
	}
}

type ScriptProvisioningState string

const (
	ScriptProvisioningStateCanceled              ScriptProvisioningState = "Canceled"
	ScriptProvisioningStateCreating              ScriptProvisioningState = "Creating"
	ScriptProvisioningStateFailed                ScriptProvisioningState = "Failed"
	ScriptProvisioningStateProvisioningResources ScriptProvisioningState = "ProvisioningResources"
	ScriptProvisioningStateRunning               ScriptProvisioningState = "Running"
	ScriptProvisioningStateSucceeded             ScriptProvisioningState = "Succeeded"
)

func PossibleValuesForScriptProvisioningState() []string {
	return []string{
		string(ScriptProvisioningStateCanceled),
		string(ScriptProvisioningStateCreating),
		string(ScriptProvisioningStateFailed),
		string(ScriptProvisioningStateProvisioningResources),
		string(ScriptProvisioningStateRunning),
		string(ScriptProvisioningStateSucceeded),
	}
}

type ScriptType string

const (
	ScriptTypeAzureCLI        ScriptType = "AzureCLI"
	ScriptTypeAzurePowerShell ScriptType = "AzurePowerShell"
)

func PossibleValuesForScriptType() []string {
	return []string{
		string(ScriptTypeAzureCLI),
		string(ScriptTypeAzurePowerShell),
	}
}
