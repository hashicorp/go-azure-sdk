package pricesheets

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DownloadURL struct {
	DownloadURL *string `json:"downloadUrl,omitempty"`
	ExpiryTime  *string `json:"expiryTime,omitempty"`
	ValidTill   *string `json:"validTill,omitempty"`
}

func (o *DownloadURL) GetExpiryTimeAsTime() (*time.Time, error) {
	if o.ExpiryTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DownloadURL) SetExpiryTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpiryTime = &formatted
}

func (o *DownloadURL) GetValidTillAsTime() (*time.Time, error) {
	if o.ValidTill == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ValidTill, "2006-01-02T15:04:05Z07:00")
}

func (o *DownloadURL) SetValidTillAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ValidTill = &formatted
}
