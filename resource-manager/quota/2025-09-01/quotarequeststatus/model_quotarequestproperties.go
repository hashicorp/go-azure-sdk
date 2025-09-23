package quotarequeststatus

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QuotaRequestProperties struct {
	Error             *ServiceErrorDetail `json:"error,omitempty"`
	Message           *string             `json:"message,omitempty"`
	ProvisioningState *QuotaRequestState  `json:"provisioningState,omitempty"`
	RequestSubmitTime *string             `json:"requestSubmitTime,omitempty"`
	Value             *[]SubRequest       `json:"value,omitempty"`
}

func (o *QuotaRequestProperties) GetRequestSubmitTimeAsTime() (*time.Time, error) {
	if o.RequestSubmitTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RequestSubmitTime, "2006-01-02T15:04:05Z07:00")
}

func (o *QuotaRequestProperties) SetRequestSubmitTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RequestSubmitTime = &formatted
}
