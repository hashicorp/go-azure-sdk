package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallTranscriptionInfo struct {
	LastModifiedDateTime *string                     `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                     `json:"@odata.type,omitempty"`
	State                *CallTranscriptionInfoState `json:"state,omitempty"`
}
