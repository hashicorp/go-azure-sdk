package skus

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabServicesSkuTier string

const (
	LabServicesSkuTierPremium  LabServicesSkuTier = "Premium"
	LabServicesSkuTierStandard LabServicesSkuTier = "Standard"
)

func PossibleValuesForLabServicesSkuTier() []string {
	return []string{
		string(LabServicesSkuTierPremium),
		string(LabServicesSkuTierStandard),
	}
}

func parseLabServicesSkuTier(input string) (*LabServicesSkuTier, error) {
	vals := map[string]LabServicesSkuTier{
		"premium":  LabServicesSkuTierPremium,
		"standard": LabServicesSkuTierStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LabServicesSkuTier(input)
	return &out, nil
}

type RestrictionReasonCode string

const (
	RestrictionReasonCodeNotAvailableForSubscription RestrictionReasonCode = "NotAvailableForSubscription"
	RestrictionReasonCodeQuotaId                     RestrictionReasonCode = "QuotaId"
)

func PossibleValuesForRestrictionReasonCode() []string {
	return []string{
		string(RestrictionReasonCodeNotAvailableForSubscription),
		string(RestrictionReasonCodeQuotaId),
	}
}

func parseRestrictionReasonCode(input string) (*RestrictionReasonCode, error) {
	vals := map[string]RestrictionReasonCode{
		"notavailableforsubscription": RestrictionReasonCodeNotAvailableForSubscription,
		"quotaid":                     RestrictionReasonCodeQuotaId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictionReasonCode(input)
	return &out, nil
}

type RestrictionType string

const (
	RestrictionTypeLocation RestrictionType = "Location"
)

func PossibleValuesForRestrictionType() []string {
	return []string{
		string(RestrictionTypeLocation),
	}
}

func parseRestrictionType(input string) (*RestrictionType, error) {
	vals := map[string]RestrictionType{
		"location": RestrictionTypeLocation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictionType(input)
	return &out, nil
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

func parseScaleType(input string) (*ScaleType, error) {
	vals := map[string]ScaleType{
		"automatic": ScaleTypeAutomatic,
		"manual":    ScaleTypeManual,
		"none":      ScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ScaleType(input)
	return &out, nil
}
