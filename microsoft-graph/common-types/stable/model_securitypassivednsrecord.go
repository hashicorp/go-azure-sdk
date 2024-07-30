package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityPassiveDnsRecord struct {
	Artifact          *SecurityArtifact `json:"artifact,omitempty"`
	CollectedDateTime *string           `json:"collectedDateTime,omitempty"`
	FirstSeenDateTime *string           `json:"firstSeenDateTime,omitempty"`
	Id                *string           `json:"id,omitempty"`
	LastSeenDateTime  *string           `json:"lastSeenDateTime,omitempty"`
	ODataType         *string           `json:"@odata.type,omitempty"`
	ParentHost        *SecurityHost     `json:"parentHost,omitempty"`
	RecordType        *string           `json:"recordType,omitempty"`
}
