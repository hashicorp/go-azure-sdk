package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostCookie struct {
	Domain            *string       `json:"domain,omitempty"`
	FirstSeenDateTime *string       `json:"firstSeenDateTime,omitempty"`
	Host              *SecurityHost `json:"host,omitempty"`
	Id                *string       `json:"id,omitempty"`
	LastSeenDateTime  *string       `json:"lastSeenDateTime,omitempty"`
	Name              *string       `json:"name,omitempty"`
	ODataType         *string       `json:"@odata.type,omitempty"`
}
