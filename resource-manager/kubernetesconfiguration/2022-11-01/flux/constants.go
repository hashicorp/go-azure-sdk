package flux

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FluxComplianceState string

const (
	FluxComplianceStateCompliant            FluxComplianceState = "Compliant"
	FluxComplianceStateNonNegativeCompliant FluxComplianceState = "Non-Compliant"
	FluxComplianceStatePending              FluxComplianceState = "Pending"
	FluxComplianceStateSuspended            FluxComplianceState = "Suspended"
	FluxComplianceStateUnknown              FluxComplianceState = "Unknown"
)

func PossibleValuesForFluxComplianceState() []string {
	return []string{
		string(FluxComplianceStateCompliant),
		string(FluxComplianceStateNonNegativeCompliant),
		string(FluxComplianceStatePending),
		string(FluxComplianceStateSuspended),
		string(FluxComplianceStateUnknown),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type ScopeType string

const (
	ScopeTypeCluster   ScopeType = "cluster"
	ScopeTypeNamespace ScopeType = "namespace"
)

func PossibleValuesForScopeType() []string {
	return []string{
		string(ScopeTypeCluster),
		string(ScopeTypeNamespace),
	}
}

type SourceKindType string

const (
	SourceKindTypeAzureBlob     SourceKindType = "AzureBlob"
	SourceKindTypeBucket        SourceKindType = "Bucket"
	SourceKindTypeGitRepository SourceKindType = "GitRepository"
)

func PossibleValuesForSourceKindType() []string {
	return []string{
		string(SourceKindTypeAzureBlob),
		string(SourceKindTypeBucket),
		string(SourceKindTypeGitRepository),
	}
}
