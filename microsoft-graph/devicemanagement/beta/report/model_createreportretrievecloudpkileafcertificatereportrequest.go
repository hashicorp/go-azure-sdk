package report

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateReportRetrieveCloudPkiLeafCertificateReportRequest struct {
	CertificationAuthorityId nullable.Type[string] `json:"certificationAuthorityId,omitempty"`
	Filter                   nullable.Type[string] `json:"filter,omitempty"`
	OrderBy                  *[]string             `json:"orderBy,omitempty"`
	Search                   nullable.Type[string] `json:"search,omitempty"`
	Select                   *[]string             `json:"select,omitempty"`
	Skip                     *int64                `json:"skip,omitempty"`
	Top                      *int64                `json:"top,omitempty"`
}
