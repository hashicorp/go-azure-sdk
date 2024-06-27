package reservationorder

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationOrderProperty struct {
	BenefitStartTime   *string                                 `json:"benefitStartTime,omitempty"`
	BillingAccountId   *string                                 `json:"billingAccountId,omitempty"`
	BillingPlan        *ReservationBillingPlan                 `json:"billingPlan,omitempty"`
	BillingProfileId   *string                                 `json:"billingProfileId,omitempty"`
	CreatedDateTime    *string                                 `json:"createdDateTime,omitempty"`
	CustomerId         *string                                 `json:"customerId,omitempty"`
	DisplayName        *string                                 `json:"displayName,omitempty"`
	EnrollmentId       *string                                 `json:"enrollmentId,omitempty"`
	ExpiryDate         *string                                 `json:"expiryDate,omitempty"`
	ExpiryDateTime     *string                                 `json:"expiryDateTime,omitempty"`
	ExtendedStatusInfo *ReservationExtendedStatusInfo          `json:"extendedStatusInfo,omitempty"`
	OriginalQuantity   *int64                                  `json:"originalQuantity,omitempty"`
	PlanInformation    *ReservationOrderBillingPlanInformation `json:"planInformation,omitempty"`
	ProductCode        *string                                 `json:"productCode,omitempty"`
	ProvisioningState  *string                                 `json:"provisioningState,omitempty"`
	RequestDateTime    *string                                 `json:"requestDateTime,omitempty"`
	Reservations       *[]Reservation                          `json:"reservations,omitempty"`
	ReviewDateTime     *string                                 `json:"reviewDateTime,omitempty"`
	Term               *string                                 `json:"term,omitempty"`
}

func (o *ReservationOrderProperty) GetBenefitStartTimeAsTime() (*time.Time, error) {
	if o.BenefitStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BenefitStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationOrderProperty) SetBenefitStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BenefitStartTime = &formatted
}

func (o *ReservationOrderProperty) GetCreatedDateTimeAsTime() (*time.Time, error) {
	if o.CreatedDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreatedDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationOrderProperty) SetCreatedDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreatedDateTime = &formatted
}

func (o *ReservationOrderProperty) GetExpiryDateTimeAsTime() (*time.Time, error) {
	if o.ExpiryDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationOrderProperty) SetExpiryDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryDateTime = &formatted
}

func (o *ReservationOrderProperty) GetRequestDateTimeAsTime() (*time.Time, error) {
	if o.RequestDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RequestDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationOrderProperty) SetRequestDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RequestDateTime = &formatted
}

func (o *ReservationOrderProperty) GetReviewDateTimeAsTime() (*time.Time, error) {
	if o.ReviewDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReviewDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ReservationOrderProperty) SetReviewDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReviewDateTime = &formatted
}
