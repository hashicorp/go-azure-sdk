package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemAnalytics struct {
	AllTime           *ItemActivityStat   `json:"allTime,omitempty"`
	Id                *string             `json:"id,omitempty"`
	ItemActivityStats *[]ItemActivityStat `json:"itemActivityStats,omitempty"`
	LastSevenDays     *ItemActivityStat   `json:"lastSevenDays,omitempty"`
	ODataType         *string             `json:"@odata.type,omitempty"`
}
