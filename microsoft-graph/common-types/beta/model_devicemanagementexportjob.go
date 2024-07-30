package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementExportJob struct {
	ExpirationDateTime *string                                    `json:"expirationDateTime,omitempty"`
	Filter             *string                                    `json:"filter,omitempty"`
	Format             *DeviceManagementExportJobFormat           `json:"format,omitempty"`
	Id                 *string                                    `json:"id,omitempty"`
	LocalizationType   *DeviceManagementExportJobLocalizationType `json:"localizationType,omitempty"`
	ODataType          *string                                    `json:"@odata.type,omitempty"`
	ReportName         *string                                    `json:"reportName,omitempty"`
	RequestDateTime    *string                                    `json:"requestDateTime,omitempty"`
	Select             *[]string                                  `json:"select,omitempty"`
	SnapshotId         *string                                    `json:"snapshotId,omitempty"`
	Status             *DeviceManagementExportJobStatus           `json:"status,omitempty"`
	Url                *string                                    `json:"url,omitempty"`
}
