package proxy

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NameUnavailableReason string

const (
	NameUnavailableReasonAlreadyExists NameUnavailableReason = "AlreadyExists"
	NameUnavailableReasonInvalid       NameUnavailableReason = "Invalid"
)

func PossibleValuesForNameUnavailableReason() []string {
	return []string{
		string(NameUnavailableReasonAlreadyExists),
		string(NameUnavailableReasonInvalid),
	}
}

func parseNameUnavailableReason(input string) (*NameUnavailableReason, error) {
	vals := map[string]NameUnavailableReason{
		"alreadyexists": NameUnavailableReasonAlreadyExists,
		"invalid":       NameUnavailableReasonInvalid,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NameUnavailableReason(input)
	return &out, nil
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

func parseSkuTier(input string) (*SkuTier, error) {
	vals := map[string]SkuTier{
		"basic":    SkuTierBasic,
		"free":     SkuTierFree,
		"premium":  SkuTierPremium,
		"standard": SkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuTier(input)
	return &out, nil
}
