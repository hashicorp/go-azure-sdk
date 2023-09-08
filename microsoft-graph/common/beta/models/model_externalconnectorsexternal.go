package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternal struct {
	Connections  *[]ExternalConnectorsExternalConnection `json:"connections,omitempty"`
	IndustryData *IndustryDataIndustryDataRoot           `json:"industryData,omitempty"`
	ODataType    *string                                 `json:"@odata.type,omitempty"`
}
