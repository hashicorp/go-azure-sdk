package hostpool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DayOfWeek string

const (
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeekFriday),
		string(DayOfWeekMonday),
		string(DayOfWeekSaturday),
		string(DayOfWeekSunday),
		string(DayOfWeekThursday),
		string(DayOfWeekTuesday),
		string(DayOfWeekWednesday),
	}
}

type HostPoolType string

const (
	HostPoolTypeBYODesktop HostPoolType = "BYODesktop"
	HostPoolTypePersonal   HostPoolType = "Personal"
	HostPoolTypePooled     HostPoolType = "Pooled"
)

func PossibleValuesForHostPoolType() []string {
	return []string{
		string(HostPoolTypeBYODesktop),
		string(HostPoolTypePersonal),
		string(HostPoolTypePooled),
	}
}

type HostpoolPublicNetworkAccess string

const (
	HostpoolPublicNetworkAccessDisabled                   HostpoolPublicNetworkAccess = "Disabled"
	HostpoolPublicNetworkAccessEnabled                    HostpoolPublicNetworkAccess = "Enabled"
	HostpoolPublicNetworkAccessEnabledForClientsOnly      HostpoolPublicNetworkAccess = "EnabledForClientsOnly"
	HostpoolPublicNetworkAccessEnabledForSessionHostsOnly HostpoolPublicNetworkAccess = "EnabledForSessionHostsOnly"
)

func PossibleValuesForHostpoolPublicNetworkAccess() []string {
	return []string{
		string(HostpoolPublicNetworkAccessDisabled),
		string(HostpoolPublicNetworkAccessEnabled),
		string(HostpoolPublicNetworkAccessEnabledForClientsOnly),
		string(HostpoolPublicNetworkAccessEnabledForSessionHostsOnly),
	}
}

type LoadBalancerType string

const (
	LoadBalancerTypeBreadthFirst LoadBalancerType = "BreadthFirst"
	LoadBalancerTypeDepthFirst   LoadBalancerType = "DepthFirst"
	LoadBalancerTypePersistent   LoadBalancerType = "Persistent"
)

func PossibleValuesForLoadBalancerType() []string {
	return []string{
		string(LoadBalancerTypeBreadthFirst),
		string(LoadBalancerTypeDepthFirst),
		string(LoadBalancerTypePersistent),
	}
}

type Operation string

const (
	OperationComplete Operation = "Complete"
	OperationHide     Operation = "Hide"
	OperationRevoke   Operation = "Revoke"
	OperationStart    Operation = "Start"
	OperationUnhide   Operation = "Unhide"
)

func PossibleValuesForOperation() []string {
	return []string{
		string(OperationComplete),
		string(OperationHide),
		string(OperationRevoke),
		string(OperationStart),
		string(OperationUnhide),
	}
}

type PersonalDesktopAssignmentType string

const (
	PersonalDesktopAssignmentTypeAutomatic PersonalDesktopAssignmentType = "Automatic"
	PersonalDesktopAssignmentTypeDirect    PersonalDesktopAssignmentType = "Direct"
)

func PossibleValuesForPersonalDesktopAssignmentType() []string {
	return []string{
		string(PersonalDesktopAssignmentTypeAutomatic),
		string(PersonalDesktopAssignmentTypeDirect),
	}
}

type PreferredAppGroupType string

const (
	PreferredAppGroupTypeDesktop          PreferredAppGroupType = "Desktop"
	PreferredAppGroupTypeNone             PreferredAppGroupType = "None"
	PreferredAppGroupTypeRailApplications PreferredAppGroupType = "RailApplications"
)

func PossibleValuesForPreferredAppGroupType() []string {
	return []string{
		string(PreferredAppGroupTypeDesktop),
		string(PreferredAppGroupTypeNone),
		string(PreferredAppGroupTypeRailApplications),
	}
}

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating  PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleting  PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateFailed    PrivateEndpointConnectionProvisioningState = "Failed"
	PrivateEndpointConnectionProvisioningStateSucceeded PrivateEndpointConnectionProvisioningState = "Succeeded"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateFailed),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
	}
}

type PrivateEndpointServiceConnectionStatus string

const (
	PrivateEndpointServiceConnectionStatusApproved PrivateEndpointServiceConnectionStatus = "Approved"
	PrivateEndpointServiceConnectionStatusPending  PrivateEndpointServiceConnectionStatus = "Pending"
	PrivateEndpointServiceConnectionStatusRejected PrivateEndpointServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateEndpointServiceConnectionStatus() []string {
	return []string{
		string(PrivateEndpointServiceConnectionStatusApproved),
		string(PrivateEndpointServiceConnectionStatusPending),
		string(PrivateEndpointServiceConnectionStatusRejected),
	}
}

type RegistrationTokenOperation string

const (
	RegistrationTokenOperationDelete RegistrationTokenOperation = "Delete"
	RegistrationTokenOperationNone   RegistrationTokenOperation = "None"
	RegistrationTokenOperationUpdate RegistrationTokenOperation = "Update"
)

func PossibleValuesForRegistrationTokenOperation() []string {
	return []string{
		string(RegistrationTokenOperationDelete),
		string(RegistrationTokenOperationNone),
		string(RegistrationTokenOperationUpdate),
	}
}

type SSOSecretType string

const (
	SSOSecretTypeCertificate           SSOSecretType = "Certificate"
	SSOSecretTypeCertificateInKeyVault SSOSecretType = "CertificateInKeyVault"
	SSOSecretTypeSharedKey             SSOSecretType = "SharedKey"
	SSOSecretTypeSharedKeyInKeyVault   SSOSecretType = "SharedKeyInKeyVault"
)

func PossibleValuesForSSOSecretType() []string {
	return []string{
		string(SSOSecretTypeCertificate),
		string(SSOSecretTypeCertificateInKeyVault),
		string(SSOSecretTypeSharedKey),
		string(SSOSecretTypeSharedKeyInKeyVault),
	}
}

type SessionHostComponentUpdateType string

const (
	SessionHostComponentUpdateTypeDefault   SessionHostComponentUpdateType = "Default"
	SessionHostComponentUpdateTypeScheduled SessionHostComponentUpdateType = "Scheduled"
)

func PossibleValuesForSessionHostComponentUpdateType() []string {
	return []string{
		string(SessionHostComponentUpdateTypeDefault),
		string(SessionHostComponentUpdateTypeScheduled),
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
