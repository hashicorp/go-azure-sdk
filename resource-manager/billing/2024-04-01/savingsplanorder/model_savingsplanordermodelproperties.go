package savingsplanorder

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanOrderModelProperties struct {
	BenefitStartTime   *string                 `json:"benefitStartTime,omitempty"`
	BillingAccountId   *string                 `json:"billingAccountId,omitempty"`
	BillingPlan        *BillingPlan            `json:"billingPlan,omitempty"`
	BillingProfileId   *string                 `json:"billingProfileId,omitempty"`
	BillingScopeId     *string                 `json:"billingScopeId,omitempty"`
	CustomerId         *string                 `json:"customerId,omitempty"`
	DisplayName        *string                 `json:"displayName,omitempty"`
	ExpiryDate         *string                 `json:"expiryDate,omitempty"`
	ExpiryDateTime     *string                 `json:"expiryDateTime,omitempty"`
	ExtendedStatusInfo *ExtendedStatusInfo     `json:"extendedStatusInfo,omitempty"`
	PlanInformation    *BillingPlanInformation `json:"planInformation,omitempty"`
	ProductCode        *string                 `json:"productCode,omitempty"`
	ProvisioningState  *string                 `json:"provisioningState,omitempty"`
	SavingsPlans       *[]string               `json:"savingsPlans,omitempty"`
	Term               *SavingsPlanTerm        `json:"term,omitempty"`
}

func (o *SavingsPlanOrderModelProperties) GetBenefitStartTimeAsTime() (*time.Time, error) {
	if o.BenefitStartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BenefitStartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SavingsPlanOrderModelProperties) SetBenefitStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BenefitStartTime = &formatted
}

func (o *SavingsPlanOrderModelProperties) GetExpiryDateTimeAsTime() (*time.Time, error) {
	if o.ExpiryDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SavingsPlanOrderModelProperties) SetExpiryDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryDateTime = &formatted
}
