package skus

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApiManagementSkuCapacityScaleType string

const (
	ApiManagementSkuCapacityScaleTypeAutomatic ApiManagementSkuCapacityScaleType = "Automatic"
	ApiManagementSkuCapacityScaleTypeManual    ApiManagementSkuCapacityScaleType = "Manual"
	ApiManagementSkuCapacityScaleTypeNone      ApiManagementSkuCapacityScaleType = "None"
)

func PossibleValuesForApiManagementSkuCapacityScaleType() []string {
	return []string{
		string(ApiManagementSkuCapacityScaleTypeAutomatic),
		string(ApiManagementSkuCapacityScaleTypeManual),
		string(ApiManagementSkuCapacityScaleTypeNone),
	}
}

func parseApiManagementSkuCapacityScaleType(input string) (*ApiManagementSkuCapacityScaleType, error) {
	vals := map[string]ApiManagementSkuCapacityScaleType{
		"automatic": ApiManagementSkuCapacityScaleTypeAutomatic,
		"manual":    ApiManagementSkuCapacityScaleTypeManual,
		"none":      ApiManagementSkuCapacityScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiManagementSkuCapacityScaleType(input)
	return &out, nil
}

type ApiManagementSkuRestrictionsReasonCode string

const (
	ApiManagementSkuRestrictionsReasonCodeNotAvailableForSubscription ApiManagementSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	ApiManagementSkuRestrictionsReasonCodeQuotaId                     ApiManagementSkuRestrictionsReasonCode = "QuotaId"
)

func PossibleValuesForApiManagementSkuRestrictionsReasonCode() []string {
	return []string{
		string(ApiManagementSkuRestrictionsReasonCodeNotAvailableForSubscription),
		string(ApiManagementSkuRestrictionsReasonCodeQuotaId),
	}
}

func parseApiManagementSkuRestrictionsReasonCode(input string) (*ApiManagementSkuRestrictionsReasonCode, error) {
	vals := map[string]ApiManagementSkuRestrictionsReasonCode{
		"notavailableforsubscription": ApiManagementSkuRestrictionsReasonCodeNotAvailableForSubscription,
		"quotaid":                     ApiManagementSkuRestrictionsReasonCodeQuotaId,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiManagementSkuRestrictionsReasonCode(input)
	return &out, nil
}

type ApiManagementSkuRestrictionsType string

const (
	ApiManagementSkuRestrictionsTypeLocation ApiManagementSkuRestrictionsType = "Location"
	ApiManagementSkuRestrictionsTypeZone     ApiManagementSkuRestrictionsType = "Zone"
)

func PossibleValuesForApiManagementSkuRestrictionsType() []string {
	return []string{
		string(ApiManagementSkuRestrictionsTypeLocation),
		string(ApiManagementSkuRestrictionsTypeZone),
	}
}

func parseApiManagementSkuRestrictionsType(input string) (*ApiManagementSkuRestrictionsType, error) {
	vals := map[string]ApiManagementSkuRestrictionsType{
		"location": ApiManagementSkuRestrictionsTypeLocation,
		"zone":     ApiManagementSkuRestrictionsTypeZone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ApiManagementSkuRestrictionsType(input)
	return &out, nil
}
