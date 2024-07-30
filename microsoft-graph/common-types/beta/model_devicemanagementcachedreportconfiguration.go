package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCachedReportConfiguration struct {
	ExpirationDateTime  *string                                          `json:"expirationDateTime,omitempty"`
	Filter              *string                                          `json:"filter,omitempty"`
	Id                  *string                                          `json:"id,omitempty"`
	LastRefreshDateTime *string                                          `json:"lastRefreshDateTime,omitempty"`
	Metadata            *string                                          `json:"metadata,omitempty"`
	ODataType           *string                                          `json:"@odata.type,omitempty"`
	OrderBy             *[]string                                        `json:"orderBy,omitempty"`
	ReportName          *string                                          `json:"reportName,omitempty"`
	Select              *[]string                                        `json:"select,omitempty"`
	Status              *DeviceManagementCachedReportConfigurationStatus `json:"status,omitempty"`
}
