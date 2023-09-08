package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsFailureInfo struct {
	ODataType *string                      `json:"@odata.type,omitempty"`
	Reason    *string                      `json:"reason,omitempty"`
	Stage     *CallRecordsFailureInfoStage `json:"stage,omitempty"`
}
