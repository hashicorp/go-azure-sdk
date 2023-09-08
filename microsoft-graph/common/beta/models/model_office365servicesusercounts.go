package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365ServicesUserCounts struct {
	ExchangeActive           *int64  `json:"exchangeActive,omitempty"`
	ExchangeInactive         *int64  `json:"exchangeInactive,omitempty"`
	Id                       *string `json:"id,omitempty"`
	ODataType                *string `json:"@odata.type,omitempty"`
	Office365Active          *int64  `json:"office365Active,omitempty"`
	Office365Inactive        *int64  `json:"office365Inactive,omitempty"`
	OneDriveActive           *int64  `json:"oneDriveActive,omitempty"`
	OneDriveInactive         *int64  `json:"oneDriveInactive,omitempty"`
	ReportPeriod             *string `json:"reportPeriod,omitempty"`
	ReportRefreshDate        *string `json:"reportRefreshDate,omitempty"`
	SharePointActive         *int64  `json:"sharePointActive,omitempty"`
	SharePointInactive       *int64  `json:"sharePointInactive,omitempty"`
	SkypeForBusinessActive   *int64  `json:"skypeForBusinessActive,omitempty"`
	SkypeForBusinessInactive *int64  `json:"skypeForBusinessInactive,omitempty"`
	TeamsActive              *int64  `json:"teamsActive,omitempty"`
	TeamsInactive            *int64  `json:"teamsInactive,omitempty"`
	YammerActive             *int64  `json:"yammerActive,omitempty"`
	YammerInactive           *int64  `json:"yammerInactive,omitempty"`
}
