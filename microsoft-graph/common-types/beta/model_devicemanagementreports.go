package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementReports struct {
	CachedReportConfigurations *[]DeviceManagementCachedReportConfiguration `json:"cachedReportConfigurations,omitempty"`
	ExportJobs                 *[]DeviceManagementExportJob                 `json:"exportJobs,omitempty"`
	Id                         *string                                      `json:"id,omitempty"`
	ODataType                  *string                                      `json:"@odata.type,omitempty"`
}
