package migrations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidationDetails struct {
	DbLevelValidationDetails     *[]DbLevelValidationStatus `json:"dbLevelValidationDetails,omitempty"`
	ServerLevelValidationDetails *[]ValidationSummaryItem   `json:"serverLevelValidationDetails,omitempty"`
	Status                       *ValidationState           `json:"status,omitempty"`
	ValidationEndTimeInUtc       *string                    `json:"validationEndTimeInUtc,omitempty"`
	ValidationStartTimeInUtc     *string                    `json:"validationStartTimeInUtc,omitempty"`
}

func (o *ValidationDetails) GetValidationEndTimeInUtcAsTime() (*time.Time, error) {
	if o.ValidationEndTimeInUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidationEndTimeInUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *ValidationDetails) SetValidationEndTimeInUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidationEndTimeInUtc = &formatted
}

func (o *ValidationDetails) GetValidationStartTimeInUtcAsTime() (*time.Time, error) {
	if o.ValidationStartTimeInUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidationStartTimeInUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *ValidationDetails) SetValidationStartTimeInUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidationStartTimeInUtc = &formatted
}
