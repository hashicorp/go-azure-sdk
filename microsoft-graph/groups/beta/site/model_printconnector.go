package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintConnector struct {
	AppVersion               *string          `json:"appVersion,omitempty"`
	DeviceHealth             *DeviceHealth    `json:"deviceHealth,omitempty"`
	DisplayName              *string          `json:"displayName,omitempty"`
	FullyQualifiedDomainName *string          `json:"fullyQualifiedDomainName,omitempty"`
	Id                       *string          `json:"id,omitempty"`
	Location                 *PrinterLocation `json:"location,omitempty"`
	Name                     *string          `json:"name,omitempty"`
	ODataType                *string          `json:"@odata.type,omitempty"`
	OperatingSystem          *string          `json:"operatingSystem,omitempty"`
	RegisteredDateTime       *string          `json:"registeredDateTime,omitempty"`
}
