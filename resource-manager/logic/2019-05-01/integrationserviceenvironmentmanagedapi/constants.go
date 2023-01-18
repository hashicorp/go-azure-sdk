package integrationserviceenvironmentmanagedapi

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiDeploymentParameterVisibility string

const (
	ApiDeploymentParameterVisibilityDefault      ApiDeploymentParameterVisibility = "Default"
	ApiDeploymentParameterVisibilityInternal     ApiDeploymentParameterVisibility = "Internal"
	ApiDeploymentParameterVisibilityNotSpecified ApiDeploymentParameterVisibility = "NotSpecified"
)

func PossibleValuesForApiDeploymentParameterVisibility() []string {
	return []string{
		string(ApiDeploymentParameterVisibilityDefault),
		string(ApiDeploymentParameterVisibilityInternal),
		string(ApiDeploymentParameterVisibilityNotSpecified),
	}
}

func parseApiDeploymentParameterVisibility(input string) (*ApiDeploymentParameterVisibility, error) {
	vals := map[string]ApiDeploymentParameterVisibility{
		"default":      ApiDeploymentParameterVisibilityDefault,
		"internal":     ApiDeploymentParameterVisibilityInternal,
		"notspecified": ApiDeploymentParameterVisibilityNotSpecified,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiDeploymentParameterVisibility(input)
	return &out, nil
}

type ApiTier string

const (
	ApiTierEnterprise   ApiTier = "Enterprise"
	ApiTierNotSpecified ApiTier = "NotSpecified"
	ApiTierPremium      ApiTier = "Premium"
	ApiTierStandard     ApiTier = "Standard"
)

func PossibleValuesForApiTier() []string {
	return []string{
		string(ApiTierEnterprise),
		string(ApiTierNotSpecified),
		string(ApiTierPremium),
		string(ApiTierStandard),
	}
}

func parseApiTier(input string) (*ApiTier, error) {
	vals := map[string]ApiTier{
		"enterprise":   ApiTierEnterprise,
		"notspecified": ApiTierNotSpecified,
		"premium":      ApiTierPremium,
		"standard":     ApiTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiTier(input)
	return &out, nil
}

type ApiType string

const (
	ApiTypeNotSpecified ApiType = "NotSpecified"
	ApiTypeRest         ApiType = "Rest"
	ApiTypeSoap         ApiType = "Soap"
)

func PossibleValuesForApiType() []string {
	return []string{
		string(ApiTypeNotSpecified),
		string(ApiTypeRest),
		string(ApiTypeSoap),
	}
}

func parseApiType(input string) (*ApiType, error) {
	vals := map[string]ApiType{
		"notspecified": ApiTypeNotSpecified,
		"rest":         ApiTypeRest,
		"soap":         ApiTypeSoap,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiType(input)
	return &out, nil
}

type WorkflowProvisioningState string

const (
	WorkflowProvisioningStateAccepted      WorkflowProvisioningState = "Accepted"
	WorkflowProvisioningStateCanceled      WorkflowProvisioningState = "Canceled"
	WorkflowProvisioningStateCompleted     WorkflowProvisioningState = "Completed"
	WorkflowProvisioningStateCreated       WorkflowProvisioningState = "Created"
	WorkflowProvisioningStateCreating      WorkflowProvisioningState = "Creating"
	WorkflowProvisioningStateDeleted       WorkflowProvisioningState = "Deleted"
	WorkflowProvisioningStateDeleting      WorkflowProvisioningState = "Deleting"
	WorkflowProvisioningStateFailed        WorkflowProvisioningState = "Failed"
	WorkflowProvisioningStateInProgress    WorkflowProvisioningState = "InProgress"
	WorkflowProvisioningStateMoving        WorkflowProvisioningState = "Moving"
	WorkflowProvisioningStateNotSpecified  WorkflowProvisioningState = "NotSpecified"
	WorkflowProvisioningStatePending       WorkflowProvisioningState = "Pending"
	WorkflowProvisioningStateReady         WorkflowProvisioningState = "Ready"
	WorkflowProvisioningStateRegistered    WorkflowProvisioningState = "Registered"
	WorkflowProvisioningStateRegistering   WorkflowProvisioningState = "Registering"
	WorkflowProvisioningStateRenewing      WorkflowProvisioningState = "Renewing"
	WorkflowProvisioningStateRunning       WorkflowProvisioningState = "Running"
	WorkflowProvisioningStateSucceeded     WorkflowProvisioningState = "Succeeded"
	WorkflowProvisioningStateUnregistered  WorkflowProvisioningState = "Unregistered"
	WorkflowProvisioningStateUnregistering WorkflowProvisioningState = "Unregistering"
	WorkflowProvisioningStateUpdating      WorkflowProvisioningState = "Updating"
	WorkflowProvisioningStateWaiting       WorkflowProvisioningState = "Waiting"
)

func PossibleValuesForWorkflowProvisioningState() []string {
	return []string{
		string(WorkflowProvisioningStateAccepted),
		string(WorkflowProvisioningStateCanceled),
		string(WorkflowProvisioningStateCompleted),
		string(WorkflowProvisioningStateCreated),
		string(WorkflowProvisioningStateCreating),
		string(WorkflowProvisioningStateDeleted),
		string(WorkflowProvisioningStateDeleting),
		string(WorkflowProvisioningStateFailed),
		string(WorkflowProvisioningStateInProgress),
		string(WorkflowProvisioningStateMoving),
		string(WorkflowProvisioningStateNotSpecified),
		string(WorkflowProvisioningStatePending),
		string(WorkflowProvisioningStateReady),
		string(WorkflowProvisioningStateRegistered),
		string(WorkflowProvisioningStateRegistering),
		string(WorkflowProvisioningStateRenewing),
		string(WorkflowProvisioningStateRunning),
		string(WorkflowProvisioningStateSucceeded),
		string(WorkflowProvisioningStateUnregistered),
		string(WorkflowProvisioningStateUnregistering),
		string(WorkflowProvisioningStateUpdating),
		string(WorkflowProvisioningStateWaiting),
	}
}

func parseWorkflowProvisioningState(input string) (*WorkflowProvisioningState, error) {
	vals := map[string]WorkflowProvisioningState{
		"accepted":      WorkflowProvisioningStateAccepted,
		"canceled":      WorkflowProvisioningStateCanceled,
		"completed":     WorkflowProvisioningStateCompleted,
		"created":       WorkflowProvisioningStateCreated,
		"creating":      WorkflowProvisioningStateCreating,
		"deleted":       WorkflowProvisioningStateDeleted,
		"deleting":      WorkflowProvisioningStateDeleting,
		"failed":        WorkflowProvisioningStateFailed,
		"inprogress":    WorkflowProvisioningStateInProgress,
		"moving":        WorkflowProvisioningStateMoving,
		"notspecified":  WorkflowProvisioningStateNotSpecified,
		"pending":       WorkflowProvisioningStatePending,
		"ready":         WorkflowProvisioningStateReady,
		"registered":    WorkflowProvisioningStateRegistered,
		"registering":   WorkflowProvisioningStateRegistering,
		"renewing":      WorkflowProvisioningStateRenewing,
		"running":       WorkflowProvisioningStateRunning,
		"succeeded":     WorkflowProvisioningStateSucceeded,
		"unregistered":  WorkflowProvisioningStateUnregistered,
		"unregistering": WorkflowProvisioningStateUnregistering,
		"updating":      WorkflowProvisioningStateUpdating,
		"waiting":       WorkflowProvisioningStateWaiting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WorkflowProvisioningState(input)
	return &out, nil
}

type WsdlImportMethod string

const (
	WsdlImportMethodNotSpecified    WsdlImportMethod = "NotSpecified"
	WsdlImportMethodSoapPassThrough WsdlImportMethod = "SoapPassThrough"
	WsdlImportMethodSoapToRest      WsdlImportMethod = "SoapToRest"
)

func PossibleValuesForWsdlImportMethod() []string {
	return []string{
		string(WsdlImportMethodNotSpecified),
		string(WsdlImportMethodSoapPassThrough),
		string(WsdlImportMethodSoapToRest),
	}
}

func parseWsdlImportMethod(input string) (*WsdlImportMethod, error) {
	vals := map[string]WsdlImportMethod{
		"notspecified":    WsdlImportMethodNotSpecified,
		"soappassthrough": WsdlImportMethodSoapPassThrough,
		"soaptorest":      WsdlImportMethodSoapToRest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WsdlImportMethod(input)
	return &out, nil
}
