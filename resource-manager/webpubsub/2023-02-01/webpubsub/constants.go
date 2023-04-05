package webpubsub

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ACLAction string

const (
	ACLActionAllow ACLAction = "Allow"
	ACLActionDeny  ACLAction = "Deny"
)

func PossibleValuesForACLAction() []string {
	return []string{
		string(ACLActionAllow),
		string(ACLActionDeny),
	}
}

type EventListenerEndpointDiscriminator string

const (
	EventListenerEndpointDiscriminatorEventHub EventListenerEndpointDiscriminator = "EventHub"
)

func PossibleValuesForEventListenerEndpointDiscriminator() []string {
	return []string{
		string(EventListenerEndpointDiscriminatorEventHub),
	}
}

type EventListenerFilterDiscriminator string

const (
	EventListenerFilterDiscriminatorEventName EventListenerFilterDiscriminator = "EventName"
)

func PossibleValuesForEventListenerFilterDiscriminator() []string {
	return []string{
		string(EventListenerFilterDiscriminatorEventName),
	}
}

type KeyType string

const (
	KeyTypePrimary   KeyType = "Primary"
	KeyTypeSalt      KeyType = "Salt"
	KeyTypeSecondary KeyType = "Secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimary),
		string(KeyTypeSalt),
		string(KeyTypeSecondary),
	}
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkServiceConnectionStatus() []string {
	return []string{
		string(PrivateLinkServiceConnectionStatusApproved),
		string(PrivateLinkServiceConnectionStatusDisconnected),
		string(PrivateLinkServiceConnectionStatusPending),
		string(PrivateLinkServiceConnectionStatusRejected),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateRunning   ProvisioningState = "Running"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
		string(ProvisioningStateUpdating),
	}
}

type ScaleType string

const (
	ScaleTypeAutomatic ScaleType = "Automatic"
	ScaleTypeManual    ScaleType = "Manual"
	ScaleTypeNone      ScaleType = "None"
)

func PossibleValuesForScaleType() []string {
	return []string{
		string(ScaleTypeAutomatic),
		string(ScaleTypeManual),
		string(ScaleTypeNone),
	}
}

type SharedPrivateLinkResourceStatus string

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = "Approved"
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = "Pending"
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = "Rejected"
	SharedPrivateLinkResourceStatusTimeout      SharedPrivateLinkResourceStatus = "Timeout"
)

func PossibleValuesForSharedPrivateLinkResourceStatus() []string {
	return []string{
		string(SharedPrivateLinkResourceStatusApproved),
		string(SharedPrivateLinkResourceStatusDisconnected),
		string(SharedPrivateLinkResourceStatusPending),
		string(SharedPrivateLinkResourceStatusRejected),
		string(SharedPrivateLinkResourceStatusTimeout),
	}
}

type UpstreamAuthType string

const (
	UpstreamAuthTypeManagedIdentity UpstreamAuthType = "ManagedIdentity"
	UpstreamAuthTypeNone            UpstreamAuthType = "None"
)

func PossibleValuesForUpstreamAuthType() []string {
	return []string{
		string(UpstreamAuthTypeManagedIdentity),
		string(UpstreamAuthTypeNone),
	}
}

type WebPubSubRequestType string

const (
	WebPubSubRequestTypeClientConnection WebPubSubRequestType = "ClientConnection"
	WebPubSubRequestTypeRESTAPI          WebPubSubRequestType = "RESTAPI"
	WebPubSubRequestTypeServerConnection WebPubSubRequestType = "ServerConnection"
	WebPubSubRequestTypeTrace            WebPubSubRequestType = "Trace"
)

func PossibleValuesForWebPubSubRequestType() []string {
	return []string{
		string(WebPubSubRequestTypeClientConnection),
		string(WebPubSubRequestTypeRESTAPI),
		string(WebPubSubRequestTypeServerConnection),
		string(WebPubSubRequestTypeTrace),
	}
}

type WebPubSubSkuTier string

const (
	WebPubSubSkuTierBasic    WebPubSubSkuTier = "Basic"
	WebPubSubSkuTierFree     WebPubSubSkuTier = "Free"
	WebPubSubSkuTierPremium  WebPubSubSkuTier = "Premium"
	WebPubSubSkuTierStandard WebPubSubSkuTier = "Standard"
)

func PossibleValuesForWebPubSubSkuTier() []string {
	return []string{
		string(WebPubSubSkuTierBasic),
		string(WebPubSubSkuTierFree),
		string(WebPubSubSkuTierPremium),
		string(WebPubSubSkuTierStandard),
	}
}
