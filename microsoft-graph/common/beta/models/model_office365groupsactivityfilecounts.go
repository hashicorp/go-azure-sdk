package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Office365GroupsActivityFileCounts struct {
	Active            *int64  `json:"active,omitempty"`
	Id                *string `json:"id,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	ReportDate        *string `json:"reportDate,omitempty"`
	ReportPeriod      *string `json:"reportPeriod,omitempty"`
	ReportRefreshDate *string `json:"reportRefreshDate,omitempty"`
	Total             *int64  `json:"total,omitempty"`
}
