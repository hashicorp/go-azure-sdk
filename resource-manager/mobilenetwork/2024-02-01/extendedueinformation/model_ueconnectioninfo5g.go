package extendedueinformation

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
