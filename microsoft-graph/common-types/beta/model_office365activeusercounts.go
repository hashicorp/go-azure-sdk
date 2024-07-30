package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365ActiveUserCounts struct {
	Exchange          *int64  `json:"exchange,omitempty"`
	Id                *string `json:"id,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	Office365         *int64  `json:"office365,omitempty"`
	OneDrive          *int64  `json:"oneDrive,omitempty"`
	ReportDate        *string `json:"reportDate,omitempty"`
	ReportPeriod      *string `json:"reportPeriod,omitempty"`
	ReportRefreshDate *string `json:"reportRefreshDate,omitempty"`
	SharePoint        *int64  `json:"sharePoint,omitempty"`
	SkypeForBusiness  *int64  `json:"skypeForBusiness,omitempty"`
	Teams             *int64  `json:"teams,omitempty"`
	Yammer            *int64  `json:"yammer,omitempty"`
}
