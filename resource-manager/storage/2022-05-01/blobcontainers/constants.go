package blobcontainers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImmutabilityPolicyState string

const (
	ImmutabilityPolicyStateLocked   ImmutabilityPolicyState = "Locked"
	ImmutabilityPolicyStateUnlocked ImmutabilityPolicyState = "Unlocked"
)

func PossibleValuesForImmutabilityPolicyState() []string {
	return []string{
		string(ImmutabilityPolicyStateLocked),
		string(ImmutabilityPolicyStateUnlocked),
	}
}

type ImmutabilityPolicyUpdateType string

const (
	ImmutabilityPolicyUpdateTypeExtend ImmutabilityPolicyUpdateType = "extend"
	ImmutabilityPolicyUpdateTypeLock   ImmutabilityPolicyUpdateType = "lock"
	ImmutabilityPolicyUpdateTypePut    ImmutabilityPolicyUpdateType = "put"
)

func PossibleValuesForImmutabilityPolicyUpdateType() []string {
	return []string{
		string(ImmutabilityPolicyUpdateTypeExtend),
		string(ImmutabilityPolicyUpdateTypeLock),
		string(ImmutabilityPolicyUpdateTypePut),
	}
}

type LeaseContainerRequestAction string

const (
	LeaseContainerRequestActionAcquire LeaseContainerRequestAction = "Acquire"
	LeaseContainerRequestActionBreak   LeaseContainerRequestAction = "Break"
	LeaseContainerRequestActionChange  LeaseContainerRequestAction = "Change"
	LeaseContainerRequestActionRelease LeaseContainerRequestAction = "Release"
	LeaseContainerRequestActionRenew   LeaseContainerRequestAction = "Renew"
)

func PossibleValuesForLeaseContainerRequestAction() []string {
	return []string{
		string(LeaseContainerRequestActionAcquire),
		string(LeaseContainerRequestActionBreak),
		string(LeaseContainerRequestActionChange),
		string(LeaseContainerRequestActionRelease),
		string(LeaseContainerRequestActionRenew),
	}
}

type LeaseDuration string

const (
	LeaseDurationFixed    LeaseDuration = "Fixed"
	LeaseDurationInfinite LeaseDuration = "Infinite"
)

func PossibleValuesForLeaseDuration() []string {
	return []string{
		string(LeaseDurationFixed),
		string(LeaseDurationInfinite),
	}
}

type LeaseState string

const (
	LeaseStateAvailable LeaseState = "Available"
	LeaseStateBreaking  LeaseState = "Breaking"
	LeaseStateBroken    LeaseState = "Broken"
	LeaseStateExpired   LeaseState = "Expired"
	LeaseStateLeased    LeaseState = "Leased"
)

func PossibleValuesForLeaseState() []string {
	return []string{
		string(LeaseStateAvailable),
		string(LeaseStateBreaking),
		string(LeaseStateBroken),
		string(LeaseStateExpired),
		string(LeaseStateLeased),
	}
}

type LeaseStatus string

const (
	LeaseStatusLocked   LeaseStatus = "Locked"
	LeaseStatusUnlocked LeaseStatus = "Unlocked"
)

func PossibleValuesForLeaseStatus() []string {
	return []string{
		string(LeaseStatusLocked),
		string(LeaseStatusUnlocked),
	}
}

type ListContainersInclude string

const (
	ListContainersIncludeDeleted ListContainersInclude = "deleted"
)

func PossibleValuesForListContainersInclude() []string {
	return []string{
		string(ListContainersIncludeDeleted),
	}
}

type MigrationState string

const (
	MigrationStateCompleted  MigrationState = "Completed"
	MigrationStateInProgress MigrationState = "InProgress"
)

func PossibleValuesForMigrationState() []string {
	return []string{
		string(MigrationStateCompleted),
		string(MigrationStateInProgress),
	}
}

type PublicAccess string

const (
	PublicAccessBlob      PublicAccess = "Blob"
	PublicAccessContainer PublicAccess = "Container"
	PublicAccessNone      PublicAccess = "None"
)

func PossibleValuesForPublicAccess() []string {
	return []string{
		string(PublicAccessBlob),
		string(PublicAccessContainer),
		string(PublicAccessNone),
	}
}
