package extendedueinformation

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeConnectionInfo4G struct {
	EnbS1apId             int64                 `json:"enbS1apId"`
	GlobalRanNodeId       GlobalRanNodeId       `json:"globalRanNodeId"`
	LastActivityTime      *string               `json:"lastActivityTime,omitempty"`
	LastVisitedTai        *string               `json:"lastVisitedTai,omitempty"`
	LocationInfo          *UeLocationInfo       `json:"locationInfo,omitempty"`
	MmeS1apId             int64                 `json:"mmeS1apId"`
	PerUeTnla             *string               `json:"perUeTnla,omitempty"`
	RrcEstablishmentCause RrcEstablishmentCause `json:"rrcEstablishmentCause"`
	UeState               UeState               `json:"ueState"`
	UeUsageSetting        *UeUsageSetting       `json:"ueUsageSetting,omitempty"`
}

func (o *UeConnectionInfo4G) GetLastActivityTimeAsTime() (*time.Time, error) {
	if o.LastActivityTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastActivityTime, "2006-01-02T15:04:05Z07:00")
}

func (o *UeConnectionInfo4G) SetLastActivityTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastActivityTime = &formatted
}
