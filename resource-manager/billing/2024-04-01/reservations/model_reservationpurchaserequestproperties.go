package reservations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationPurchaseRequestProperties struct {
	AppliedScopeProperties     *ReservationAppliedScopeProperties                              `json:"appliedScopeProperties,omitempty"`
	AppliedScopeType           *AppliedScopeType                                               `json:"appliedScopeType,omitempty"`
	AppliedScopes              *[]string                                                       `json:"appliedScopes,omitempty"`
	BillingPlan                *ReservationBillingPlan                                         `json:"billingPlan,omitempty"`
	BillingScopeId             *string                                                         `json:"billingScopeId,omitempty"`
	DisplayName                *string                                                         `json:"displayName,omitempty"`
	InstanceFlexibility        *InstanceFlexibility                                            `json:"instanceFlexibility,omitempty"`
	Quantity                   *int64                                                          `json:"quantity,omitempty"`
	Renew                      *bool                                                           `json:"renew,omitempty"`
	ReservedResourceProperties *ReservationPurchaseRequestPropertiesReservedResourceProperties `json:"reservedResourceProperties,omitempty"`
	ReservedResourceType       *string                                                         `json:"reservedResourceType,omitempty"`
	ReviewDateTime             *string                                                         `json:"reviewDateTime,omitempty"`
	Term                       *string                                                         `json:"term,omitempty"`
}

func (o *ReservationPurchaseRequestProperties) GetReviewDateTimeAsTime() (*time.Time, error) {
	if o.ReviewDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReviewDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationPurchaseRequestProperties) SetReviewDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReviewDateTime = &formatted
}
