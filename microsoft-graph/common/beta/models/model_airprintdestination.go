package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AirPrintDestination struct {
	ForceTls     *bool   `json:"forceTls,omitempty"`
	IpAddress    *string `json:"ipAddress,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Port         *int64  `json:"port,omitempty"`
	ResourcePath *string `json:"resourcePath,omitempty"`
}
