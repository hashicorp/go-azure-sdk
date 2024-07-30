package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostPair struct {
	ChildHost         *SecurityHost `json:"childHost,omitempty"`
	FirstSeenDateTime *string       `json:"firstSeenDateTime,omitempty"`
	Id                *string       `json:"id,omitempty"`
	LastSeenDateTime  *string       `json:"lastSeenDateTime,omitempty"`
	LinkKind          *string       `json:"linkKind,omitempty"`
	ODataType         *string       `json:"@odata.type,omitempty"`
	ParentHost        *SecurityHost `json:"parentHost,omitempty"`
}
