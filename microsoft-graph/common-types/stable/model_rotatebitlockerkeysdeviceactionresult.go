package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RotateBitLockerKeysDeviceActionResult struct {
	ActionName          *string                                           `json:"actionName,omitempty"`
	ActionState         *RotateBitLockerKeysDeviceActionResultActionState `json:"actionState,omitempty"`
	ErrorCode           *int64                                            `json:"errorCode,omitempty"`
	LastUpdatedDateTime *string                                           `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                                           `json:"@odata.type,omitempty"`
	StartDateTime       *string                                           `json:"startDateTime,omitempty"`
}
