package endpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationType string

const (
	AuthenticationTypeIdentityBased AuthenticationType = "IdentityBased"
	AuthenticationTypeKeyBased      AuthenticationType = "KeyBased"
)

func PossibleValuesForAuthenticationType() []string {
	return []string{
		string(AuthenticationTypeIdentityBased),
		string(AuthenticationTypeKeyBased),
	}
}

type EndpointProvisioningState string

const (
	EndpointProvisioningStateCanceled     EndpointProvisioningState = "Canceled"
	EndpointProvisioningStateDeleted      EndpointProvisioningState = "Deleted"
	EndpointProvisioningStateDeleting     EndpointProvisioningState = "Deleting"
	EndpointProvisioningStateDisabled     EndpointProvisioningState = "Disabled"
	EndpointProvisioningStateFailed       EndpointProvisioningState = "Failed"
	EndpointProvisioningStateMoving       EndpointProvisioningState = "Moving"
	EndpointProvisioningStateProvisioning EndpointProvisioningState = "Provisioning"
	EndpointProvisioningStateRestoring    EndpointProvisioningState = "Restoring"
	EndpointProvisioningStateSucceeded    EndpointProvisioningState = "Succeeded"
	EndpointProvisioningStateSuspending   EndpointProvisioningState = "Suspending"
	EndpointProvisioningStateWarning      EndpointProvisioningState = "Warning"
)

func PossibleValuesForEndpointProvisioningState() []string {
	return []string{
		string(EndpointProvisioningStateCanceled),
		string(EndpointProvisioningStateDeleted),
		string(EndpointProvisioningStateDeleting),
		string(EndpointProvisioningStateDisabled),
		string(EndpointProvisioningStateFailed),
		string(EndpointProvisioningStateMoving),
		string(EndpointProvisioningStateProvisioning),
		string(EndpointProvisioningStateRestoring),
		string(EndpointProvisioningStateSucceeded),
		string(EndpointProvisioningStateSuspending),
		string(EndpointProvisioningStateWarning),
	}
}

type EndpointType string

const (
	EndpointTypeEventGrid  EndpointType = "EventGrid"
	EndpointTypeEventHub   EndpointType = "EventHub"
	EndpointTypeServiceBus EndpointType = "ServiceBus"
)

func PossibleValuesForEndpointType() []string {
	return []string{
		string(EndpointTypeEventGrid),
		string(EndpointTypeEventHub),
		string(EndpointTypeServiceBus),
	}
}
