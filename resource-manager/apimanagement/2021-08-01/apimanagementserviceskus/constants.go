package apimanagementserviceskus

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceSkuCapacityScaleType string

const (
	ResourceSkuCapacityScaleTypeAutomatic ResourceSkuCapacityScaleType = "automatic"
	ResourceSkuCapacityScaleTypeManual    ResourceSkuCapacityScaleType = "manual"
	ResourceSkuCapacityScaleTypeNone      ResourceSkuCapacityScaleType = "none"
)

func PossibleValuesForResourceSkuCapacityScaleType() []string {
	return []string{
		string(ResourceSkuCapacityScaleTypeAutomatic),
		string(ResourceSkuCapacityScaleTypeManual),
		string(ResourceSkuCapacityScaleTypeNone),
	}
}

func parseResourceSkuCapacityScaleType(input string) (*ResourceSkuCapacityScaleType, error) {
	vals := map[string]ResourceSkuCapacityScaleType{
		"automatic": ResourceSkuCapacityScaleTypeAutomatic,
		"manual":    ResourceSkuCapacityScaleTypeManual,
		"none":      ResourceSkuCapacityScaleTypeNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResourceSkuCapacityScaleType(input)
	return &out, nil
}

type SkuType string

const (
	SkuTypeBasic       SkuType = "Basic"
	SkuTypeConsumption SkuType = "Consumption"
	SkuTypeDeveloper   SkuType = "Developer"
	SkuTypeIsolated    SkuType = "Isolated"
	SkuTypePremium     SkuType = "Premium"
	SkuTypeStandard    SkuType = "Standard"
)

func PossibleValuesForSkuType() []string {
	return []string{
		string(SkuTypeBasic),
		string(SkuTypeConsumption),
		string(SkuTypeDeveloper),
		string(SkuTypeIsolated),
		string(SkuTypePremium),
		string(SkuTypeStandard),
	}
}

func parseSkuType(input string) (*SkuType, error) {
	vals := map[string]SkuType{
		"basic":       SkuTypeBasic,
		"consumption": SkuTypeConsumption,
		"developer":   SkuTypeDeveloper,
		"isolated":    SkuTypeIsolated,
		"premium":     SkuTypePremium,
		"standard":    SkuTypeStandard,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SkuType(input)
	return &out, nil
}
