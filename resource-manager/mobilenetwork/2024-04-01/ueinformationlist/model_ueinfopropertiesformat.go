package ueinformationlist

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeInfoPropertiesFormat struct {
	LastReadAt    *string      `json:"lastReadAt,omitempty"`
	RatType       RatType      `json:"ratType"`
	UeIPAddresses *[]DnnIPPair `json:"ueIpAddresses,omitempty"`
	UeState       UeState      `json:"ueState"`
}

func (o *UeInfoPropertiesFormat) GetLastReadAtAsTime() (*time.Time, error) {
	if o.LastReadAt == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastReadAt, "2006-01-02T15:04:05Z07:00")
}

func (o *UeInfoPropertiesFormat) SetLastReadAtAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastReadAt = &formatted
}
