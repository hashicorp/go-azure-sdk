package identityassignments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityValidationStatus struct {
	Error          *ErrorDetail                             `json:"error,omitempty"`
	State          ChangeDataCaptureIdentityValidationState `json:"state"`
	ValidatedAtUtc *string                                  `json:"validatedAtUtc,omitempty"`
}

func (o *IdentityValidationStatus) GetValidatedAtUtcAsTime() (*time.Time, error) {
	if o.ValidatedAtUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidatedAtUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *IdentityValidationStatus) SetValidatedAtUtcAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidatedAtUtc = &formatted
}
