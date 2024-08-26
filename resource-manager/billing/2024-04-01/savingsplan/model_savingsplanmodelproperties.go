package savingsplan

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanModelProperties struct {
	AppliedScopeProperties       *AppliedScopeProperties `json:"appliedScopeProperties,omitempty"`
	AppliedScopeType             *AppliedScopeType       `json:"appliedScopeType,omitempty"`
	BenefitStartTime             *string                 `json:"benefitStartTime,omitempty"`
	BillingAccountId             *string                 `json:"billingAccountId,omitempty"`
	BillingPlan                  *BillingPlan            `json:"billingPlan,omitempty"`
	BillingProfileId             *string                 `json:"billingProfileId,omitempty"`
	BillingScopeId               *string                 `json:"billingScopeId,omitempty"`
	Commitment                   *Commitment             `json:"commitment,omitempty"`
	CustomerId                   *string                 `json:"customerId,omitempty"`
	DisplayName                  *string                 `json:"displayName,omitempty"`
	DisplayProvisioningState     *string                 `json:"displayProvisioningState,omitempty"`
	EffectiveDateTime            *string                 `json:"effectiveDateTime,omitempty"`
	ExpiryDateTime               *string                 `json:"expiryDateTime,omitempty"`
	ExtendedStatusInfo           *ExtendedStatusInfo     `json:"extendedStatusInfo,omitempty"`
	ProductCode                  *string                 `json:"productCode,omitempty"`
	ProvisioningState            *ProvisioningState      `json:"provisioningState,omitempty"`
	PurchaseDateTime             *string                 `json:"purchaseDateTime,omitempty"`
	Renew                        *bool                   `json:"renew,omitempty"`
	RenewDestination             *string                 `json:"renewDestination,omitempty"`
	RenewProperties              *RenewProperties        `json:"renewProperties,omitempty"`
	RenewSource                  *string                 `json:"renewSource,omitempty"`
	Term                         *SavingsPlanTerm        `json:"term,omitempty"`
	UserFriendlyAppliedScopeType *string                 `json:"userFriendlyAppliedScopeType,omitempty"`
	Utilization                  *Utilization            `json:"utilization,omitempty"`
}

func (o *SavingsPlanModelProperties) GetBenefitStartTimeAsTime() (*time.Time, error) {
	if o.BenefitStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BenefitStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SavingsPlanModelProperties) SetBenefitStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BenefitStartTime = &formatted
}

func (o *SavingsPlanModelProperties) GetEffectiveDateTimeAsTime() (*time.Time, error) {
	if o.EffectiveDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EffectiveDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SavingsPlanModelProperties) SetEffectiveDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EffectiveDateTime = &formatted
}

func (o *SavingsPlanModelProperties) GetExpiryDateTimeAsTime() (*time.Time, error) {
	if o.ExpiryDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SavingsPlanModelProperties) SetExpiryDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryDateTime = &formatted
}

func (o *SavingsPlanModelProperties) GetPurchaseDateTimeAsTime() (*time.Time, error) {
	if o.PurchaseDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PurchaseDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SavingsPlanModelProperties) SetPurchaseDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PurchaseDateTime = &formatted
}
