package apimanagementserviceskus

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
