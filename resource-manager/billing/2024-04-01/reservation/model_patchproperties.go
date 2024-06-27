package reservation

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PatchProperties struct {
	AppliedScopeProperties *ReservationAppliedScopeProperties `json:"appliedScopeProperties,omitempty"`
	AppliedScopeType       *AppliedScopeType                  `json:"appliedScopeType,omitempty"`
	DisplayName            *string                            `json:"displayName,omitempty"`
	InstanceFlexibility    *InstanceFlexibility               `json:"instanceFlexibility,omitempty"`
	Renew                  *bool                              `json:"renew,omitempty"`
	RenewProperties        *PatchPropertiesRenewProperties    `json:"renewProperties,omitempty"`
	ReviewDateTime         *string                            `json:"reviewDateTime,omitempty"`
}

func (o *PatchProperties) GetReviewDateTimeAsTime() (*time.Time, error) {
	if o.ReviewDateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ReviewDateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *PatchProperties) SetReviewDateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ReviewDateTime = &formatted
}
