package report

import (
	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateReportRetrieveCloudPkiLeafCertificateSummaryReportRequest struct {
	CertificationAuthorityId nullable.Type[string] `json:"certificationAuthorityId,omitempty"`
	Select                   *[]string             `json:"select,omitempty"`
}
