package emailregistration

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailRegistration struct {
	ActivationCode           *string             `json:"activationCode,omitempty"`
	ActivationExpirationDate *string             `json:"activationExpirationDate,omitempty"`
	Email                    *string             `json:"email,omitempty"`
	RegistrationStatus       *RegistrationStatus `json:"registrationStatus,omitempty"`
	TenantId                 *string             `json:"tenantId,omitempty"`
}

func (o *EmailRegistration) GetActivationExpirationDateAsTime() (*time.Time, error) {
	if o.ActivationExpirationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ActivationExpirationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *EmailRegistration) SetActivationExpirationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ActivationExpirationDate = &formatted
}
