package report

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceManagementReportsDeviceNonComplianceReportRequest struct {
	Filter    *string   `json:"filter,omitempty"`
	GroupBy   *[]string `json:"groupBy,omitempty"`
	Name      *string   `json:"name,omitempty"`
	OrderBy   *[]string `json:"orderBy,omitempty"`
	Search    *string   `json:"search,omitempty"`
	Select    *[]string `json:"select,omitempty"`
	SessionId *string   `json:"sessionId,omitempty"`
	Skip      *int64    `json:"skip,omitempty"`
	Top       *int64    `json:"top,omitempty"`
}
