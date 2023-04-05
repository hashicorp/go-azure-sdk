package iotdpsresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessRightsDescription string

const (
	AccessRightsDescriptionDeviceConnect           AccessRightsDescription = "DeviceConnect"
	AccessRightsDescriptionEnrollmentRead          AccessRightsDescription = "EnrollmentRead"
	AccessRightsDescriptionEnrollmentWrite         AccessRightsDescription = "EnrollmentWrite"
	AccessRightsDescriptionRegistrationStatusRead  AccessRightsDescription = "RegistrationStatusRead"
	AccessRightsDescriptionRegistrationStatusWrite AccessRightsDescription = "RegistrationStatusWrite"
	AccessRightsDescriptionServiceConfig           AccessRightsDescription = "ServiceConfig"
)

func PossibleValuesForAccessRightsDescription() []string {
	return []string{
		string(AccessRightsDescriptionDeviceConnect),
		string(AccessRightsDescriptionEnrollmentRead),
		string(AccessRightsDescriptionEnrollmentWrite),
		string(AccessRightsDescriptionRegistrationStatusRead),
		string(AccessRightsDescriptionRegistrationStatusWrite),
		string(AccessRightsDescriptionServiceConfig),
	}
}

type AllocationPolicy string

const (
	AllocationPolicyGeoLatency AllocationPolicy = "GeoLatency"
	AllocationPolicyHashed     AllocationPolicy = "Hashed"
	AllocationPolicyStatic     AllocationPolicy = "Static"
)

func PossibleValuesForAllocationPolicy() []string {
	return []string{
		string(AllocationPolicyGeoLatency),
		string(AllocationPolicyHashed),
		string(AllocationPolicyStatic),
	}
}

type IPFilterActionType string

const (
	IPFilterActionTypeAccept IPFilterActionType = "Accept"
	IPFilterActionTypeReject IPFilterActionType = "Reject"
)

func PossibleValuesForIPFilterActionType() []string {
	return []string{
		string(IPFilterActionTypeAccept),
		string(IPFilterActionTypeReject),
	}
}

type IPFilterTargetType string

const (
	IPFilterTargetTypeAll        IPFilterTargetType = "all"
	IPFilterTargetTypeDeviceApi  IPFilterTargetType = "deviceApi"
	IPFilterTargetTypeServiceApi IPFilterTargetType = "serviceApi"
)

func PossibleValuesForIPFilterTargetType() []string {
	return []string{
		string(IPFilterTargetTypeAll),
		string(IPFilterTargetTypeDeviceApi),
		string(IPFilterTargetTypeServiceApi),
	}
}

type IotDpsSku string

const (
	IotDpsSkuSOne IotDpsSku = "S1"
)

func PossibleValuesForIotDpsSku() []string {
	return []string{
		string(IotDpsSkuSOne),
	}
}

type NameUnavailabilityReason string

const (
	NameUnavailabilityReasonAlreadyExists NameUnavailabilityReason = "AlreadyExists"
	NameUnavailabilityReasonInvalid       NameUnavailabilityReason = "Invalid"
)

func PossibleValuesForNameUnavailabilityReason() []string {
	return []string{
		string(NameUnavailabilityReasonAlreadyExists),
		string(NameUnavailabilityReasonInvalid),
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

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled  PublicNetworkAccess = "Enabled"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
	}
}

type State string

const (
	StateActivating       State = "Activating"
	StateActivationFailed State = "ActivationFailed"
	StateActive           State = "Active"
	StateDeleted          State = "Deleted"
	StateDeleting         State = "Deleting"
	StateDeletionFailed   State = "DeletionFailed"
	StateFailingOver      State = "FailingOver"
	StateFailoverFailed   State = "FailoverFailed"
	StateResuming         State = "Resuming"
	StateSuspended        State = "Suspended"
	StateSuspending       State = "Suspending"
	StateTransitioning    State = "Transitioning"
)

func PossibleValuesForState() []string {
	return []string{
		string(StateActivating),
		string(StateActivationFailed),
		string(StateActive),
		string(StateDeleted),
		string(StateDeleting),
		string(StateDeletionFailed),
		string(StateFailingOver),
		string(StateFailoverFailed),
		string(StateResuming),
		string(StateSuspended),
		string(StateSuspending),
		string(StateTransitioning),
	}
}
