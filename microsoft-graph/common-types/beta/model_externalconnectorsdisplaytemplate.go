package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsDisplayTemplate struct {
	Id        *string                           `json:"id,omitempty"`
	Layout    *Json                             `json:"layout,omitempty"`
	ODataType *string                           `json:"@odata.type,omitempty"`
	Priority  *int64                            `json:"priority,omitempty"`
	Rules     *[]ExternalConnectorsPropertyRule `json:"rules,omitempty"`
}
