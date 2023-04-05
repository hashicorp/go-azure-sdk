package skus

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
