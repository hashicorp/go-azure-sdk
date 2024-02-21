package extendedueinformation

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeConnectionInfo5G struct {
	AllowedNssai          *[]Snssai             `json:"allowedNssai,omitempty"`
	AmfUeNgapId           int64                 `json:"amfUeNgapId"`
	GlobalRanNodeId       GlobalRanNodeId       `json:"globalRanNodeId"`
	LastActivityTime      *string               `json:"lastActivityTime,omitempty"`
	LastVisitedTai        *string               `json:"lastVisitedTai,omitempty"`
	LocationInfo          *UeLocationInfo       `json:"locationInfo,omitempty"`
	PerUeTnla             *string               `json:"perUeTnla,omitempty"`
	RanUeNgapId           int64                 `json:"ranUeNgapId"`
	RrcEstablishmentCause RrcEstablishmentCause `json:"rrcEstablishmentCause"`
	UeState               UeState               `json:"ueState"`
	UeUsageSetting        *UeUsageSetting       `json:"ueUsageSetting,omitempty"`
}

func (o *UeConnectionInfo5G) GetLastActivityTimeAsTime() (*time.Time, error) {
	if o.LastActivityTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastActivityTime, "2006-01-02T15:04:05Z07:00")
}

func (o *UeConnectionInfo5G) SetLastActivityTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastActivityTime = &formatted
}
