package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubjectRightsRequestHistory struct {
	ChangedBy     *IdentitySet                            `json:"changedBy,omitempty"`
	EventDateTime *string                                 `json:"eventDateTime,omitempty"`
	ODataType     *string                                 `json:"@odata.type,omitempty"`
	Stage         *SubjectRightsRequestHistoryStage       `json:"stage,omitempty"`
	StageStatus   *SubjectRightsRequestHistoryStageStatus `json:"stageStatus,omitempty"`
	Type          *string                                 `json:"type,omitempty"`
}
