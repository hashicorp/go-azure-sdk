package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsExternalItem struct {
	Acl        *[]ExternalConnectorsAcl               `json:"acl,omitempty"`
	Activities *[]ExternalConnectorsExternalActivity  `json:"activities,omitempty"`
	Content    *ExternalConnectorsExternalItemContent `json:"content,omitempty"`
	Id         *string                                `json:"id,omitempty"`
	ODataType  *string                                `json:"@odata.type,omitempty"`
	Properties *ExternalConnectorsProperties          `json:"properties,omitempty"`
}
