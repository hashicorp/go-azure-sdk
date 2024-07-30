package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostTracker struct {
	FirstSeenDateTime *string       `json:"firstSeenDateTime,omitempty"`
	Host              *SecurityHost `json:"host,omitempty"`
	Id                *string       `json:"id,omitempty"`
	Kind              *string       `json:"kind,omitempty"`
	LastSeenDateTime  *string       `json:"lastSeenDateTime,omitempty"`
	ODataType         *string       `json:"@odata.type,omitempty"`
	Value             *string       `json:"value,omitempty"`
}
