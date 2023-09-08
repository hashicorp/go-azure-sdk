package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityIndicator struct {
	Artifact  *SecurityArtifact        `json:"artifact,omitempty"`
	Id        *string                  `json:"id,omitempty"`
	ODataType *string                  `json:"@odata.type,omitempty"`
	Source    *SecurityIndicatorSource `json:"source,omitempty"`
}
