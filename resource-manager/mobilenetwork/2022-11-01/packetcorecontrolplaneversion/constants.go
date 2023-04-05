package packetcorecontrolplaneversion

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObsoleteVersion string

const (
	ObsoleteVersionNotObsolete ObsoleteVersion = "NotObsolete"
	ObsoleteVersionObsolete    ObsoleteVersion = "Obsolete"
)

func PossibleValuesForObsoleteVersion() []string {
	return []string{
		string(ObsoleteVersionNotObsolete),
		string(ObsoleteVersionObsolete),
	}
}

type PlatformType string

const (
	PlatformTypeAKSNegativeHCI                              PlatformType = "AKS-HCI"
	PlatformTypeThreePNegativeAZURENegativeSTACKNegativeHCI PlatformType = "3P-AZURE-STACK-HCI"
)

func PossibleValuesForPlatformType() []string {
	return []string{
		string(PlatformTypeAKSNegativeHCI),
		string(PlatformTypeThreePNegativeAZURENegativeSTACKNegativeHCI),
	}
}

type ProvisioningState string

const (
	ProvisioningStateAccepted  ProvisioningState = "Accepted"
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateDeleted   ProvisioningState = "Deleted"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUnknown   ProvisioningState = "Unknown"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
	}
}

type RecommendedVersion string

const (
	RecommendedVersionNotRecommended RecommendedVersion = "NotRecommended"
	RecommendedVersionRecommended    RecommendedVersion = "Recommended"
)

func PossibleValuesForRecommendedVersion() []string {
	return []string{
		string(RecommendedVersionNotRecommended),
		string(RecommendedVersionRecommended),
	}
}

type VersionState string

const (
	VersionStateActive           VersionState = "Active"
	VersionStateDeprecated       VersionState = "Deprecated"
	VersionStatePreview          VersionState = "Preview"
	VersionStateUnknown          VersionState = "Unknown"
	VersionStateValidating       VersionState = "Validating"
	VersionStateValidationFailed VersionState = "ValidationFailed"
)

func PossibleValuesForVersionState() []string {
	return []string{
		string(VersionStateActive),
		string(VersionStateDeprecated),
		string(VersionStatePreview),
		string(VersionStateUnknown),
		string(VersionStateValidating),
		string(VersionStateValidationFailed),
	}
}
