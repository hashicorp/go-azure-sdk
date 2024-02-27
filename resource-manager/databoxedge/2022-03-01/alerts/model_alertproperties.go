package alerts

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertProperties struct {
	AlertType           *string            `json:"alertType,omitempty"`
	AppearedAtDateTime  *string            `json:"appearedAtDateTime,omitempty"`
	DetailedInformation *map[string]string `json:"detailedInformation,omitempty"`
	ErrorDetails        *AlertErrorDetails `json:"errorDetails,omitempty"`
	Recommendation      *string            `json:"recommendation,omitempty"`
	Severity            *AlertSeverity     `json:"severity,omitempty"`
	Title               *string            `json:"title,omitempty"`
}

func (o *AlertProperties) GetAppearedAtDateTimeAsTime() (*time.Time, error) {
	if o.AppearedAtDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AppearedAtDateTime, "2006-01-02T15:04:05Z07:00")
}
