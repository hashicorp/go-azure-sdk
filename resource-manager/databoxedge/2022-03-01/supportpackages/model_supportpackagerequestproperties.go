package supportpackages

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SupportPackageRequestProperties struct {
	Include          *string `json:"include,omitempty"`
	MaximumTimeStamp *string `json:"maximumTimeStamp,omitempty"`
	MinimumTimeStamp *string `json:"minimumTimeStamp,omitempty"`
}

func (o *SupportPackageRequestProperties) GetMaximumTimeStampAsTime() (*time.Time, error) {
	if o.MaximumTimeStamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MaximumTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SupportPackageRequestProperties) SetMaximumTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MaximumTimeStamp = &formatted
}

func (o *SupportPackageRequestProperties) GetMinimumTimeStampAsTime() (*time.Time, error) {
	if o.MinimumTimeStamp == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.MinimumTimeStamp, "2006-01-02T15:04:05Z07:00")
}

func (o *SupportPackageRequestProperties) SetMinimumTimeStampAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.MinimumTimeStamp = &formatted
}
