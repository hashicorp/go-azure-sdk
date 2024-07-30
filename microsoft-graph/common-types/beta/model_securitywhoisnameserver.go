package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityWhoisNameserver struct {
	FirstSeenDateTime *string       `json:"firstSeenDateTime,omitempty"`
	Host              *SecurityHost `json:"host,omitempty"`
	LastSeenDateTime  *string       `json:"lastSeenDateTime,omitempty"`
	ODataType         *string       `json:"@odata.type,omitempty"`
}
