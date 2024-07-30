package virtualendpointreport

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementVirtualEndpointReportsActionStatusReportRequest struct {
	Filter  *string   `json:"filter,omitempty"`
	GroupBy *[]string `json:"groupBy,omitempty"`
	OrderBy *[]string `json:"orderBy,omitempty"`
	Search  *string   `json:"search,omitempty"`
	Select  *[]string `json:"select,omitempty"`
	Skip    *int64    `json:"skip,omitempty"`
	Top     *int64    `json:"top,omitempty"`
}
