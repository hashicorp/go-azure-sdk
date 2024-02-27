package billingsubscriptions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnrollmentAccountSubscriptionDetails struct {
	EnrollmentAccountStartDate          *string                              `json:"enrollmentAccountStartDate,omitempty"`
	SubscriptionEnrollmentAccountStatus *SubscriptionEnrollmentAccountStatus `json:"subscriptionEnrollmentAccountStatus,omitempty"`
}

func (o *EnrollmentAccountSubscriptionDetails) GetEnrollmentAccountStartDateAsTime() (*time.Time, error) {
	if o.EnrollmentAccountStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EnrollmentAccountStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *EnrollmentAccountSubscriptionDetails) SetEnrollmentAccountStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EnrollmentAccountStartDate = &formatted
}
