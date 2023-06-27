package packetcorecontrolplaneversion

import "strings"

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

func parseObsoleteVersion(input string) (*ObsoleteVersion, error) {
	vals := map[string]ObsoleteVersion{
		"notobsolete": ObsoleteVersionNotObsolete,
		"obsolete":    ObsoleteVersionObsolete,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ObsoleteVersion(input)
	return &out, nil
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

func parsePlatformType(input string) (*PlatformType, error) {
	vals := map[string]PlatformType{
		"aks-hci":            PlatformTypeAKSNegativeHCI,
		"3p-azure-stack-hci": PlatformTypeThreePNegativeAZURENegativeSTACKNegativeHCI,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlatformType(input)
	return &out, nil
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

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":  ProvisioningStateAccepted,
		"canceled":  ProvisioningStateCanceled,
		"deleted":   ProvisioningStateDeleted,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"succeeded": ProvisioningStateSucceeded,
		"unknown":   ProvisioningStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
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

func parseRecommendedVersion(input string) (*RecommendedVersion, error) {
	vals := map[string]RecommendedVersion{
		"notrecommended": RecommendedVersionNotRecommended,
		"recommended":    RecommendedVersionRecommended,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecommendedVersion(input)
	return &out, nil
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

func parseVersionState(input string) (*VersionState, error) {
	vals := map[string]VersionState{
		"active":           VersionStateActive,
		"deprecated":       VersionStateDeprecated,
		"preview":          VersionStatePreview,
		"unknown":          VersionStateUnknown,
		"validating":       VersionStateValidating,
		"validationfailed": VersionStateValidationFailed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VersionState(input)
	return &out, nil
}
