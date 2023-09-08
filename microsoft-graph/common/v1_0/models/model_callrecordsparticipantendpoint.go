package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsParticipantEndpoint struct {
	CpuCoresCount          *int64                   `json:"cpuCoresCount,omitempty"`
	CpuName                *string                  `json:"cpuName,omitempty"`
	CpuProcessorSpeedInMhz *int64                   `json:"cpuProcessorSpeedInMhz,omitempty"`
	Feedback               *CallRecordsUserFeedback `json:"feedback,omitempty"`
	Identity               *IdentitySet             `json:"identity,omitempty"`
	Name                   *string                  `json:"name,omitempty"`
	ODataType              *string                  `json:"@odata.type,omitempty"`
	UserAgent              *CallRecordsUserAgent    `json:"userAgent,omitempty"`
}
