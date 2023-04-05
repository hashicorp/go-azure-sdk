package integrationserviceenvironmentmanagedapi

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
