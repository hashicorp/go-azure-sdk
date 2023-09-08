package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityRunDetails struct {
	ErrorCode       *SecurityRunDetailsErrorCode `json:"errorCode,omitempty"`
	FailureReason   *string                      `json:"failureReason,omitempty"`
	LastRunDateTime *string                      `json:"lastRunDateTime,omitempty"`
	ODataType       *string                      `json:"@odata.type,omitempty"`
	Status          *SecurityRunDetailsStatus    `json:"status,omitempty"`
}
