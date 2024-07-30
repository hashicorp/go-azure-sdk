package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddIn struct {
	Id         *string     `json:"id,omitempty"`
	ODataType  *string     `json:"@odata.type,omitempty"`
	Properties *[]KeyValue `json:"properties,omitempty"`
	Type       *string     `json:"type,omitempty"`
}
