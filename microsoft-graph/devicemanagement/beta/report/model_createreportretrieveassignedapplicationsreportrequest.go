package report

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateReportRetrieveAssignedApplicationsReportRequest struct {
	Filter  nullable.Type[string] `json:"filter,omitempty"`
	Groupby nullable.Type[string] `json:"groupby,omitempty"`
	OrderBy nullable.Type[string] `json:"orderby,omitempty"`
	Search  nullable.Type[string] `json:"search,omitempty"`
	Select  nullable.Type[string] `json:"select,omitempty"`
	Skip    nullable.Type[int64]  `json:"skip,omitempty"`
	Top     nullable.Type[int64]  `json:"top,omitempty"`
}
