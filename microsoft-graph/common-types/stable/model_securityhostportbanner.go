package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostPortBanner struct {
	Banner            *string `json:"banner,omitempty"`
	FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`
	LastSeenDateTime  *string `json:"lastSeenDateTime,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	ScanProtocol      *string `json:"scanProtocol,omitempty"`
	TimesObserved     *int64  `json:"timesObserved,omitempty"`
}
