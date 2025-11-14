package openapis

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PricesheetDownloadProperties struct {
	DownloadURL *string `json:"downloadUrl,omitempty"`
	ValidTill   *string `json:"validTill,omitempty"`
}

func (o *PricesheetDownloadProperties) GetValidTillAsTime() (*time.Time, error) {
	if o.ValidTill == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidTill, "2006-01-02T15:04:05Z07:00")
}

func (o *PricesheetDownloadProperties) SetValidTillAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidTill = &formatted
}
