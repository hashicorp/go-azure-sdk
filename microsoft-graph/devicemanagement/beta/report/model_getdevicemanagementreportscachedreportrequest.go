package report

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementReportsCachedReportRequest struct {
	GroupBy *[]string `json:"groupBy,omitempty"`
	Id      *string   `json:"id,omitempty"`
	OrderBy *[]string `json:"orderBy,omitempty"`
	Search  *string   `json:"search,omitempty"`
	Select  *[]string `json:"select,omitempty"`
	Skip    *int64    `json:"skip,omitempty"`
	Top     *int64    `json:"top,omitempty"`
}
