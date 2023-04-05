package fileshares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnabledProtocols string

const (
	EnabledProtocolsNFS EnabledProtocols = "NFS"
	EnabledProtocolsSMB EnabledProtocols = "SMB"
)

func PossibleValuesForEnabledProtocols() []string {
	return []string{
		string(EnabledProtocolsNFS),
		string(EnabledProtocolsSMB),
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

type LeaseShareAction string

const (
	LeaseShareActionAcquire LeaseShareAction = "Acquire"
	LeaseShareActionBreak   LeaseShareAction = "Break"
	LeaseShareActionChange  LeaseShareAction = "Change"
	LeaseShareActionRelease LeaseShareAction = "Release"
	LeaseShareActionRenew   LeaseShareAction = "Renew"
)

func PossibleValuesForLeaseShareAction() []string {
	return []string{
		string(LeaseShareActionAcquire),
		string(LeaseShareActionBreak),
		string(LeaseShareActionChange),
		string(LeaseShareActionRelease),
		string(LeaseShareActionRenew),
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

type RootSquashType string

const (
	RootSquashTypeAllSquash    RootSquashType = "AllSquash"
	RootSquashTypeNoRootSquash RootSquashType = "NoRootSquash"
	RootSquashTypeRootSquash   RootSquashType = "RootSquash"
)

func PossibleValuesForRootSquashType() []string {
	return []string{
		string(RootSquashTypeAllSquash),
		string(RootSquashTypeNoRootSquash),
		string(RootSquashTypeRootSquash),
	}
}

type ShareAccessTier string

const (
	ShareAccessTierCool                 ShareAccessTier = "Cool"
	ShareAccessTierHot                  ShareAccessTier = "Hot"
	ShareAccessTierPremium              ShareAccessTier = "Premium"
	ShareAccessTierTransactionOptimized ShareAccessTier = "TransactionOptimized"
)

func PossibleValuesForShareAccessTier() []string {
	return []string{
		string(ShareAccessTierCool),
		string(ShareAccessTierHot),
		string(ShareAccessTierPremium),
		string(ShareAccessTierTransactionOptimized),
	}
}
