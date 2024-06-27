package reservationorder

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationProperty struct {
	AppliedScopeProperties       *ReservationAppliedScopeProperties `json:"appliedScopeProperties,omitempty"`
	AppliedScopeType             *string                            `json:"appliedScopeType,omitempty"`
	AppliedScopes                *[]string                          `json:"appliedScopes,omitempty"`
	Archived                     *bool                              `json:"archived,omitempty"`
	BenefitStartTime             *string                            `json:"benefitStartTime,omitempty"`
	BillingPlan                  *ReservationBillingPlan            `json:"billingPlan,omitempty"`
	BillingScopeId               *string                            `json:"billingScopeId,omitempty"`
	Capabilities                 *string                            `json:"capabilities,omitempty"`
	DisplayName                  *string                            `json:"displayName,omitempty"`
	DisplayProvisioningState     *string                            `json:"displayProvisioningState,omitempty"`
	EffectiveDateTime            *string                            `json:"effectiveDateTime,omitempty"`
	ExpiryDate                   *string                            `json:"expiryDate,omitempty"`
	ExpiryDateTime               *string                            `json:"expiryDateTime,omitempty"`
	ExtendedStatusInfo           *ReservationExtendedStatusInfo     `json:"extendedStatusInfo,omitempty"`
	InstanceFlexibility          *InstanceFlexibility               `json:"instanceFlexibility,omitempty"`
	LastUpdatedDateTime          *string                            `json:"lastUpdatedDateTime,omitempty"`
	MergeProperties              *ReservationMergeProperties        `json:"mergeProperties,omitempty"`
	ProductCode                  *string                            `json:"productCode,omitempty"`
	ProvisioningState            *string                            `json:"provisioningState,omitempty"`
	ProvisioningSubState         *string                            `json:"provisioningSubState,omitempty"`
	PurchaseDate                 *string                            `json:"purchaseDate,omitempty"`
	PurchaseDateTime             *string                            `json:"purchaseDateTime,omitempty"`
	Quantity                     *float64                           `json:"quantity,omitempty"`
	Renew                        *bool                              `json:"renew,omitempty"`
	RenewDestination             *string                            `json:"renewDestination,omitempty"`
	RenewProperties              *RenewPropertiesResponse           `json:"renewProperties,omitempty"`
	RenewSource                  *string                            `json:"renewSource,omitempty"`
	ReservedResourceType         *string                            `json:"reservedResourceType,omitempty"`
	ReviewDateTime               *string                            `json:"reviewDateTime,omitempty"`
	SkuDescription               *string                            `json:"skuDescription,omitempty"`
	SplitProperties              *ReservationSplitProperties        `json:"splitProperties,omitempty"`
	SwapProperties               *ReservationSwapProperties         `json:"swapProperties,omitempty"`
	Term                         *string                            `json:"term,omitempty"`
	UserFriendlyAppliedScopeType *string                            `json:"userFriendlyAppliedScopeType,omitempty"`
	UserFriendlyRenewState       *string                            `json:"userFriendlyRenewState,omitempty"`
	Utilization                  *ReservationPropertyUtilization    `json:"utilization,omitempty"`
}

func (o *ReservationProperty) GetBenefitStartTimeAsTime() (*time.Time, error) {
	if o.BenefitStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BenefitStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationProperty) SetBenefitStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BenefitStartTime = &formatted
}

func (o *ReservationProperty) GetEffectiveDateTimeAsTime() (*time.Time, error) {
	if o.EffectiveDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EffectiveDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationProperty) SetEffectiveDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EffectiveDateTime = &formatted
}

func (o *ReservationProperty) GetExpiryDateTimeAsTime() (*time.Time, error) {
	if o.ExpiryDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationProperty) SetExpiryDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryDateTime = &formatted
}

func (o *ReservationProperty) GetLastUpdatedDateTimeAsTime() (*time.Time, error) {
	if o.LastUpdatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdatedDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationProperty) SetLastUpdatedDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUpdatedDateTime = &formatted
}

func (o *ReservationProperty) GetPurchaseDateTimeAsTime() (*time.Time, error) {
	if o.PurchaseDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PurchaseDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationProperty) SetPurchaseDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PurchaseDateTime = &formatted
}

func (o *ReservationProperty) GetReviewDateTimeAsTime() (*time.Time, error) {
	if o.ReviewDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReviewDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationProperty) SetReviewDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReviewDateTime = &formatted
}
