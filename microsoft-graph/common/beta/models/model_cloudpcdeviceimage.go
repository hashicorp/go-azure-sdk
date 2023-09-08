package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDeviceImage struct {
	DisplayName           *string                          `json:"displayName,omitempty"`
	ExpirationDate        *string                          `json:"expirationDate,omitempty"`
	Id                    *string                          `json:"id,omitempty"`
	LastModifiedDateTime  *string                          `json:"lastModifiedDateTime,omitempty"`
	ODataType             *string                          `json:"@odata.type,omitempty"`
	OperatingSystem       *string                          `json:"operatingSystem,omitempty"`
	OsBuildNumber         *string                          `json:"osBuildNumber,omitempty"`
	OsStatus              *CloudPCDeviceImageOsStatus      `json:"osStatus,omitempty"`
	SourceImageResourceId *string                          `json:"sourceImageResourceId,omitempty"`
	Status                *CloudPCDeviceImageStatus        `json:"status,omitempty"`
	StatusDetails         *CloudPCDeviceImageStatusDetails `json:"statusDetails,omitempty"`
	Version               *string                          `json:"version,omitempty"`
}
