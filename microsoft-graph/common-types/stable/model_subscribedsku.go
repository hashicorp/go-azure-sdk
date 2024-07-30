package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscribedSku struct {
	AccountId        *string             `json:"accountId,omitempty"`
	AccountName      *string             `json:"accountName,omitempty"`
	AppliesTo        *string             `json:"appliesTo,omitempty"`
	CapabilityStatus *string             `json:"capabilityStatus,omitempty"`
	ConsumedUnits    *int64              `json:"consumedUnits,omitempty"`
	Id               *string             `json:"id,omitempty"`
	ODataType        *string             `json:"@odata.type,omitempty"`
	PrepaidUnits     *LicenseUnitsDetail `json:"prepaidUnits,omitempty"`
	ServicePlans     *[]ServicePlanInfo  `json:"servicePlans,omitempty"`
	SkuId            *string             `json:"skuId,omitempty"`
	SkuPartNumber    *string             `json:"skuPartNumber,omitempty"`
	SubscriptionIds  *[]string           `json:"subscriptionIds,omitempty"`
}
