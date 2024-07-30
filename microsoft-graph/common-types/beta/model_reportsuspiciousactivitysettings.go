package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReportSuspiciousActivitySettings struct {
	IncludeTarget      *IncludeTarget                         `json:"includeTarget,omitempty"`
	ODataType          *string                                `json:"@odata.type,omitempty"`
	State              *ReportSuspiciousActivitySettingsState `json:"state,omitempty"`
	VoiceReportingCode *int64                                 `json:"voiceReportingCode,omitempty"`
}
