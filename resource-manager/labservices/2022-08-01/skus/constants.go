package skus

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

type RestrictionType string

const (
	RestrictionTypeLocation RestrictionType = "Location"
)

func PossibleValuesForRestrictionType() []string {
	return []string{
		string(RestrictionTypeLocation),
	}
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
