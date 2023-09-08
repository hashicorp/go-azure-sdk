package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIntelligenceProfileIndicator struct {
	Artifact          *SecurityArtifact                           `json:"artifact,omitempty"`
	FirstSeenDateTime *string                                     `json:"firstSeenDateTime,omitempty"`
	Id                *string                                     `json:"id,omitempty"`
	LastSeenDateTime  *string                                     `json:"lastSeenDateTime,omitempty"`
	ODataType         *string                                     `json:"@odata.type,omitempty"`
	Source            *SecurityIntelligenceProfileIndicatorSource `json:"source,omitempty"`
}
