package domains

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DomainPurchaseConsent struct {
	AgreedAt      *string   `json:"agreedAt,omitempty"`
	AgreedBy      *string   `json:"agreedBy,omitempty"`
	AgreementKeys *[]string `json:"agreementKeys,omitempty"`
}

func (o *DomainPurchaseConsent) GetAgreedAtAsTime() (*time.Time, error) {
	if o.AgreedAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AgreedAt, "2006-01-02T15:04:05Z07:00")
}

func (o *DomainPurchaseConsent) SetAgreedAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AgreedAt = &formatted
}
