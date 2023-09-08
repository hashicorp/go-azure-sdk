package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCExportJob struct {
	ExpirationDateTime *string                          `json:"expirationDateTime,omitempty"`
	ExportJobStatus    *CloudPCExportJobExportJobStatus `json:"exportJobStatus,omitempty"`
	ExportUrl          *string                          `json:"exportUrl,omitempty"`
	Filter             *string                          `json:"filter,omitempty"`
	Format             *string                          `json:"format,omitempty"`
	Id                 *string                          `json:"id,omitempty"`
	ODataType          *string                          `json:"@odata.type,omitempty"`
	ReportName         *CloudPCExportJobReportName      `json:"reportName,omitempty"`
	RequestDateTime    *string                          `json:"requestDateTime,omitempty"`
	Select             *[]string                        `json:"select,omitempty"`
}
