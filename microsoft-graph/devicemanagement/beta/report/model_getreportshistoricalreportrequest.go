package report

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetReportsHistoricalReportRequest struct {
	Filter  nullable.Type[string] `json:"filter,omitempty"`
	GroupBy *[]string             `json:"groupBy,omitempty"`
	Name    nullable.Type[string] `json:"name,omitempty"`
	OrderBy *[]string             `json:"orderBy,omitempty"`
	Search  nullable.Type[string] `json:"search,omitempty"`
	Select  *[]string             `json:"select,omitempty"`
	Skip    nullable.Type[int64]  `json:"skip,omitempty"`
	Top     nullable.Type[int64]  `json:"top,omitempty"`
}
