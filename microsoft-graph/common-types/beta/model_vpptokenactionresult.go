package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenActionResult struct {
	ActionName          *string                          `json:"actionName,omitempty"`
	ActionState         *VppTokenActionResultActionState `json:"actionState,omitempty"`
	LastUpdatedDateTime *string                          `json:"lastUpdatedDateTime,omitempty"`
	ODataType           *string                          `json:"@odata.type,omitempty"`
	StartDateTime       *string                          `json:"startDateTime,omitempty"`
}
