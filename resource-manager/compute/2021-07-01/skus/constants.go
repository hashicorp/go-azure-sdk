package skus

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExtendedLocationType string

const (
	ExtendedLocationTypeEdgeZone ExtendedLocationType = "EdgeZone"
)

func PossibleValuesForExtendedLocationType() []string {
	return []string{
		string(ExtendedLocationTypeEdgeZone),
	}
}

type ResourceSkuCapacityScaleType string

const (
	ResourceSkuCapacityScaleTypeAutomatic ResourceSkuCapacityScaleType = "Automatic"
	ResourceSkuCapacityScaleTypeManual    ResourceSkuCapacityScaleType = "Manual"
	ResourceSkuCapacityScaleTypeNone      ResourceSkuCapacityScaleType = "None"
)

func PossibleValuesForResourceSkuCapacityScaleType() []string {
	return []string{
		string(ResourceSkuCapacityScaleTypeAutomatic),
		string(ResourceSkuCapacityScaleTypeManual),
		string(ResourceSkuCapacityScaleTypeNone),
	}
}

type ResourceSkuRestrictionsReasonCode string

const (
	ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription ResourceSkuRestrictionsReasonCode = "NotAvailableForSubscription"
	ResourceSkuRestrictionsReasonCodeQuotaId                     ResourceSkuRestrictionsReasonCode = "QuotaId"
)

func PossibleValuesForResourceSkuRestrictionsReasonCode() []string {
	return []string{
		string(ResourceSkuRestrictionsReasonCodeNotAvailableForSubscription),
		string(ResourceSkuRestrictionsReasonCodeQuotaId),
	}
}

type ResourceSkuRestrictionsType string

const (
	ResourceSkuRestrictionsTypeLocation ResourceSkuRestrictionsType = "Location"
	ResourceSkuRestrictionsTypeZone     ResourceSkuRestrictionsType = "Zone"
)

func PossibleValuesForResourceSkuRestrictionsType() []string {
	return []string{
		string(ResourceSkuRestrictionsTypeLocation),
		string(ResourceSkuRestrictionsTypeZone),
	}
}
