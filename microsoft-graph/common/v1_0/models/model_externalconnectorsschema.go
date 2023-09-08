package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsSchema struct {
	BaseType   *string                       `json:"baseType,omitempty"`
	Id         *string                       `json:"id,omitempty"`
	ODataType  *string                       `json:"@odata.type,omitempty"`
	Properties *[]ExternalConnectorsProperty `json:"properties,omitempty"`
}
