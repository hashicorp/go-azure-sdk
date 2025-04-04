package distributedavailabilitygroups

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CertificateInfo struct {
	CertificateName *string `json:"certificateName,omitempty"`
	ExpiryDate      *string `json:"expiryDate,omitempty"`
}

func (o *CertificateInfo) GetExpiryDateAsTime() (*time.Time, error) {
	if o.ExpiryDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryDate, "2006-01-02T15:04:05Z07:00")
}

func (o *CertificateInfo) SetExpiryDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryDate = &formatted
}
