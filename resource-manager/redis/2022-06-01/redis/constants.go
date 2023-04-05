package redis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DayOfWeek string

const (
	DayOfWeekEveryday  DayOfWeek = "Everyday"
	DayOfWeekFriday    DayOfWeek = "Friday"
	DayOfWeekMonday    DayOfWeek = "Monday"
	DayOfWeekSaturday  DayOfWeek = "Saturday"
	DayOfWeekSunday    DayOfWeek = "Sunday"
	DayOfWeekThursday  DayOfWeek = "Thursday"
	DayOfWeekTuesday   DayOfWeek = "Tuesday"
	DayOfWeekWednesday DayOfWeek = "Wednesday"
	DayOfWeekWeekend   DayOfWeek = "Weekend"
)

func PossibleValuesForDayOfWeek() []string {
	return []string{
		string(DayOfWeekEveryday),
		string(DayOfWeekFriday),
		string(DayOfWeekMonday),
		string(DayOfWeekSaturday),
		string(DayOfWeekSunday),
		string(DayOfWeekThursday),
		string(DayOfWeekTuesday),
		string(DayOfWeekWednesday),
		string(DayOfWeekWeekend),
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

type ProvisioningState string

const (
	ProvisioningStateCreating               ProvisioningState = "Creating"
	ProvisioningStateDeleting               ProvisioningState = "Deleting"
	ProvisioningStateDisabled               ProvisioningState = "Disabled"
	ProvisioningStateFailed                 ProvisioningState = "Failed"
	ProvisioningStateLinking                ProvisioningState = "Linking"
	ProvisioningStateProvisioning           ProvisioningState = "Provisioning"
	ProvisioningStateRecoveringScaleFailure ProvisioningState = "RecoveringScaleFailure"
	ProvisioningStateScaling                ProvisioningState = "Scaling"
	ProvisioningStateSucceeded              ProvisioningState = "Succeeded"
	ProvisioningStateUnlinking              ProvisioningState = "Unlinking"
	ProvisioningStateUnprovisioning         ProvisioningState = "Unprovisioning"
	ProvisioningStateUpdating               ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateDisabled),
		string(ProvisioningStateFailed),
		string(ProvisioningStateLinking),
		string(ProvisioningStateProvisioning),
		string(ProvisioningStateRecoveringScaleFailure),
		string(ProvisioningStateScaling),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnlinking),
		string(ProvisioningStateUnprovisioning),
		string(ProvisioningStateUpdating),
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

type RebootType string

const (
	RebootTypeAllNodes      RebootType = "AllNodes"
	RebootTypePrimaryNode   RebootType = "PrimaryNode"
	RebootTypeSecondaryNode RebootType = "SecondaryNode"
)

func PossibleValuesForRebootType() []string {
	return []string{
		string(RebootTypeAllNodes),
		string(RebootTypePrimaryNode),
		string(RebootTypeSecondaryNode),
	}
}

type RedisKeyType string

const (
	RedisKeyTypePrimary   RedisKeyType = "Primary"
	RedisKeyTypeSecondary RedisKeyType = "Secondary"
)

func PossibleValuesForRedisKeyType() []string {
	return []string{
		string(RedisKeyTypePrimary),
		string(RedisKeyTypeSecondary),
	}
}

type ReplicationRole string

const (
	ReplicationRolePrimary   ReplicationRole = "Primary"
	ReplicationRoleSecondary ReplicationRole = "Secondary"
)

func PossibleValuesForReplicationRole() []string {
	return []string{
		string(ReplicationRolePrimary),
		string(ReplicationRoleSecondary),
	}
}

type SkuFamily string

const (
	SkuFamilyC SkuFamily = "C"
	SkuFamilyP SkuFamily = "P"
)

func PossibleValuesForSkuFamily() []string {
	return []string{
		string(SkuFamilyC),
		string(SkuFamilyP),
	}
}

type SkuName string

const (
	SkuNameBasic    SkuName = "Basic"
	SkuNamePremium  SkuName = "Premium"
	SkuNameStandard SkuName = "Standard"
)

func PossibleValuesForSkuName() []string {
	return []string{
		string(SkuNameBasic),
		string(SkuNamePremium),
		string(SkuNameStandard),
	}
}

type TlsVersion string

const (
	TlsVersionOnePointOne  TlsVersion = "1.1"
	TlsVersionOnePointTwo  TlsVersion = "1.2"
	TlsVersionOnePointZero TlsVersion = "1.0"
)

func PossibleValuesForTlsVersion() []string {
	return []string{
		string(TlsVersionOnePointOne),
		string(TlsVersionOnePointTwo),
		string(TlsVersionOnePointZero),
	}
}
