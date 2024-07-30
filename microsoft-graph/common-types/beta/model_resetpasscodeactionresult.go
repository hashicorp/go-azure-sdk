package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetPasscodeActionResult struct {
	ActionName          *string                               `json:"actionName,omitempty"`
	ActionState         *ResetPasscodeActionResultActionState `json:"actionState,omitempty"`
	ErrorCode           *int64                                `json:"errorCode,omitempty"`
	LastUpdatedDateTime *string                               `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                               `json:"@odata.type,omitempty"`
	Passcode            *string                               `json:"passcode,omitempty"`
	StartDateTime       *string                               `json:"startDateTime,omitempty"`
}
