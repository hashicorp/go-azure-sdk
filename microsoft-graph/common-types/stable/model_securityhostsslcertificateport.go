package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityHostSslCertificatePort struct {
	FirstSeenDateTime *string `json:"firstSeenDateTime,omitempty"`
	LastSeenDateTime  *string `json:"lastSeenDateTime,omitempty"`
	ODataType         *string `json:"@odata.type,omitempty"`
	Port              *int64  `json:"port,omitempty"`
}
