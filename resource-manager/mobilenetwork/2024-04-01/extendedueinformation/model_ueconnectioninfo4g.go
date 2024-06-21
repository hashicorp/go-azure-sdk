package extendedueinformation

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
